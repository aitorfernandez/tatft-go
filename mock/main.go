package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Placeholder struct {
	Name string
}

func Data(w http.ResponseWriter, r *http.Request) {
	foo := &Placeholder{Name: "foo"}
	d, _ := json.Marshal(foo)

	w.Header().Add("Content-Type", "application/json")
	w.Write(d)
}

func main() {
	fmt.Println("tatft")
}
