package demo

import (
	"errors"

	"github.com/jaimemartinez88/go-sql-benchmark/pkg/types"
)

type Store interface {
	InsertUser(*types.User) (*types.User, error)
	InsertAddress(*types.Address) (*types.Address, error)
	GetUserByEmail(email string) (*types.User, error)
	GetAddressByUserID(userID string) (*types.Address, error)

	GetAllUsersAndAddresses() ([]*types.UserAddress, error)
}

type Demo struct {
	store Store
}

func New(store Store) *Demo {
	return &Demo{store}
}

func (d *Demo) CreateUser(u *types.User) (*types.User, error) {
	return d.store.InsertUser(u)
}

func (d *Demo) CreateAddress(u *types.Address) (*types.Address, error) {
	return d.store.InsertAddress(u)
}

func (d *Demo) GetUser(email string) (*types.User, error) {
	return d.store.GetUserByEmail(email)
}

func (d *Demo) GetAddress(userID string) (*types.Address, error) {
	return d.store.GetAddressByUserID(userID)
}

func (d *Demo) GetAllUsersAndAddresses() ([]*types.UserAddress, error) {
	return d.store.GetAllUsersAndAddresses()
}

func (d *Demo) GenData(n int) ([]*types.UserAddress, error) {

	if n >= 100 {
		return nil, errors.New("hey! that's a bit too many")
	}
	ua := make([]*types.UserAddress, n)
	for i := 0; i < n; i++ {
		u := randUser()
		u, err := d.CreateUser(u)
		if err != nil {
			return nil, err
		}
		a := randAddr(u.ID)
		addr, err := d.CreateAddress(a)
		if err != nil {
			return nil, err
		}
		ua[i] = &types.UserAddress{
			User:    u,
			Address: addr,
		}
	}
	return ua, nil

}
