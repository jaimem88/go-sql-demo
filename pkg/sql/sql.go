package sql

import (
	"database/sql"

	"github.com/jaimemartinez88/go-sql-benchmark/pkg/types"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type SQLStore struct {
	db *sql.DB
}

// New creates a GoSQL instance
func New(connStr string) (*SQLStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db connection")
	}
	return &SQLStore{
		db: db,
	}, nil
}

func (s *SQLStore) InsertUser(u *types.User) (*types.User, error) {
	err := s.db.QueryRow(`INSERT INTO demo.user (
		name, email, mobile, age, admin
	) VALUES(
		$1, $2, $3, $4, $5
	) RETURNING id`, u.Name, u.Email, u.Mobile, u.Age, u.Admin).Scan(&u.ID)

	return u, errors.Wrap(err, "failed to insert user")
}

func (s *SQLStore) InsertAddress(addr *types.Address) (*types.Address, error) {
	err := s.db.QueryRow(`INSERT INTO demo.address (
		user_id, street_address, suburb, postcode, state, country
	) VALUES(
		$1, $2, $3, $4, $5, $6
	) RETURNING id`,
		addr.UserID, addr.StreetAddress, addr.Suburb, addr.Postcode, addr.State, addr.Country).Scan(&addr.ID)

	return addr, errors.Wrap(err, "failed to insert address")
}

func (s *SQLStore) GetUserByEmail(email string) (*types.User, error) {
	u := &types.User{}
	err := s.db.QueryRow(`SELECT id, name, email, mobile, age, admin 
	FROM demo.user WHERE email = $1`, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.Mobile, &u.Age, &u.Admin)
	return u, errors.Wrap(err, "failed to get user")
}

func (s *SQLStore) GetAddressByUserID(userID string) (*types.Address, error) {
	addr := &types.Address{}
	err := s.db.QueryRow(`SELECT id, user_id,street_address, suburb, postcode, state, country 
	FROM demo.address WHERE user_id = $1`, userID).
		Scan(&addr.ID, &addr.UserID, &addr.StreetAddress, &addr.Suburb, &addr.Postcode, &addr.State, &addr.Country)
	return addr, errors.Wrap(err, "failed to get address")
}

func (s *SQLStore) GetAllUsersAndAddresses() ([]*types.UserAddress, error) {
	rows, err := s.db.Query(`SELECT u.id, u.name, u.email, u.mobile, u.age, u.admin
		a.street_address, a.suburb, a.postcode, a.state, a.country
		FROM demo.user u JOIN demo.address a ON a.user_id = u.id`,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all users and addresses")
	}
	results := []*types.UserAddress{}

	for rows.Next() {
		uAddr := &types.UserAddress{}
		if err := rows.Scan(&uAddr.User.ID, &uAddr.Name, &uAddr.Email, &uAddr.Mobile, &uAddr.Age, &uAddr.Admin,
			&uAddr.StreetAddress, &uAddr.Suburb, &uAddr.Postcode, &uAddr.State, &uAddr.Country,
		); err != nil {
			return nil, errors.Wrap(err, "failed to get all users and addresses")
		}
		results = append(results, uAddr)
	}
	return results, nil
}
