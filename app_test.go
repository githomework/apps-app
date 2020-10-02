package app

import (
	"fmt"
	"testing"

	"github.com/buger/jsonparser"
)

func TestParser(t *testing.T) {
	data := []byte(`{"x":"xx","y":"yy"}`)
	m := map[string]string{}
	jsonparser.EachKey(data, func(i int, value []byte, vt jsonparser.ValueType, err error) {
		if err != nil {
			fmt.Println(err)
		}
		switch i {
		case 0:
			m["x"] = string(value)
		case 1:
			m["y"] = string(value)
		}

	}, []string{"x"}, []string{"y"})

	for k, v := range m {
		fmt.Println(k, v)
	}
}
