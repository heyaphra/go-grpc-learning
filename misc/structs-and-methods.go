// This was an exploration in implementing 01-protobuf-vs-rest using structs and methods
// Less emphasis on optimal solution, more on learning Go

package main

import (
	"fmt"
	"os"

	proto "github.com/golang/protobuf/proto"
)

type xmlPerson struct{}

type jsonPerson struct{}

type protobufPerson struct{}

type bufPkg struct {
	unserialized *Person
	raw          []byte
}

func (x *xmlPerson) size() int64 {
	file, err := os.Stat("./payload-sizes/data/person.xml")
	check(err)
	return file.Size()
}

func (x *jsonPerson) size() int64 {
	file, err := os.Stat("./payload-sizes/data/person.json")
	check(err)
	return file.Size()
}

func (p *protobufPerson) size(buf []byte) int {
	return len(buf)
}

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	xml := xmlPerson{}
	json := jsonPerson{}
	buf := protobufPerson{}
	natalie := &Person{
		Name: "Natalie",
		Age:  25,
	}
	fmt.Println("XML size:", xml.size(), "bytes")
	fmt.Println("JSON size:", json.size(), "bytes")
	fmt.Println("Protobuf size:", buf.size(buf.data(natalie).raw), "bytes")
}
