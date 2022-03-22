package main

import (
	F "fmt"
	str "strconv"
	"strings"
)

type Foo struct {
	Number int
	Text   string
}

func main() {
	// What would you like to do?
	// convert base 10 to any base
	// convert base any to base 10
	// add two numbers in any base
	// ask for input for the number that wats to be conerted
	// ask for the base that the number  is to be taken to

	intro()

}

func intro() {

	F.Println("Welcome to the base converter")
	F.Println("What would you like to do?")
	F.Println("1. Convert base 10 to any base")
	F.Println("2. Convert base any to base 10")
	F.Println("3. Add two numbers in any base")
	F.Println("4. Exit")

	// get input from user
	var choice int
	F.Scanln(&choice)

	// switch case to handle the user input
	switch choice {

	case 1:
		F.Println("Please enter your value in base 10")
		var base10 int
		F.Scanln(&base10)
		F.Println("Please enter the base you would like to convert to")
		var base int
		F.Scanln(&base)
		F.Println("Your answer is", baseConverter(base10, base))

	case 2:
		F.Println("Please enter your value in it's base")
		var base10 string
		F.Scanln(&base10)
		F.Println("Please enter the base you would like to convert to")
		var base int
		F.Scanln(&base)
		output, err := str.ParseUint(hexaToInteger(base10), base, 64)
		if err != nil {
			F.Println(err)
			panic(err)
		}
		F.Println("Output is: ", output)

	case 3:
		F.Println("Please enter your first value ")
		var base10 string
		F.Scanln(&base10)
		F.Println("what base is this value in?")
		var base int
		F.Scanln(&base)
		F.Println("Please enter your second value")
		var base10Two string
		F.Scanln(&base10Two)
		F.Println("what base is this value in?")
		var baseTwo int
		F.Scanln(&baseTwo)
		F.Println("what base would you like to convert to?")
		var baseThree int
		F.Scanln(&baseThree)
		output, err := str.ParseUint(hexaToInteger(base10), base, 64)
		if err != nil {
			panic(err)
		}
		outputTwo, err := str.ParseUint(hexaToInteger(base10Two), baseTwo, 64)
		if err != nil {
			panic(err)
		}
		F.Println("Output is: ", int(output+outputTwo))
		F.Println("Output is: ", baseConverter(int(output+outputTwo), baseThree))

		// baseToBase10(2)
	case 4:
		F.Println("Goodbye")
		break
	default:
		F.Println("Please enter a valid choice")
	}
}

// convert base 10 to any base
func baseConverter(number, base int) []string {

	var remainderList = make([]Foo, 0)
	var newList = make([]string, 0)
	var remainder int

	for number > 0 {

		F.Println("number at the top", number)

		remainder = number % base

		F.Println("remainder", remainder)

		switch remainder {
		case 10:
			remainderList = append(remainderList, Foo{Text: "A"})
		case 11:
			remainderList = append(remainderList, Foo{Text: "B"})

		case 12:
			remainderList = append(remainderList, Foo{Text: "C"})

		case 13:
			remainderList = append(remainderList, Foo{Text: "D"})

		case 14:
			remainderList = append(remainderList, Foo{Text: "E"})

		case 15:
			remainderList = append(remainderList, Foo{Text: "F"})

		case 16:
			remainderList = append(remainderList, Foo{Text: "G"})

		case 17:
			remainderList = append(remainderList, Foo{Text: "H"})

		case 18:
			remainderList = append(remainderList, Foo{Text: "I"})

		case 19:
			remainderList = append(remainderList, Foo{Text: "J"})

		case 20:
			remainderList = append(remainderList, Foo{Text: "K"})

		case 21:
			remainderList = append(remainderList, Foo{Text: "L"})

		case 22:
			remainderList = append(remainderList, Foo{Text: "M"})

		case 23:
			remainderList = append(remainderList, Foo{Text: "N"})

		case 24:
			remainderList = append(remainderList, Foo{Text: "O"})

		case 25:
			remainderList = append(remainderList, Foo{Text: "P"})

		case 26:
			remainderList = append(remainderList, Foo{Text: "Q"})

		case 27:
			remainderList = append(remainderList, Foo{Text: "R"})

		case 28:
			remainderList = append(remainderList, Foo{Text: "S"})

		case 29:
			remainderList = append(remainderList, Foo{Text: "T"})

		case 30:
			remainderList = append(remainderList, Foo{Text: "U"})

		case 31:
			remainderList = append(remainderList, Foo{Text: "V"})

		default:
			remainderList = append(remainderList, Foo{Number: remainder})

		}

		number = number / base

	}
	F.Println("remainderList", len(remainderList))
	for j := len(remainderList) - 1; j >= 0; j-- {

		if remainderList[j].Number == 0 && remainderList[j].Text != "" {

			newList = append(newList, string(remainderList[j].Text))

		} else if remainderList[j].Number >= 1 {

			newList = append(newList, str.Itoa(remainderList[j].Number))

		} else if remainderList[j].Number == 0 && remainderList[j].Text == "" {

			newList = append(newList, str.Itoa(remainderList[j].Number))
		}
	}
	F.Println("my own answer", remainderList)
	return newList

}

// convert from any base to base 10
func hexaToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}
