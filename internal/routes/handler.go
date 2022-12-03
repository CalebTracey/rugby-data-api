package routes

import (
	"encoding/json"
	"github.com/calebtracey/rugby-data-api/internal/facade"
	_ "github.com/calebtracey/rugby-data-api/internal/routes/statik"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Handler struct {
	Service facade.APIFacadeI
}

func (h *Handler) InitializeRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/health", h.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/leaderboard", h.LeaderboardHandler).Methods(http.MethodPost)
	r.HandleFunc("/leaderboards", h.AllLeaderboardsHandler).Methods(http.MethodGet)

	staticFs, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(staticFs)
	sh := http.StripPrefix("/swagger-ui/", staticServer)
	r.PathPrefix("/swagger-ui/").Handler(sh)

	return r
}

func (h *Handler) LeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	apiRequest := leaderboard.RequestFromJSON(r.Body)
	apiResponse := h.Service.GetLeaderboardData(r.Context(), apiRequest)

	if err := apiResponse.ResponseToJSON(w); err != nil {
		log.Errorf("failed to marshal response: %s", err.RootCause)
		apiResponse.Message.ErrorLog = response.ErrorLogs{*err}
	}
	statusCode := apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.LeaderboardData))
	apiResponse.Message.AddMessageDetails(startTime)

	response.WriteHeader(w, statusCode)
}

func (h *Handler) AllLeaderboardsHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	apiResponse := h.Service.GetAllLeaderboardData(r.Context())

	if err := apiResponse.ResponseToJSON(w); err != nil {
		log.Errorf("failed to marshal response: %s", err.RootCause)
		apiResponse.Message.ErrorLog = response.ErrorLogs{*err}
	}
	statusCode := apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.LeaderboardData))
	apiResponse.Message.AddMessageDetails(startTime)

	response.WriteHeader(w, statusCode)

}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	defer r.Context().Done()
	if err := json.NewEncoder(w).Encode(map[string]bool{"ok": true}); err != nil {
		log.Errorln(err.Error())
		response.WriteHeader(w, 503)
		return
	}
}
