package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
)

var hexOverflow = []string{"A", "B", "C", "D", "E", "F"}

func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
} 

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

func fromHex(theHex string) (int, int, int) {
	var decimal, octal, binary int

	// find decimal
	var hasil int
	zeroCheck, err := strconv.Atoi(theHex)
	if zeroCheck == 0 && err == nil {
		hasil = zeroCheck
	}
	lenHex := len(theHex)
	var value,i int
	for ; lenHex>0; lenHex-- {
		charCheck, err := strconv.Atoi(strings.Split(theHex, "")[i])
		if err == nil {
			value = int(math.Pow(16,float64(lenHex-1)))*charCheck
		} else {
			value = int(math.Pow(16,float64(lenHex-1)))*(indexOf(strings.Split(theHex, "")[i], hexOverflow)+10)
		}
		hasil += value
		i += 1
	}
	decimal = hasil

	// find octal
	hasil = 0
	if zeroCheck == 0 && err == nil {
		hasil = zeroCheck
	}
	for a := decimal; a > 0; {
		sisa := a%8
		a = a/8
		x := strconv.Itoa(sisa)
		y :=  strconv.Itoa(hasil)
		if y == "0" && a > 1 {
			hasil, _ = strconv.Atoi(x)
		} else {
			hasil, _ = strconv.Atoi(x + y)
		}
	}
	octal = hasil

	return decimal, octal, binary
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
	theCase, _ := inputSanitazed()
	switch theCase {
		case 1:
			fmt.Print("Input decimal number : ")
			number, err := inputSanitazed()
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
			decimal, octal, binary := fromHex(strings.ToUpper(number))
			fmt.Println("Decimal :", decimal, "\nOctal :", octal, "\nBinary :", binary)
		// case 3:
		// 	fromOctal()
		// case 4:
		// 	fromBinary()
		default:
			fmt.Println("Tidak ada di pilihan")
	}
}