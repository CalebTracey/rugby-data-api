package comp

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/calebtracey/rugby-data-api/internal/dao/comp"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockFacade.go -package=compmocks . FacadeI
type FacadeI interface {
	LeaderboardData(ctx context.Context, req leaderboard.Request) (resp leaderboard.Response)
	AllLeaderboardData(ctx context.Context) (resp leaderboard.Response)
}

type Facade struct {
	CompDAO  comp.DAOI
	DbMapper psql.MapperI
}

func (s Facade) LeaderboardData(ctx context.Context, req leaderboard.Request) (resp leaderboard.Response) {
	g, ctx := errgroup.WithContext(ctx)
	results := make([]dtos.CompetitionLeaderboardData, len(req.Competitions))

	for i, competition := range req.Competitions {
		i, competition := i, competition

		compName, compId := compId(competition.Name)
		teamsQuery := s.DbMapper.LeaderboardByIdQuery(compId)

		g.Go(func() error {
			leaderboardData, err := s.CompDAO.LeaderboardData(ctx, teamsQuery)

			if err == nil {
				compData := s.DbMapper.LeaderboardDataToResponse(compName, compId, leaderboardData)
				results[i] = compData
			}
			return err
		})
	}

	if err := g.Wait(); err != nil {
		resp.Message.ErrorLog = response.ErrorLogs{
			*mapError(err, fmt.Sprintf("%s", req)),
		}
	}
	resp.LeaderboardData = results

	return resp
}

func (s Facade) AllLeaderboardData(ctx context.Context) (resp leaderboard.Response) {
	leaderboardData, err := s.CompDAO.AllLeaderboardData(ctx)
	if err != nil {
		resp.Message.ErrorLog = response.ErrorLogs{
			*mapError(err, "all leaderboard request"),
		}
		return resp
	}
	resp = s.DbMapper.AllLeaderboardDataToResponse(leaderboardData)

	return resp
}

func mapError(err error, query string) (errLog *response.ErrorLog) {
	log.Error(err)
	errLog = &response.ErrorLog{
		Query: query,
	}
	if err == sql.ErrNoRows {
		errLog.RootCause = "Not found in database"
		errLog.StatusCode = "404"
		return errLog
	}

	if err != nil {
		errLog.RootCause = err.Error()
	}
	errLog.StatusCode = "500"
	return errLog
}

func compId(compName string) (string, string) {
	c := cases.Title(language.English)
	switch c.String(compName) {
	case SixNations:
		return SixNations, SixNationsId
	case RugbyWorldCup:
		return RugbyWorldCup, RugbyWorldCupId
	case Premiership:
		return Premiership, PremiershipId
	case Top14:
		return Top14, Top14Id
	case UnitedRugbyChampionship:
		return UnitedRugbyChampionship, UnitedRugbyChampionshipId
	case RugbyChampionship:
		return RugbyChampionship, RugbyChampionshipId
	default:
		return "", ""
	}
}

const (
	SixNations   = "Six Nations"
	SixNationsId = "180659"

	RugbyWorldCup   = "Rugby World Cup"
	RugbyWorldCupId = "164205"

	Premiership   = "Premiership"
	PremiershipId = "267979"

	Top14   = "Top 14"
	Top14Id = "270559"

	UnitedRugbyChampionship   = "United Rugby Championship"
	UnitedRugbyChampionshipId = "270557"

	RugbyChampionship   = "Rugby Championship"
	RugbyChampionshipId = "244293"
)
