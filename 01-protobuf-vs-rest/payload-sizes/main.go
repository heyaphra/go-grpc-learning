package main

import (
	"fmt"
	"os"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	xml := xmlPerson{}
	json := jsonPerson{}
	buf := protobufPerson{}
	natalie := &Person{
		Name: "Natalie",
		Age:  25,
	}
	fmt.Println("XML size:", xml.size())
	fmt.Println("JSON size:", json.size())
	fmt.Println("Protobuf size:", buf.size(buf.data(natalie).raw))
}

type xmlPerson struct{}

func (x *xmlPerson) size() int64 {
	file, err := os.Stat("./payload-sizes/data/person.xml")
	check(err)
	return file.Size()
}

type jsonPerson struct{}

func (x *jsonPerson) size() int64 {
	file, err := os.Stat("./payload-sizes/data/person.json")
	check(err)
	return file.Size()
}

type file interface {
	size() int
}

type bufPkg struct {
	unserialized *Person
	raw          []byte
}

type protobufPerson struct{}

func (p *protobufPerson) data(buf *Person) *bufPkg {
	return &bufPkg{
		unserialized: buf,
		raw: func() []byte {
			data, err := proto.Marshal(buf)
			check(err)
			return data
		}(),
	}
}

func (p *protobufPerson) size(buf []byte) int {
	return len(buf)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
