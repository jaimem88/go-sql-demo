package gopg

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/types"
)

type GoPGStore struct {
	db *pg.DB
}

// New creates a SQLXStore instance
func New(user, pass, dbName, host, port string) (*GoPGStore, error) {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: pass,
		Database: dbName,
		Addr:     fmt.Sprintf("%s:%s", host, port),
	})

	return &GoPGStore{
		db: db,
	}, nil
}

func (s *GoPGStore) InsertUser(u *types.User) (*types.User, error) {
	err := s.db.Insert(u)
	return u, err
}
func (s *GoPGStore) InsertAddress(addr *types.Address) (*types.Address, error) {
	err := s.db.Insert(addr)
	return addr, err
}

func (s *GoPGStore) GetUserByEmail(email string) (*types.User, error) {
	user := &types.User{}
	err := s.db.
		Model(user).
		Where("email = ?", email).
		Select()
	return user, err
}
func (s *GoPGStore) GetAddressByUserID(userID string) (*types.Address, error) {
	addr := &types.Address{}
	err := s.db.
		Model(addr).
		Where("user_id = ?", userID).
		Select()
	return addr, err
}
func (s *GoPGStore) GetAllUsersAndAddresses() ([]*types.UserAddress, error) {

	var userAddr []*types.UserAddress
	err := s.db.
		Model(&userAddr).
		Column("user_address.*", "address.*").
		Join("LEFT JOIN demo.address ").
		JoinOn("address.user_id = \"user_address\".id").
		Select()

	for _, u := range userAddr {
		u.User.ID = u.Address.UserID
	}

	return userAddr, err
}
