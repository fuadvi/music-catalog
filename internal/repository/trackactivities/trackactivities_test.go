package trackactivities

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fuadvi/music-catalog/internal/models/trackactivities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	assert.NoError(t, err)

	type args struct {
		model trackactivities.TrackActivity
	}

	now := time.Now()
	isLiked := true
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				model: trackactivities.TrackActivity{
					Model: gorm.Model{
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					SpotifyID: "spotifyID",
					Isliked:   &isLiked,
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "track_activities" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserID,
						args.model.SpotifyID,
						args.model.Isliked,
						args.model.CreatedBy,
						args.model.UpdatedBy,
					).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				model: trackactivities.TrackActivity{
					Model: gorm.Model{
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					SpotifyID: "spotifyID",
					Isliked:   &isLiked,
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "track_activities" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserID,
						args.model.SpotifyID,
						args.model.Isliked,
						args.model.CreatedBy,
						args.model.UpdatedBy,
					).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &Repository{
				db: gormDB,
			}
			if err := r.Create(context.Background(), tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	assert.NoError(t, err)
	type args struct {
		model trackactivities.TrackActivity
	}

	now := time.Now()
	isLiked := true
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				model: trackactivities.TrackActivity{
					Model: gorm.Model{
						ID:        uint(1),
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					SpotifyID: "spotifyID",
					Isliked:   &isLiked,
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "track_activities" SET (.+) WHERE (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserID,
						args.model.SpotifyID,
						args.model.Isliked,
						args.model.CreatedBy,
						args.model.UpdatedBy,
						args.model.ID,
					).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				model: trackactivities.TrackActivity{
					Model: gorm.Model{
						ID:        uint(1),
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					SpotifyID: "spotifyID",
					Isliked:   &isLiked,
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "track_activities" SET (.+) WHERE (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserID,
						args.model.SpotifyID,
						args.model.Isliked,
						args.model.CreatedBy,
						args.model.UpdatedBy,
						args.model.ID,
					).WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &Repository{
				db: gormDB,
			}
			if err := r.Update(context.Background(), tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
