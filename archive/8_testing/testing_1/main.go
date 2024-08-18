package main

import (
	"errors"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Hello World")
	concatenateStrings("firststring", "secondstring")
	addNumbers(1, 2)
	concatenateNumbers("2")
}

func concatenateStrings(str1 string, str2 string) string {
	return fmt.Sprintf("%s - %s", str1, str2)
}

func addNumbers(int1 int, int2 int) int {
	return int1 + int2
}

func concatenateNumbers(n1 string) error {

	var re = regexp.MustCompile(`^[0-9]+$`)

	if re.MatchString(n1) {
		//fmt.Println("String contains only digits.")
		//return errors.New("String contains only digits.")
		return nil
	} else {
		//fmt.Println("String does not contain only digits.")
		return errors.New("string does not contain only digits")
	}
	//return fmt.Sprintf("%s %s", n1, n2)
}
