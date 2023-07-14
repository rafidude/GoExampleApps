package main

import "fmt"

func main() {
	arr := []any{1, "hello", 3.14}

	for _, v := range arr {
		switch v.(type) {
		case int:
			fmt.Println(v.(int))
		case string:
			fmt.Println(v.(string))
		case float64:
			fmt.Println(v.(float64))
		}
	}
}
