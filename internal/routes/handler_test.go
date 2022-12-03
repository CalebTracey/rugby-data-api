package routes

import (
	"github.com/calebtracey/rugby-data-api/internal/facade"
	"github.com/calebtracey/rugby-data-api/internal/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	leaderboardBody = "{\"competitions\":[{\"name\":\"six nations\"},{\"name\":\"premiership\"}]}"
)

func TestHandler_InitializeRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFacade := mocks.NewMockAPIFacadeI(ctrl)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		Service facade.APIFacadeI
		args    args
	}{
		{
			name:    "Happy Path - /leaderboard",
			Service: mockFacade,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/leaderboard", strings.NewReader(leaderboardBody)),
			},
		},
		{
			name:    "Happy Path - /leaderboards",
			Service: mockFacade,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/leaderboards", nil),
			},
		},
		{
			name:    "Happy Path - /health",
			Service: mockFacade,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/health", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.Service,
			}

			switch tt.args.r.RequestURI {
			case "/leaderboard":
				mockFacade.EXPECT().GetLeaderboardData(gomock.Any(), gomock.Any())
			case "/leaderboards":
				mockFacade.EXPECT().GetAllLeaderboardData(gomock.Any())
			default:
			}

			h.InitializeRoutes().ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}
