package routes

import (
	"bytes"
	"encoding/json"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

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

func setMessage(startTime time.Time, msg response.Message) (int, response.Message) {
	status, _ := strconv.Atoi(msg.Status)
	hn, _ := os.Hostname()
	msg.HostName = hn
	msg.TimeTaken = time.Since(startTime).String()
	return status, msg
}

func errorLogs(errors []error, rootCause string, status int) response.ErrorLogs {
	var errLogs response.ErrorLogs
	for _, err := range errors {
		errLogs = append(errLogs, response.ErrorLog{
			RootCause:  rootCause,
			StatusCode: strconv.Itoa(status),
			Trace:      err.Error(),
		})
	}
	return errLogs
}
