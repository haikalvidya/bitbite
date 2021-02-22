package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

var hexOverflow = [6]string{"A", "B", "C", "D", "E", "F"}

func fromDecimal(theDecimal int) (string, int, int) {
	var neg bool
	if theDecimal < 0 {
		theDecimal = theDecimal*-1
		neg = true
	}

	// find the binary
	var hasil string
	if theDecimal < 2 {
		hasil = strconv.Itoa(theDecimal)
	}
	for a := theDecimal; a > 1; {
		sisa := a%2
		a = a/2
		hasil = strconv.Itoa(sisa) + hasil
		if a < 2 {
			hasil = strconv.Itoa(a) + hasil
		}
	}
	binary, _ := strconv.Atoi(hasil)
	if neg {
		binary = binary*-1
	}
	
	// find hexadecimal
	hasil = ""
	if theDecimal == 0 {
		hasil = strconv.Itoa(theDecimal)
	}
	for a := theDecimal; a > 0; {
		sisa := a%16
		a = a/16
		if sisa >= 10 && sisa <= 15 {
			sisa := hexOverflow[sisa-10]
			hasil = sisa + hasil
			fmt.Println("Masuk")
		} else {
			hasil = strconv.Itoa(sisa) + hasil
		}
	}
	hex := hasil 
	if neg {
		hex = "-" + hex
	}

	// find octal
	hasil = ""
	if theDecimal == 0 {
		hasil = strconv.Itoa(theDecimal)
	}
	for a := theDecimal; a > 0; {
		sisa := a%8
		a = a/8
		hasil = strconv.Itoa(sisa) + hasil
	}
	octal, _ := strconv.Atoi(hasil) 
	if neg {
		octal = octal*-1
	}

	return hex, octal, binary
}

func fromHex(theHex int) (int, int, int) {
	var decimal, octal, binary int
	var neg bool
	if theHex < 0 {
		theHex = theHex*-1
		neg = True
	}

	// find decimal
	var hasil string
	for if theHex == 0 {
		hasil = strconv.Itoa(theHex)
	}
	lenHex = len(theHex)

}

func inputSanitazed() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	theInput, _ := reader.ReadString('\n')
	theInput = strings.Replace(theInput, "\n", "", -1)
	theValue, err := strconv.Atoi(theInput)
	return theValue, err
}

func main(){
	fmt.Println("List Converter from : \n1. Decimal\n2. Hex\n3. Octal\n4. Binary\nctrl+c to exit")
	theCase, _ := inputSanitazedToInteger()
	switch theCase {
		case 1:
			fmt.Print("Input decimal number : ")
			number, err := inputSanitazedToInteger()
			if err != nil {
				fmt.Println("The Input must be number")
			} else {
				hex, octal, binary := fromDecimal(number)
				fmt.Println("Hexadecimal :", hex, "\nOctal :", octal, "\nBinary :", binary)
			}
		case 2:
			fmt.Print("Input hex : ")
			var number string
			fmt.Scanln(&number)
			decimal, octal, binary := fromHex(number)
			fmt.Println("Decimal :", hex, "\nOctal :", octal, "\nBinary :", binary)
		// case 3:
		// 	fromOctal()
		// case 4:
		// 	fromBinary()
		default:
			fmt.Println("Tidak ada di pilihan")
	}
}