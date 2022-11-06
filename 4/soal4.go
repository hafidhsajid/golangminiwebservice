package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	// "github.com/thedevsaddam/renderer"
)

// var rnd *renderer.Render

var data = []string{"apple", "grape", "banana", "melon"}

// var data
// {"W001", "wick", 22},
// {"B001", "bourne", 23},
// {"B002", "bond", 23},]

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// var _ id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			fmt.Println(each)
			result, err = json.Marshal(each)
			// if each.ID == id {
			// 	result, err = json.Marshal(each)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(result)
			// 	return
			// }
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "text/html")
	const html = `
	<!DOCTYPE html>
<html>
<body>

<h2>Palindrom</h2>

<form action="/sent" method="POST">
  <label for="fname">Angka Minimal :</label><br>
  <input type="text" id="min" name="min" value=""><br>
  <label for="lname">Angka Maksimal:</label><br>
  <input type="text" id="max" name="max" value=""><br><br>
  <input type="submit" value="Submit">
</form> 

<p>If you click the "Submit" button, the form-data will be sent to a result many palindrome ".</p>

</body>
</html>`

	if r.Method == "GET" {
		// w.Write([]byte("Hello World!"))
		w.Write([]byte(html))
		// rnd.HTML(w, http.StatusOK, "home", nil)

	}
	if r.Method == "POST" {
		min := r.FormValue("min")
		max := r.FormValue("max")
		// address := r.FormValue("address")
		fmt.Println(min, max)
		w.Write([]byte(min + "data" + max))
		// w.Write([]byte("Hello World!"))
		// rnd.HTML(w, http.StatusOK, "home", nil)

	}
}
func sent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// w.Header().Set("Content Type", "text/html")
	if r.Method == "POST" {

		var _res int
		min := r.FormValue("min")
		max := r.FormValue("max")
		_res = calculate(min, max)
		var data = map[string]interface{}{
			"Title": "Soal1 Palindrom",
			"min":   min,
			"max":   max,
			"hasil": _res,
		}
		res, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		w.Write(res)
		fmt.Println(min, max)

	}

}

func calculate(awal string, akhir string) int {

	_awal, err := strconv.Atoi(awal)
	_akhir, err := strconv.Atoi(akhir)

	if err != nil {
		fmt.Println(err)

	}

	var res = 0
	for i := _awal; i < _akhir; i++ {
		if palindromecheck(fmt.Sprintf("%d", i)) {
			res += 1
		}
	}
	return res
}

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

func main() {
	http.HandleFunc("/sent", sent)
	http.HandleFunc("/", helloworld)
	http.HandleFunc("/satu", user)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
