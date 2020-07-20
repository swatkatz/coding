package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Foo struct {
	Name string
	Price string
}

type Response struct {
	Foos []Foo
	User int64
}
func main() {
	http.HandleFunc("/user/123", func(w http.ResponseWriter, r *http.Request) {
		res := Response{
			Foos: []Foo{
				{
					Name: "swati",
					Price: "100$/hr",
				},
				{
					Name: "whatever",
					Price: "200$/hr",
				},
			},
			User: 123,
		}
		resp, err := json.Marshal(res)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err.Error())
		}
		fmt.Fprintf(w, string(resp))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
