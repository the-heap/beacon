package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompt gets text input from the user and returns it as a string.
func Prompt(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprint(out, question)
	message, _ := reader.ReadString('\n')
	return message
}

// ByDate implements sort.Interface for []Log based
// on the Date field
type ByDate []Log

func (d ByDate) Len() int           { return len(d) }
func (d ByDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDate) Less(i, j int) bool { return d[i].Date > d[j].Date }

// CheckArgs errors if user fails to provider correct num/types of os.Args.
func CheckArgs(typesNeeded []string, args []string, commandName string) {
	// bail early if args passed in doesn't match length of types expected
	if len(args) != len(typesNeeded) {
		fmt.Fprintln(out, "\n"+"You did not provide the correct arguments to 'beacon "+commandName+"'")
		fmt.Fprintln(out, "You can type 'beacon --help' for more information on how to use Beacon")
		os.Exit(1)
	}

	// TODO: iterate over `typesNeeded` and check each arg matches each typeneeded

}

// ToStringCutNewLine converts a byte to a string and removes the newline
// - This is used for converting Output() values from os/exec that have a newline,
// and need to be a string to write to a file.
func ToStringCutNewLine(byteArr []byte) string {
	return strings.TrimSpace(string(byteArr))
}
