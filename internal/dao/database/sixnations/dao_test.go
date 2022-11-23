package sixnations

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	"github.com/calebtracey/rugby-data-api/internal/dao/database/psql"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestSNDAO_GetTeams(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDao := psql.NewMockDAOI(ctrl)
	mockMapper := NewMockMapperI(ctrl)
	type fields struct {
		PSQLDAO    psql.DAOI
		PSQLMapper MapperI
	}
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		wantPsqlTeamsResponse response.PSQLTeamsResponse
		wantErr               *response.ErrorLog
	}{
		{
			name: "Happy Path",
			fields: fields{
				PSQLDAO:    mockDao,
				PSQLMapper: mockMapper,
			},
			args: args{
				ctx:   context.Background(),
				query: "{}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SNDAO{
				PSQLDAO:    tt.fields.PSQLDAO,
				PSQLMapper: tt.fields.PSQLMapper,
			}
			gotPsqlTeamsResponse, gotErr := s.GetTeams(tt.args.ctx, tt.args.query)
			if !reflect.DeepEqual(gotPsqlTeamsResponse, tt.wantPsqlTeamsResponse) {
				t.Errorf("GetTeams() gotPsqlTeamsResponse = %v, want %v", gotPsqlTeamsResponse, tt.wantPsqlTeamsResponse)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetTeams() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
