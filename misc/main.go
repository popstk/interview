package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type Doctor struct {
	Person
	Score int
}

func main() {
	d := Doctor{
		Person: Person{
			Name: "Sam",
			Age: 35,
		},
		Score:  15,
	}

	data, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	fmt.Println(string(data))
}
