package routes

import (
	"bytes"
	"embed"
	"encoding/json"
	"github.com/calebtracey/rugby-data-api/external/models/request"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	"github.com/calebtracey/rugby-data-api/internal/facade"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"time"
)

//go:embed static
var content embed.FS

type Handler struct {
	Service facade.APIFacadeI
}

func (h *Handler) InitializeRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Handle("/health", h.HealthCheck()).Methods(http.MethodGet)
	r.Handle("/competition", h.CompetitionHandler()).Methods(http.MethodPost)

	fsys, _ := fs.Sub(content, "static")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(fsys))))

	return r
}

func (h *Handler) CompetitionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		var compResponse response.CompetitionResponse
		var compRequest request.CompetitionRequest

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

		compResponse = h.Service.SixNationsResults(r.Context(), compRequest)
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

func writeHeader(w http.ResponseWriter, code int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	return w
}

func renderResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	if _, err = w.Write(content); err != nil {
		logrus.Error(err)
	}
}

func readBody(body io.ReadCloser) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, copyErr := io.Copy(buf, body)

	if copyErr != nil {
		return nil, copyErr
	}
	return buf.Bytes(), nil
}

func errorLogs(errors []error, rootCause string, status int) []response.ErrorLog {
	var errLogs []response.ErrorLog
	for _, err := range errors {
		errLogs = append(errLogs, response.ErrorLog{
			RootCause:  rootCause,
			StatusCode: strconv.Itoa(status),
			Trace:      err.Error(),
		})
	}
	return errLogs
}
