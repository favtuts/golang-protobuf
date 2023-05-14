package main

import (
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	person := &Person{
		Firstname: "John",
		Lastname:  "Doe",
	}

	serializedPerson, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshalling error: ", err)
	}

	ioutil.WriteFile("person.data", serializedPerson, 0644)
}
