package main

import (
	"fmt"
	"os"
)

func main() {
	v := S("hoge")
	if s, err := ToStringer(v); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Println(s.String())
	}
}


func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, MyError("CastError")
}