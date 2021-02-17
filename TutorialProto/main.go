package main

import (
	"fmt"
	"log"
	"math"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Krish! running main")
	fmt.Println("krish 2")
	fmt.Println(math.Pi * 2)

	elliot := &Person{
		Firstname: "Elliot",
		Lastname:  "Perderson",
		Age:       "23",
	}

	//now we convert data into byte means serialization for saving to file or wire transfer

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Println("err")
	}

	fmt.Println(data)

}
