package main

import (
	fastjson "github.com/valyala/fastjson"
	"fmt"
	"log"
)

func main() {
	s := `{"foo": [123, "bar"]}`
	var p fastjson.Parser
	v, err := p.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("foo.1=%s\n", v.GetStringBytes("foo", "1"))
}
