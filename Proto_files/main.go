package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Running main")

	elliot := &Person{
		Firstname: "Ellot",
		Lastname:  "KRish",
		Age:       23,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)

	Toni := &Person{}
	erv := proto.Unmarshal(data, Toni)
	if erv != nil {
		log.Fatalln(err)
	}

	fmt.Println(Toni.GetFirstname())
}
