package comp

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/dao/comp"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/request"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockFacade.go -package=compmocks . FacadeI
type FacadeI interface {
	LeaderboardData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse)
	AllLeaderboardData(ctx context.Context) (resp response.LeaderboardResponse)
}

type Facade struct {
	CompDAO  comp.DAOI
	DbMapper psql.MapperI
}

func (s Facade) LeaderboardData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse) {
	var data dtos.CompetitionLeaderboardDataList
	//TODO make this concurrent
	for _, c := range req.Competitions {
		compName, compId := getCompId(c.Name)
		teamsQuery := s.DbMapper.CreatePSQLLeaderboardByIdQuery(compId)
		leaderboardData, err := s.CompDAO.GetLeaderboardData(ctx, teamsQuery)
		if err != nil {
			log.Error(err)
			return response.LeaderboardResponse{
				Message: response.Message{
					ErrorLog: response.ErrorLogs{
						*err,
					},
				},
			}
		}
		compData := s.DbMapper.MapPSQLLeaderboardDataToResponse(compName, compId, leaderboardData)
		data = append(data, compData)
	}
	resp.LeaderboardData = data
	return resp
}

func (s Facade) AllLeaderboardData(ctx context.Context) (resp response.LeaderboardResponse) {
	leaderboardData, err := s.CompDAO.GetAllLeaderboardData(ctx)
	if err != nil {
		log.Error(err)
		return response.LeaderboardResponse{
			Message: response.Message{
				ErrorLog: response.ErrorLogs{
					*err,
				},
			},
		}
	}
	resp = s.DbMapper.MapPSQLAllLeaderboardDataToResponse(leaderboardData)
	return resp
}

func getCompId(compName string) (string, string) {
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

var (
	CompMap = map[string]string{
		SixNations:              SixNationsId,
		RugbyWorldCup:           RugbyWorldCupId,
		Premiership:             PremiershipId,
		Top14:                   Top14Id,
		UnitedRugbyChampionship: UnitedRugbyChampionshipId,
		RugbyChampionship:       RugbyChampionshipId,
	}
)

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
