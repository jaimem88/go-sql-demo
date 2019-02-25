package gopg

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-pg/pg"
	"github.com/go-test/deep"
	"github.com/jaimemartinez88/go-sql-demo/pkg/types"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")
)

func TestGoPGStore_InsertUser(t *testing.T) {
	s, err := New(dbUser, dbPass, dbName, dbHost, dbPort)
	require.NoError(t, err, "open db connection")
	u := types.User{
		Name:   "test name",
		Email:  "test_pg@name.com",
		Mobile: sql.NullString{String: "0433333312", Valid: true},
		Age:    sql.NullInt64{Int64: 33, Valid: true},
		Admin:  true,
	}
	cleanUser(t, s.db, u.Email)
	tests := []struct {
		name    string
		u       types.User
		want    types.User
		wantErr bool
	}{
		{
			name: "insert_ok",
			u:    u,
			want: u,
		},
		//TODO: add more test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.InsertUser(&tt.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoPGStore.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer cleanUser(t, s.db, got.Email)

			// don't do this at home
			got.ID = uuid.Nil
			diff := deep.Equal(tt.want, *got)
			assert.Nil(t, diff, "want != got")
		})
	}
}

func cleanUser(t *testing.T, db *pg.DB, email string) {
	t.Helper()

	_, err := db.Exec(`DELETE FROM demo.user WHERE email = ?`, email)
	require.NoError(t, err, "clean user")
}
