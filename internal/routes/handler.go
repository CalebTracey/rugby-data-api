package routes

import (
	"encoding/json"
	"github.com/calebtracey/rugby-data-api/internal/facade"
	_ "github.com/calebtracey/rugby-data-api/internal/routes/statik"
	"github.com/calebtracey/rugby-models/request"
	"github.com/calebtracey/rugby-models/response"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Handler struct {
	Service facade.APIFacadeI
}

func (h *Handler) InitializeRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Handle("/health", h.HealthCheck()).Methods(http.MethodGet)
	r.Handle("/leaderboard", h.LeaderboardHandler()).Methods(http.MethodPost)

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
		var compResponse response.LeaderboardResponse
		var compRequest request.LeaderboardRequest

		defer func() {
			status, _ := strconv.Atoi(compResponse.Message.Status)
			hn, _ := os.Hostname()
			compResponse.Message.HostName = hn
			compResponse.Message.TimeTaken = time.Since(startTime).String()
			_ = json.NewEncoder(writeHeader(w, status)).Encode(compResponse)
		}()
		body, bodyErr := readBody(r.Body)

		if bodyErr != nil {
			compResponse.Message.ErrorLog = errorLogs([]error{bodyErr}, "Unable to read psqlRequest body", http.StatusBadRequest)
			return
		}
		err := json.Unmarshal(body, &compRequest)
		if err != nil {
			compResponse.Message.ErrorLog = errorLogs([]error{err}, "Unable to parse psqlRequest", http.StatusBadRequest)
			return
		}

		compResponse = h.Service.GetCompetitionData(r.Context(), compRequest)
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
