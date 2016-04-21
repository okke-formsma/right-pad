package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


func main() {
	http.HandleFunc("/", rightpadHandler)
	http.ListenAndServe(":8080", nil)
}

func rightpadHandler(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("str")
	ch := r.URL.Query().Get("ch")
	if ch == "" {
		ch = " "
	}
	length, _ := strconv.Atoi(r.URL.Query().Get("len")) //length zero in case of errors
	if length > 1024 {
		fmt.Fprintf(w, "{ \"err\": \"Len exceeds 1024 characters. Please contact sales for an enterprise license.\"}")
		return
	}

	str = rightpad(str, length, ch)

	res, err := json.Marshal(struct {
		string `json:"str"`
	}{str})
	if err != nil {
		fmt.Fprintf(w, "{ \"err\": \"unable to encode\" }", str)
	}

	fmt.Fprintf(w, "%s", res)
}

func rightpad(str string, length int, pad string) (string) {
	// This is where the magic happens. The inner sanctum. The inner loop. The padding magic.
	// Will return the original string if the length is larger than the string.
	for len(str) < length {
		str = str + pad
	}
	return str
}
