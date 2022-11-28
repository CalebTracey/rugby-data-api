package routes

import (
	"encoding/json"
	"github.com/calebtracey/rugby-data-api/internal/facade"
	_ "github.com/calebtracey/rugby-data-api/internal/routes/statik"
	lbReq "github.com/calebtracey/rugby-models/pkg/dtos/request/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	lbRes "github.com/calebtracey/rugby-models/pkg/dtos/response/leaderboard"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Handler struct {
	Service facade.APIFacadeI
}

func (h *Handler) InitializeRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Handle("/health", h.HealthCheck()).Methods(http.MethodGet)
	r.Handle("/leaderboard", h.LeaderboardHandler()).Methods(http.MethodPost)
	r.Handle("/leaderboards", h.AllLeaderboardsHandler()).Methods(http.MethodGet)

	staticFs, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(staticFs)
	sh := http.StripPrefix("/swagger-ui/", staticServer)
	r.PathPrefix("/swagger-ui/").Handler(sh)

	return r
}

func (h *Handler) LeaderboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		var leaderboardRequest lbReq.Request
		var leaderboardResponse lbRes.Response
		if err := leaderboardRequest.FromJSON(r.Body); err != nil {
			leaderboardResponse.Message.ErrorLog = response.ErrorLogs{*err}
		}
		leaderboardResponse = h.Service.GetLeaderboardData(r.Context(), leaderboardRequest)

		if err := leaderboardResponse.ToJSON(w); err != nil {
			logrus.Errorf("failed to marshal response: %s", err.RootCause)
			leaderboardResponse.Message.ErrorLog = response.ErrorLogs{*err}
		}
		statusCode := leaderboardResponse.Message.ErrorLog.GetHTTPStatus(len(leaderboardResponse.LeaderboardData))
		leaderboardResponse.Message.AddMessageDetails(startTime)

		response.WriteHeader(w, statusCode)
	}
}

func (h *Handler) AllLeaderboardsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		leaderboardResponse := h.Service.GetAllLeaderboardData(r.Context())

		if err := leaderboardResponse.ToJSON(w); err != nil {
			logrus.Errorf("failed to marshal response: %s", err.RootCause)
			leaderboardResponse.Message.ErrorLog = response.ErrorLogs{*err}
		}
		statusCode := leaderboardResponse.Message.ErrorLog.GetHTTPStatus(len(leaderboardResponse.LeaderboardData))
		leaderboardResponse.Message.AddMessageDetails(startTime)

		response.WriteHeader(w, statusCode)
	}
}

func (h *Handler) HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		if err != nil {
			logrus.Errorln(err.Error())
			return
		}
	}
}
