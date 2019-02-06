package sqlx

import (
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type SQLXStore struct {
	dbx *sqlx.DB
}

// New creates a SQLXStore instance
func New(connStr string) (*SQLXStore, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db connection")
	}
	return &SQLXStore{
		dbx: db,
	}, nil
}
func (s *SQLXStore) InsertUser(u *types.User) error {
	return nil
}
func (s *SQLXStore) InsertAddress(*types.Address) error {
	return nil
}
func (s *SQLXStore) GetUser(id string) (*types.User, error) {
	return nil, nil
}
func (s *SQLXStore) GetAddress(id string) (*types.Address, error) {
	return nil, nil
}
func (s *SQLXStore) GetAllUsersAndAddresses() ([]*types.UserAddress, error) {
	return nil, nil
}
