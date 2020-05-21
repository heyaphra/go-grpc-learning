package main

import (
	"fmt"
	"os"

	proto "github.com/golang/protobuf/proto"
)

func getFileSize(filepath string) int64 {
	file, err := os.Stat(filepath)
	check(err)
	return file.Size()
}

func getProtobufSize(buf *Person) int {
	data, err := proto.Marshal(buf)
	check(err)
	return len(data)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	natalie := &Person{
		Name: "Natalie",
		Age:  25,
	}
	fmt.Println("XML size:", getFileSize("./payload-sizes/data/person.xml"), "bytes")
	fmt.Println("JSON size:", getFileSize("./payload-sizes/data/person.json"), "bytes")
	fmt.Println("Protobuf size:", getProtobufSize(natalie), "bytes")
}
