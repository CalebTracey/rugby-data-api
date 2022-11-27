package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"testing"
)

func TestDAO_InsertOne(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}(db)
	tests := []struct {
		name      string
		DB        *sql.DB
		ctx       context.Context
		exec      string
		wantResp  sql.Result
		wantErr   *response.ErrorLog
		mockErr   error
		expectErr bool
	}{
		{
			name:      "Happy Path",
			DB:        db,
			ctx:       context.Background(),
			exec:      ``,
			wantResp:  sqlmock.NewResult(int64(4), int64(12312123123)),
			expectErr: false,
		},
		{
			name: "Sad Path",
			DB:   db,
			ctx:  context.Background(),
			exec: ``,
			wantErr: &response.ErrorLog{
				StatusCode: "500",
				RootCause:  "error",
			},
			mockErr:   errors.New("error"),
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				DB: tt.DB,
			}
			if !tt.expectErr {
				mock.ExpectExec(tt.exec).WillReturnResult(tt.wantResp)
			}
			if tt.expectErr {
				mock.ExpectExec(tt.exec).WillReturnResult(tt.wantResp).WillReturnError(tt.mockErr)
			}
			_, gotErr := s.InsertOne(tt.ctx, tt.exec)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("InsertOne() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestDAO_FindAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}(db)
	mockRows := sqlmock.NewRows([]string{"comp_id", "comp_name", "team_id", "team_name"}).
		AddRow(123, "Test Comp", 1, "Team 1").
		AddRow(123, "Test Comp", 2, "Team 2")
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name          string
		DB            *sql.DB
		args          args
		mockCols      sqlmock.Rows
		wantRows      *sqlmock.Rows
		wantErr       *response.ErrorLog
		expectDbError bool
	}{
		{
			name: "Happy Path",
			args: args{
				ctx:   context.Background(),
				query: fmt.Sprintf(LeaderboardByIdQuery, "123"),
			},
			DB:            db,
			wantRows:      mockRows,
			wantErr:       nil,
			expectDbError: false,
		},
		{
			name: "Sad Path",
			args: args{
				ctx:   context.Background(),
				query: fmt.Sprintf(LeaderboardByIdQuery, "123"),
			},
			DB:       db,
			wantRows: mockRows,
			wantErr: &response.ErrorLog{
				StatusCode: "404",
				RootCause:  "Not found in database",
				Query:      fmt.Sprintf(LeaderboardByIdQuery, "123"),
			},
			expectDbError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				DB: tt.DB,
			}
			if tt.expectDbError {
				mock.ExpectQuery(regexp.QuoteMeta(tt.args.query)).WillReturnError(sql.ErrNoRows)
			}
			if !tt.expectDbError {
				mock.ExpectQuery(regexp.QuoteMeta(tt.args.query)).WillReturnRows(tt.wantRows)
			}
			_, gotErr := s.FindAll(tt.args.ctx, tt.args.query)
			//if !reflect.DeepEqual(gotRows, tt.wantRows) {
			//	t.Errorf("FindAll() gotRows = %v, want %v", gotRows, tt.wantRows)
			//}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("FindAll() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
