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
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/demo"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/sql"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/types"
)

var (
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
			errColor.Printf("failed to open DB connection! %s", err)
		}
		store = sqlStore

	case '2':

	case '3':

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

		input := getInputs([]string{"Name", "Email", "Mobile", "Age", "Admin"})

		u := &types.User{
			// not safe, don't do this at home!
			Name:  input[0],
			Email: input[1],
			Admin: false,
		}
		if input[2] != "" {
			u.Mobile = dsql.NullString{String: input[2], Valid: true}
		}
		if input[3] != "" {
			a, err := strconv.Atoi(input[3])
			if err != nil {
				errColor.Printf("hey! %s is not a number! \n", input[3])
				break
			}
			u.Age = dsql.NullInt64{Int64: int64(a), Valid: true}
		}
		if input[4] == "yes" || input[4] == "true" {
			u.Admin = true
		}

		user, err := d.CreateUser(u)
		if err != nil {
			errColor.Printf("failed to create user! %s\n", err)
			break
		}
		inputAddr := getInputs([]string{"Street Address", "Suburb", "Postcode", "State", "Country"})

		addr, err := d.CreateAddress(&types.Address{
			UserID:        user.ID,
			StreetAddress: inputAddr[0],
			Suburb:        inputAddr[1],
			Postcode:      inputAddr[2],
			State:         inputAddr[3],
			Country:       inputAddr[4],
		})
		if err != nil {
			errColor.Printf("failed to create user! %s\n", err)
			break
		}

		spew.Dump(user)
		spew.Dump(addr)
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
		spew.Dump(user)
		spew.Dump(addr)

	case '3':

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
