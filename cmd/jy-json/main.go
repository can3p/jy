package main

import (
	fastjson "github.com/valyala/fastjson"
	"fmt"
	"log"
)

func main() {
	parse_and_print(`{"foo": [123, "bar", null], "kkk": 123.88, "hhh": null}`)
	parse_and_print(`[123, "bar", null]`)
}

func parse_and_print(in string) {
	var p fastjson.Parser
	v, err := p.Parse(in)
	if err != nil {
		log.Fatal(err)
	}

	print_value("", v);
}


func print_value(path_so_far string, v *fastjson.Value) {
	// check object
	if v.Type() == fastjson.TypeObject {
		fmt.Printf("%s={}\n", path_so_far)
		o := v.GetObject()
		o.Visit(func(k []byte, v *fastjson.Value) {
			print_value(append_path(path_so_far, k), v);
		})
		return
	}
	// check array
	if v.Type() == fastjson.TypeArray {
		fmt.Printf("%s=[]\n", path_so_far)
		arr := v.GetArray();
		for idx, arr_member := range arr {
			field := fmt.Sprintf("%d", idx)
			print_value(append_path(path_so_far, []byte(field)), arr_member)
		}
		return
	}
	// check number
	if v.Type() == fastjson.TypeNumber {
		if in, err := v.Int64(); err == nil {
			fmt.Printf("%s=%d\n", path_so_far, in)
		} else if fl, err := v.Float64(); err == nil {
			fmt.Printf("%s=%f\n", path_so_far, fl)
		}

		return
	}
	// check string
	if v.Type() == fastjson.TypeString {
		fmt.Printf("%s=\"%s\"\n", path_so_far, v.GetStringBytes())
		return
	}
	// check null
	if v.Type() == fastjson.TypeNull {
		fmt.Printf("%s=null\n", path_so_far)
		return
	}
	// check bool
	if val, err := v.Bool(); err != nil {
		if val == true {
			fmt.Printf("%s=true\n", path_so_far, val)
		} else {
			fmt.Printf("%s=false\n", path_so_far, val)
		}
		return
	}
}

func append_path(path_so_far string, field []byte) string {
	return fmt.Sprintf("%s.%s", path_so_far, field);
}
