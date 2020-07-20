package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Transport:     http.DefaultTransport,
		Timeout:       time.Second,
	}
	resp, err := client.Get("http://localhost:8080/user/123")
	if err != nil {
		fmt.Printf("error: %v \n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// unmarshal into a map
	var val map[string]interface{}
	json.Unmarshal(body, &val)
	fmt.Printf("val: %v \n ", val)

	// unmarshal into a struct
	var res Result
	json.Unmarshal(body, &res)
	fmt.Printf("res: %v \n ", res)

	/*i := [][]int{{1, 2}, {3, 4}}
	x, _ := json.Marshal(i)
	fmt.Println(string(x))*/
}

type FooRes struct {
	Name string
	StrPrice string `json:"Price,omitempty"`
}
type Result struct {
	Foos []FooRes
	UserID int64 `json:"User,omitempty"`
}
