package sixnations

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	"github.com/calebtracey/rugby-data-api/internal/dao/database/psql"
)

//go:generate mockgen -destination=mockDao.go -package=sixnations . SNDAOI
type SNDAOI interface {
	GetTeams(ctx context.Context, query string) (psqlTeamsResponse response.CompetitionResponse, err *response.ErrorLog)
}

type SNDAO struct {
	PSQLDAO    psql.DAOI
	PSQLMapper MapperI
}

func (s SNDAO) GetTeams(ctx context.Context, query string) (psqlTeamsResponse response.CompetitionResponse, err *response.ErrorLog) {
	rows, err := s.PSQLDAO.FindAll(ctx, query)
	if err != nil {
		return psqlTeamsResponse, err
	}
	//TODO move the psql mapper into the dao where the rows are handled i
	sixNationsData := s.PSQLMapper.MapPSQLRowsToCompetitionData(rows)
	psqlTeamsResponse = s.PSQLMapper.MapCompetitionDataResponse(sixNationsData)

	return psqlTeamsResponse, err
}
