package main

import (
	"log"
	"fmt"
	"hellogoPkgs.com/utils"
	"github.com/sethvargo/go-password/password"
)

var (
	fName, lName string
	passMaxLenght, passMaxDigits, passMaxSymb int
	upperLett bool
)

func main() {
	fmt.Println("Welcome, What is your name? ")
	fmt.Println("First Name: ")
	fmt.Scanln(&fName)
	fmt.Println("Last Name: ")
	fmt.Scanln(&lName)
	fmt.Println(utils.StringConc(fName, lName))

	fmt.Println("Now enter a Max lenght of characters for a password: ")
	fmt.Scanln(&passMaxLenght)
	fmt.Println("Now enter a Max number of digits for a password: ")
	fmt.Scanln(&passMaxDigits)
	fmt.Println("Now enter a Max number of symbols for a password: ")
	fmt.Scanln(&passMaxSymb)
	fmt.Println("Do you want to add Upper Letter? (true or false): ")
	fmt.Scanln(&upperLett)

	res, err := password.Generate(passMaxLenght, passMaxDigits, passMaxSymb, upperLett, false)
  if err != nil {
    log.Fatal(err)
  }
	fmt.Println("Generating Password...")
  fmt.Println("Password Generated: " + res)

}
