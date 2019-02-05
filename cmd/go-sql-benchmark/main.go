// nolint
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var (
	mainColor    = color.New(color.FgCyan)
	infoColor    = color.New(color.FgHiBlue)
	successColor = color.New(color.FgHiGreen)
	errColor     = color.New(color.FgRed)
	s            = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
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
	infoColor.Println("(3) squirrel")
	infoColor.Println("(4) Go PG ORM")

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

	case '2':

	case '3':

	case '4':
	case '0':
		return
	default:
		errColor.Println("Invalid Choice, please try again.")
	}
	menu()
}
