package main

import "encoding/json"

type A struct {
	T string `json:"t"`
}

func main() {
	r, _ := json.Marshal(A{T: "hello"})
	println(string(r))
}
