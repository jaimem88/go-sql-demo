package demo

import (
	dbsql "database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/jaimemartinez88/go-sql-demo/pkg/gopg"
	"github.com/jaimemartinez88/go-sql-demo/pkg/sql"
	"github.com/jaimemartinez88/go-sql-demo/pkg/sqlx"
	"github.com/jaimemartinez88/go-sql-demo/pkg/types"
)

var (
	dSQL  *Demo
	dSQLx *Demo
	dGoPG *Demo
)

func BenchmarkMain(b *testing.B) {
	db, err := dbsql.Open("postgres", DBConnStr)
	if err != nil {
		b.Fatalf("failed to open db %v", err)
	}
	cleanAllUsers(b, db)
	defer db.Close()

	sqlStore, err := sql.New(DBConnStr)
	if err != nil {
		b.Fatalf("failed to init store %v", err)
	}
	dSQL = New(sqlStore)
	sqlxStore, err := sqlx.New(DBConnStr)
	if err != nil {
		b.Fatalf("failed to init store %v", err)
	}
	dSQLx = New(sqlxStore)
	gopgStore, err := gopg.New(DBUser, DBPass, DBName, DBHost, DBPort)
	if err != nil {
		b.Fatalf("failed to init store %v", err)
	}
	dGoPG = New(gopgStore)
}
func benchmarkGenData(b *testing.B, d *Demo, n int) {
	for i := 0; i < n; i++ {
		u := &types.User{
			Name:   "name",
			Email:  fmt.Sprintf("%s_email_%d@test.com", b.Name(), time.Now().UnixNano()),
			Mobile: dbsql.NullString{String: "0123456789", Valid: true},
			Age:    dbsql.NullInt64{Int64: 30, Valid: true},
			Admin:  true,
		}
		u, err := d.CreateUser(u)
		if err != nil {
			b.Errorf("failed to create user %s", err)
			return
		}
		a := &types.Address{
			UserID:        u.ID,
			StreetAddress: "sstreet_address",
			Suburb:        "suburb",
			State:         "state",
			Postcode:      "postcode",
			Country:       "country",
		}
		_, err = d.CreateAddress(a)
		if err != nil {
			b.Errorf("failed to create addrs %s", err)
			return
		}
	}

}
func BenchmarkSQLGenData2(b *testing.B)     { benchmarkGenData(b, dSQL, 2) }
func BenchmarkSQLGenData10(b *testing.B)    { benchmarkGenData(b, dSQL, 10) }
func BenchmarkSQLGenData100(b *testing.B)   { benchmarkGenData(b, dSQL, 100) }
func BenchmarkSQLGenData500(b *testing.B)   { benchmarkGenData(b, dSQL, 500) }
func BenchmarkSQLGenData1000(b *testing.B)  { benchmarkGenData(b, dSQL, 1000) }
func BenchmarkSQLGenData10000(b *testing.B) { benchmarkGenData(b, dSQL, 10000) }

func BenchmarkSQLXGenData2(b *testing.B)     { benchmarkGenData(b, dSQLx, 2) }
func BenchmarkSQLXGenData10(b *testing.B)    { benchmarkGenData(b, dSQLx, 10) }
func BenchmarkSQLXGenData100(b *testing.B)   { benchmarkGenData(b, dSQLx, 100) }
func BenchmarkSQLXGenData500(b *testing.B)   { benchmarkGenData(b, dSQLx, 500) }
func BenchmarkSQLXGenData1000(b *testing.B)  { benchmarkGenData(b, dSQLx, 1000) }
func BenchmarkSQLXGenData10000(b *testing.B) { benchmarkGenData(b, dSQLx, 10000) }

func BenchmarkGoPGGenData2(b *testing.B)     { benchmarkGenData(b, dGoPG, 2) }
func BenchmarkGoPGGenData10(b *testing.B)    { benchmarkGenData(b, dGoPG, 10) }
func BenchmarkGoPGGenData100(b *testing.B)   { benchmarkGenData(b, dGoPG, 100) }
func BenchmarkGoPGGenData500(b *testing.B)   { benchmarkGenData(b, dGoPG, 500) }
func BenchmarkGoPGGenData1000(b *testing.B)  { benchmarkGenData(b, dGoPG, 1000) }
func BenchmarkGoPGGenData10000(b *testing.B) { benchmarkGenData(b, dGoPG, 10000) }

func BenchmarkSQLGetAllUsersSQL(b *testing.B) {
	if _, err := dSQL.GetAllUsersAndAddresses(); err != nil {
		b.Errorf("failed to get all users %s", err)
	}
}
func BenchmarkSQLXGetAllUsers(b *testing.B) {
	if _, err := dSQLx.GetAllUsersAndAddresses(); err != nil {
		b.Errorf("failed to get all users %s", err)
	}
}
func BenchmarkGoPGGetAllUsers(b *testing.B) {
	if _, err := dGoPG.GetAllUsersAndAddresses(); err != nil {
		b.Errorf("failed to get all users %s", err)
	}
}

func cleanAllUsers(b *testing.B, db *dbsql.DB) {
	b.Helper()

	_, err := db.Exec(`DELETE FROM demo.user `)
	if err != nil {
		b.Fatalf("failed to clean the DB")
	}
}
