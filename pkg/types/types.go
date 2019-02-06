package types

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID     uuid.UUID      `json:"id" db:"id" sql:"id"`
	Name   string         `json:"name" db:"name" sql:"name"`
	Email  string         `json:"email" db:"email" sql:"email"`
	Mobile sql.NullString `json:"mobile" db:"mobile" sql:"mobile"`
	Age    sql.NullInt64  `json:"age" db:"age" sql:"age"`
	Admin  bool           `json:"admin" db:"admin" sql:"admin"`
}

type Address struct {
	ID            uuid.UUID `json:"id" db:"id" sql:"id"`
	UserID        uuid.UUID `json:"user_id" db:"user_id" sql:"user_id"`
	StreetAddress string    `json:"street_addressl" db:"street_address" sql:"street_address"`
	Suburb        string    `json:"suburb" db:"suburb" sql:"suburb"`
	Postcode      string    `json:"postcode" db:"postcode" sql:"postcode"`
	State         string    `json:"state" db:"state" sql:"state"`
	Country       string    `json:"country" db:"country" sql:"country"`
}

type UserAddress struct {
	*User
	*Address
}
