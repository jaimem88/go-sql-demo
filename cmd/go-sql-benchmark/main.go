// nolint
package main

import (
	"bufio"
	"fmt"
	"os"
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
	infoColor.Println("(2) create address")
	infoColor.Println("(3) get user")
	infoColor.Println("(4) get address")
	infoColor.Println("(5) get all users and addresses")

	infoColor.Println("(0) Exit")

	mainColor.Print("Your choice: ")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch char {
	case '1':
		input := getInputs([]string{"Name", "Email", "Mobile", "Age"})

		user, err := d.CreateUser(&types.User{
			Name:  input[0],
			Email: input[1],
			// Mobile: input[2],
			// Age:    imput[3],
		})
		if err != nil {
			errColor.Printf("failed to create user! %s", err)
			break
		}

		spew.Dump(user)
	case '2':

	case '3':

	case '0':
		return
	default:
		errColor.Println("Invalid Choice, please try again.")
	}
	subMenu(d)
}
func getInputs(expectedFields []string) []string {

	results := make([]string, len(expectedFields))
	for k, val := range expectedFields {
		mainColor.Printf("\n%s: ", val)
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			errColor.Printf("fatal err - failed to read character!!! %s", err)
			os.Exit(1)
		}
		results[k] = string(char)
	}
	return results
}
