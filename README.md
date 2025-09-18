# input-golang
input in golang
Package input provides simple and easy-to-use functions for console input in Go.

Functions

func Int(text string) int
    Int prompts the user to enter an integer and returns the integer value entered by the user.

        Example usage: age := input.Int("Enter your age: ")

func Float(text string) float64
    Float prompts the user to enter a floating-point number and returns the number entered by the user.

        Example usage: weight := input.Float("Enter your weight: ")

func Str(text string) string
    Str prompts the user to enter a string and returns the string entered by the user.

        Example usage: name := input.Str("Enter your name: ")

func Empty()
    Empty prompts the user to enter any character and does not return any value. This function is useful when waiting for user input before exiting a console application.

        Example usage: input.Empty()