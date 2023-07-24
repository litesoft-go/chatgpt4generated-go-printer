package printer

import (
	"fmt"
)

// PrintString - Prints any string
//
//goland:noinspection ALL
func PrintString(s string) {
	fmt.Print(s)
}

// PrintStringln - Prints any string, adding a newLine
//
//goland:noinspection ALL
func PrintStringln(s string) {
	fmt.Println(s)
}

// PrintInt - Print any integer,
//
//goland:noinspection ALL
func PrintInt(i int) {
	fmt.Print(i)
}

// PrintIntln - Print any integer, adding a newLine
//
//goland:noinspection ALL
func PrintIntln(i int) {
	fmt.Println(i)
}
