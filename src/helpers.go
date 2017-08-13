package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
)

// Prompt gets text input from the user and returns it as a string.
func Prompt(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	message, _ := reader.ReadString('\n')

	return message
}

// CheckArgs errors if user fails to provider correct num/types of os.Args.
func CheckArgs(typesNeeded []string, args []string, commandName string) {
	// bail early if args passed in doesn't match length of types expected
	if len(args) != len(typesNeeded) {
		fmt.Println("\n" + "You did not provide the correct arguments to 'beacon " + commandName + "'")
		fmt.Println("You can type 'beacon --help' for more information on how to use Beacon")
		os.Exit(1)
	}

	// TODO: iterate over `typesNeeded` and check each arg matches each typeneeded

}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
