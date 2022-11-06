package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func search(reader io.Reader) int {
	fmt.Println("Masukkan Input")
	scanner := bufio.NewScanner(reader)
	var firtInput int
	var name string
	var res int
	// var revlastInput string
	// var digitlast int
	// var lastInput int
	var err error

	if scanner.Scan() {

		name = scanner.Text()
	}

	//work with bug
	word1 := 0
	digit := 1

	if err != nil {
		fmt.Println(err)
	}

	for i := firtInput; i < len(name); i++ { // for search digit
		if i < len(name)-digit {

			var _satu string
			var _dua string
			var _tiga string
			// var _empat string

			for j := 0; j < digit; j++ {
				_satu += string(name[j])
				_dua += string(name[word1+digit+j])
				_tiga += string(name[word1+digit+j+(1*digit)])
			}
			conv1, err := strconv.Atoi(_satu)
			conv2, err := strconv.Atoi(_dua)
			conv3, err := strconv.Atoi(_tiga)

			if err != nil {
				fmt.Println(err)

			}

			if conv1+1 == conv2 && conv2+1 == conv1+2 && conv2+1 == conv3 {
				if conv1+1 != conv2 {
					break
				}
				if conv2+1 != conv3 {
					break

				}
				break
			} else {
				digit++
			}

		}
	}

	for i := 0; i < len(name); i++ {
		var input1 string
		var input2 string
		var input3 string
		var conv1 int
		var conv2 int
		var conv3 int
		// var digit1 int
		// var digit2 int
		// var digit3 int
		if digit != 0 {
			for j := 0; j < digit; j++ {
				input1 += string(name[(i*digit)+j])
				input2 += string(name[(i*digit)+(digit*1)+j])
				input3 += string(name[(i*digit)+(digit*2)+j])
			}

			conv1, err = strconv.Atoi(string(input1))
			if conv1%10 == 9 {
				// fmt.Println("modulus 9")
				// digit1++
				// digit2++
				// digit3++

			}
			if conv1%10 == 8 {
				// fmt.Println("modulus 8")
				// digit3++

			}
			conv2, err = strconv.Atoi(string(input2))

			if conv2%10 == 9 {
				// fmt.Println("modulus 9")

				// digit1++
				// digit2++
				// digit3++
			}

			conv3, err = strconv.Atoi(string(input3))
			if conv3%10 == 9 {
				// fmt.Println("modulus 9")
				// digit1++
				// digit2++
				// digit3++

			}

			if conv1+1 != conv2 || conv2+1 != conv3 {
				if conv1+1 != conv2 {
					fmt.Println("found1", conv1+1)
					res = conv1 + 1
					return conv1 + 1
					// break
				}
				if conv2+1 != conv3 {
					fmt.Println("found2", conv2+1)
					res = conv2 + 1
					return conv2 + 1
					// break
				}

				break

			}
		}
	}
	return res
}

func main() {
	search(os.Stdin)

}

func getvalue(name string, n int, digit int) int {
	// next := ""
	// next1 := ""
	// _digit := 0
	// for i := 0; i < name; i++ {
	// 	convnext, err := strconv.Atoi(next)
	// 	convnext1, err := strconv.Atoi(next1)
	// 	next += string(name[n+i])
	// 	next1 += string(name[n+digit+i])
	// 	fmt.Println(next,next1)
	// }

	return n
}
