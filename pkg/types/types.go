package types

import (
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	TableName struct{}       `sql:"demo.user"`
	ID        uuid.UUID      `json:"id" db:"id" sql:"id"`
	Name      string         `json:"name" db:"name" sql:"name"`
	Email     string         `json:"email" db:"email" sql:"email"`
	Mobile    sql.NullString `json:"mobile" db:"mobile" sql:"mobile"`
	Age       sql.NullInt64  `json:"age" db:"age" sql:"age"`
	Admin     bool           `json:"admin" db:"admin" sql:"admin"`
}

type Address struct {
	TableName     struct{}  `sql:"demo.address"`
	ID            uuid.UUID `json:"id" db:"id" sql:"id"`
	UserID        uuid.UUID `json:"user_id" db:"user_id" sql:"user_id"`
	StreetAddress string    `json:"street_addressl" db:"street_address" sql:"street_address"`
	Suburb        string    `json:"suburb" db:"suburb" sql:"suburb"`
	Postcode      string    `json:"postcode" db:"postcode" sql:"postcode"`
	State         string    `json:"state" db:"state" sql:"state"`
	Country       string    `json:"country" db:"country" sql:"country"`
}

type UserAddress struct {
	TableName struct{} `sql:"demo.user"`
	*User     `db:"user"`
	*Address  `db:"address"`
}

func (u *User) ToSlice() []interface{} {
	return []interface{}{
		u.ID,
		u.Name,
		u.Email,
		u.Mobile.String,
		u.Age.Int64,
		u.Admin,
	}
}

func (a *Address) ToSlice() []interface{} {
	return []interface{}{
		a.ID,
		a.UserID,
		a.StreetAddress,
		a.Suburb,
		a.Postcode,
		a.State,
		a.Country,
	}
}

func (ua *UserAddress) ToSlice() []interface{} {
	return []interface{}{
		ua.User.ID,
		ua.User.Name,
		ua.User.Email,
		ua.User.Mobile.String,
		ua.User.Age.Int64,
		ua.User.Admin,
		ua.Address.ID,
		ua.Address.UserID,
		ua.Address.StreetAddress,
		ua.Address.Suburb,
		ua.Address.Postcode,
		ua.Address.State,
		ua.Address.Country,
	}
}
func UAToSlice(ua []*UserAddress) []interface{} {
	r := make([]interface{}, len(ua))
	for k, v := range ua {
		r[k] = append(v.User.ToSlice(), v.Address.ToSlice()...)
		fmt.Printf("UATOSLICE USER %+v\n", v.User)
		fmt.Printf("UATOSLICE ADDR %+v\n", v.Address)
	}
	fmt.Printf("UATOSLICE %+v\n", r)
	return r
}
