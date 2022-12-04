package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

func TestDAO_InsertOne(t *testing.T) {
	db, mock, _ := sqlmock.New()

	tests := []struct {
		name      string
		DB        *sql.DB
		ctx       context.Context
		exec      string
		wantResp  sql.Result
		wantErr   error
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
			name:      "Sad Path",
			DB:        db,
			ctx:       context.Background(),
			exec:      ``,
			wantErr:   fmt.Errorf("error during leaderboard insert one: %w", errors.New("error")),
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
			} else {
				mock.ExpectExec(tt.exec).WillReturnError(tt.mockErr)
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

	tests := []struct {
		name      string
		ctx       context.Context
		query     string
		DB        *sql.DB
		wantResp  *sqlmock.Rows
		wantErr   bool
		mockErr   error
		expectErr bool
	}{
		{
			name:      "Happy Path",
			DB:        db,
			ctx:       context.Background(),
			query:     ``,
			wantResp:  sqlmock.NewRows([]string{"test_row"}).AddRow("123"),
			expectErr: false,
		},
		{
			name:      "Sad Path",
			DB:        db,
			ctx:       context.Background(),
			query:     ``,
			wantResp:  sqlmock.NewRows([]string{"test_row"}).AddRow("123").RowError(1, errors.New("error")),
			mockErr:   errors.New("error"),
			expectErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				DB: tt.DB,
			}

			mock.ExpectQuery(tt.query).WillReturnRows(tt.wantResp)

			_, err := s.FindAll(tt.ctx, tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
