package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func sorting(reader io.Reader) string {
	res := ""
	// res := "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3A13"

	fmt.Println("Masukkan Input")
	scanner := bufio.NewScanner(reader)
	var name string

	if scanner.Scan() {
		name = scanner.Text()
	}
	var arr []string
	var lastword string

	var _6 []string
	var _7 []string
	var _0 []string
	var _9 []string
	var _4 []string
	var _8 []string
	var _1 []string
	var _2 []string
	var _5 []string
	var _3 []string

	for i := 0; i < len(name); i++ {

		if name[i] == ' ' {
			var tmp string
			for j := i - 4; j < i; j++ {
				tmp += string(name[j])
			}
			arr = append(arr, tmp)
		}
	}
	for i := len(name) - 4; i < len(name); i++ {
		lastword += string(name[i])
		// fmt.Println(string(name[i]))
	}
	arr = append(arr, lastword)

	for _, v := range arr {
		switch string(v[0]) {
		case "6":
			_6 = append(_6, v)
		case "7":
			_7 = append(_7, v)
		case "0":
			_0 = append(_0, v)
		case "9":
			_9 = append(_9, v)
		case "4":
			_4 = append(_4, v)
		case "8":
			_8 = append(_8, v)
		case "1":
			_1 = append(_1, v)
		case "2":
			_2 = append(_2, v)
		case "5":
			_5 = append(_5, v)
		case "3":
			_3 = append(_3, v)

		}
	}
	if len(_6) != 0 {
		sort.Slice(_6, func(i, j int) bool {
			return _6[i][2:4] > _6[j][2:4]
		})

		for _, v := range _6 {
			// fmt.Print(v, " ")
			res += v + " "
		}

	}

	if len(_7) != 0 {
		sort.Slice(_7, func(i, j int) bool {
			return _7[i][2:4] > _7[j][2:4]
		})

		for _, v := range _7 {
			// fmt.Print(v, " ")
			res += v + " "
		}

	}

	if len(_0) != 0 {
		sort.Slice(_0, func(i, j int) bool {
			return _0[i][2:4] > _0[j][2:4]
		})

		for _, v := range _0 {
			// fmt.Print(v, " ")
			res += v + " "
		}

	}

	if len(_9) != 0 {
		sort.Slice(_9, func(i, j int) bool {
			return _9[i][2:4] > _9[j][2:4]
		})

		for _, v := range _9 {
			// fmt.Print(v, " ")
			res += v + " "
		}
	}

	if len(_4) != 0 {
		sort.Slice(_4, func(i, j int) bool {
			return _4[i][2:4] > _4[j][2:4]
		})

		for _, v := range _4 {
			// fmt.Print(v, " ")
			res += v + " "
		}

	}

	if len(_8) != 0 {
		sort.Slice(_8, func(i, j int) bool {
			return _8[i][2:4] > _8[j][2:4]
		})

		for _, v := range _8 {
			// fmt.Print(v, " ")
			res += v + " "
		}

	}

	if len(_1) != 0 {
		sort.Slice(_1, func(i, j int) bool {
			return _1[i][2:4] > _1[j][2:4]
		})

		for _, v := range _1 {
			// fmt.Print(v, " ")
			res += v + " "
		}

	}

	if len(_2) != 0 {
		sort.Slice(_2, func(i, j int) bool {
			return _2[i][2:4] > _2[j][2:4]
		})

		for _, v := range _2 {
			// fmt.Print(v, " ")
			res += v + " "
		}

	}

	if len(_5) != 0 {
		sort.Slice(_5, func(i, j int) bool {
			return _5[i][2:4] > _5[j][2:4]
		})

		for _, v := range _5 {
			// fmt.Print(v, " ")
			res += v + " "
		}

	}

	if len(_3) != 0 {
		sort.Slice(_3, func(i, j int) bool {
			return _3[i][2:4] > _3[j][2:4]
		})

		for _, v := range _3 {
			// fmt.Print(v, " ")
			res += v + " "
		}

	}

	// res += strings.Join(_6, "- ")
	// res += strings.Join(_7, "- ")
	// res += strings.Join(_0, "- ")
	// res += strings.Join(_9, "- ")
	// res += strings.Join(_4, "- ")
	// res += strings.Join(_8, "- ")
	// res += strings.Join(_1, "- ")
	// res += strings.Join(_2, "- ")
	// res += strings.Join(_5, "- ")
	// res += strings.Join(_3, "- ")
	fmt.Println(res)
	// fmt.Println(_6, _7, _0, _9, _4, _8, _1, _2, _5, _3)
	return res

}

func main() {
	sorting(os.Stdin)

	// fmt.Println(arr)
}
