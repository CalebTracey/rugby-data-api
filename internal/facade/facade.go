package facade

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/request"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	psql2 "github.com/calebtracey/rugby-data-api/internal/facade/psql"
	"strings"
)

type APIFacadeI interface {
	PSQLResults(ctx context.Context, req request.PSQLRequest) (resp response.PSQLResponse)
}

type APIFacade struct {
	PSQLDao psql2.FacadeI
}

func (s APIFacade) PSQLResults(ctx context.Context, req request.PSQLRequest) (resp response.PSQLResponse) {
	//TODO add validation
	if strings.EqualFold(req.RequestType, "Insert") {
		resp = s.PSQLDao.AddNew(ctx, req)
	}
	return resp
}
