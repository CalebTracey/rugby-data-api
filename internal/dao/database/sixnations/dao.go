package sixnations

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	"github.com/calebtracey/rugby-data-api/internal/dao/database/psql"
)

//go:generate mockgen -destination=mockDao.go -package=sixnations . SNDAOI
type SNDAOI interface {
	GetTeams(ctx context.Context, query string) (psqlTeamsResponse response.PSQLTeamsResponse, err *response.ErrorLog)
}

type SNDAO struct {
	PSQLDAO    psql.DAOI
	PSQLMapper Mapper
}

func (s SNDAO) GetTeams(ctx context.Context, query string) (psqlTeamsResponse response.PSQLTeamsResponse, err *response.ErrorLog) {
	rows, err := s.PSQLDAO.FindAll(ctx, query)
	if err != nil {
		return psqlTeamsResponse, err
	}
	psqlTeamsResponse.Teams = s.PSQLMapper.MapPSQLRowsToTeamData(rows)

	return psqlTeamsResponse, err
}
