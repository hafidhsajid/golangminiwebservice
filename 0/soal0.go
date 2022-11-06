package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func palindromecheck(n string) bool {

	var resString string
	for i := len(n); i > 0; i-- {
		resString += string((n)[i-1])
	}
	if n == resString {
		// fmt.Println("true")
		return true

	} else {
		return false
	}
}

func countpalindrome(reader io.Reader) int {
	var awal, akhir int
	var res = 0
	fmt.Println("Masukkan Input")
	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		input := scanner.Text()
		s := strings.Split(input, " ")
		awal, _ = strconv.Atoi(s[0])
		akhir, _ = strconv.Atoi(s[1])

	}
	// fmt.Scanln(&awal, &akhir)
	fmt.Println(awal, akhir)
	for i := awal; i < akhir; i++ {
		if palindromecheck(fmt.Sprintf("%d", i)) {
			res += 1
			// fmt.Println("pal", i)
		}
	}
	fmt.Println(res)
	// if scanner.Err() != nil {
	// 	return scanner.Err()
	// }

	// fmt.Println("Hello", username)
	// return nil
	return res
}

func main() {
	// fmt.Println(countpalindrome())
	countpalindrome(os.Stdin)
}
