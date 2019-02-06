package demo

import "github.com/jaimemartinez88/go-sql-benchmark/pkg/types"

type Store interface {
	InsertUser(*types.User) (*types.User, error)
	InsertAddress(*types.Address) (*types.Address, error)
	GetUser(id string) (*types.User, error)
	GetAddress(id string) (*types.Address, error)

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

func (d *Demo) GetUser(id string) (*types.User, error) {
	return d.store.GetUser(id)
}

func (d *Demo) GetAddress(id string) (*types.Address, error) {
	return d.store.GetAddress(id)
}

func (d *Demo) GetAllUsersAndAddresses() ([]*types.UserAddress, error) {
	return d.store.GetAllUsersAndAddresses()
}
