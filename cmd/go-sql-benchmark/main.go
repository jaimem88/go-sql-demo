// nolint
package main

import (
	"bufio"
	dsql "database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/demo"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/gopg"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/sql"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/sqlx"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/types"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	prettytable "github.com/tatsushid/go-prettytable"
)

var (
	userCols    = []string{"ID", "Name", "Email", "Mobile", "Age", "Admin"}
	addressCols = []string{"ID", "UserID", "Street Address", "Suburb", "Postcode", "State", "Country"}

	mainColor    = color.New(color.FgCyan)
	infoColor    = color.New(color.FgHiBlue)
	successColor = color.New(color.FgHiGreen)
	errColor     = color.New(color.FgRed)
	s            = spinner.New(spinner.CharSets[9], 100*time.Millisecond)

	dbUser    = os.Getenv("DB_USER")
	dbPass    = os.Getenv("DB_PASS")
	dbHost    = os.Getenv("DB_HOST")
	dbPort    = os.Getenv("DB_PORT")
	dbName    = os.Getenv("DB_NAME")
	dbConnStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?application_name=%s&sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName, "demo")
)

func main() {
	menu()

}

// nolint: gocyclo
func menu() {
	fmt.Println("\033[2J")
	successColor.Println("Serversiders SQL DEMO!")

	mainColor.Println("What would you like to do:")

	infoColor.Println("(1) database/sql")
	infoColor.Println("(2) sqlx")
	infoColor.Println("(3) Go PG ORM")

	infoColor.Println("(0) Exit")

	mainColor.Print("Your choice: ")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var store demo.Store
	switch char {
	case '1':
		sqlStore, err := sql.New(dbConnStr)
		if err != nil {
			errColor.Printf("failed to open sql DB connection! %s", err)
			break
		}
		store = sqlStore

	case '2':
		sqlxStore, err := sqlx.New(dbConnStr)
		if err != nil {
			errColor.Printf("failed to open sqlx DB connection! %s", err)
			break
		}
		store = sqlxStore
	case '3':
		gp, err := gopg.New(dbUser, dbPass, dbName, dbHost, dbPort)
		if err != nil {
			errColor.Printf("failed to open gopg DB connection! %s", err)
			break
		}
		store = gp
	case '0':
		return
	default:
		errColor.Println("Invalid Choice, please try again.")
	}

	d := demo.New(store)

	subMenu(d)
	menu()
}

func subMenu(d *demo.Demo) {

	mainColor.Println("What would you like to do:")

	infoColor.Println("(1) create user")
	infoColor.Println("(2) get user")
	infoColor.Println("(3) get all users and addresses")

	infoColor.Println("(0) Exit")

	mainColor.Print("Your choice: ")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\033[2J")
	switch char {
	case '1':
		user, err := user(d)
		if err != nil {
			errColor.Printf("failed to create user! %s\n", err)
			break
		}
		addr, err := address(d, user.ID)
		if err != nil {
			errColor.Printf("failed to create addr! %s\n", err)
			break
		}

		successColor.Println("New user:")
		if err := tableUserAddress(append(userCols, addressCols...), []*types.UserAddress{{User: user, Address: addr}}); err != nil {
			errColor.Printf("failed to print table %s\n", err)
			break
		}
	case '2':
		input := getInputs([]string{"Email"})
		user, err := d.GetUser(input[0])
		if err != nil {
			errColor.Printf("failed to get user! %s\n", err)
			break
		}
		addr, err := d.GetAddress(user.ID.String())
		if err != nil {
			errColor.Printf("failed to get address! %s\n", err)
			break
		}

		successColor.Printf("User: %s\n", input[0])
		if err := tableUserAddress(append(userCols, addressCols...), []*types.UserAddress{{User: user, Address: addr}}); err != nil {
			errColor.Printf("failed to print table %s\n", err)
			break
		}

	case '3':
		usersAndAddr, err := d.GetAllUsersAndAddresses()
		if err != nil {
			errColor.Printf("failed to get users! %s\n", err)
			break
		}
		successColor.Println("All users:")
		if err := tableUserAddress(append(userCols, addressCols...), usersAndAddr); err != nil {
			errColor.Printf("failed to print table %s\n", err)
			break
		}
	case '0':
		return
	default:
		errColor.Println("Invalid Choice, please try again.")
	}
	subMenu(d)
}

// dirty and hacky way to read and return input
func getInputs(expectedFields []string) []string {

	results := make([]string, len(expectedFields))
	for k, val := range expectedFields {
		mainColor.Printf("\n%s: ", val)
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			errColor.Printf("fatal err - failed to read line!!! %s\n", err)
			os.Exit(1)
		}
		results[k] = string(line)
	}
	return results
}

func tableUserAddress(colNames []string, vals []*types.UserAddress) error {
	cols := make([]prettytable.Column, len(colNames))
	for k, v := range colNames {
		cols[k] = prettytable.Column{Header: v, MinWidth: 10, MaxWidth: 18}
	}
	tbl, err := prettytable.NewTable(cols...)
	if err != nil {
		return errors.Wrap(err, "failed to create pretty table :(")
	}
	tbl.Separator = " ╪ "
	for _, v := range vals {
		if err := tbl.AddRow(v.ToSlice()...); err != nil {
			return errors.Wrap(err, "failed to add row pretty table :(")
		}
	}

	if _, err := tbl.Print(); err != nil {
		return errors.Wrap(err, "failed to print pretty table :(")
	}
	return nil
}

func user(d *demo.Demo) (*types.User, error) {
	inputUser := getInputs([]string{"Name", "Email", "Mobile", "Age", "Admin"})

	u := &types.User{
		// not safe, don't do this at home!
		Name:  inputUser[0],
		Email: inputUser[1],
		Admin: false,
	}
	if inputUser[2] != "" {
		u.Mobile = dsql.NullString{String: inputUser[2], Valid: true}
	}
	if inputUser[3] != "" {
		a, err := strconv.Atoi(inputUser[3])
		if err != nil {
			errColor.Printf("hey! %s is not a number! \n", inputUser[3])
			return nil, err
		}
		u.Age = dsql.NullInt64{Int64: int64(a), Valid: true}
	}
	if inputUser[4] == "yes" || inputUser[4] == "true" {
		u.Admin = true
	}

	return d.CreateUser(u)
}

func address(d *demo.Demo, userID uuid.UUID) (*types.Address, error) {
	inputAddr := getInputs([]string{"Street Address", "Suburb", "Postcode", "State", "Country"})

	return d.CreateAddress(&types.Address{
		UserID:        userID,
		StreetAddress: inputAddr[0],
		Suburb:        inputAddr[1],
		Postcode:      inputAddr[2],
		State:         inputAddr[3],
		Country:       inputAddr[4],
	})

}
