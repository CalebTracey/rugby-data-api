package routes

import (
	"bytes"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"io"
	"strconv"
)

func readBody(body io.ReadCloser) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, copyErr := io.Copy(buf, body)

	if copyErr != nil {
		return nil, copyErr
	}
	return buf.Bytes(), nil
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
