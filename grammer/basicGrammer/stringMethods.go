package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello world"
	fmt.Println(strings.Contains(a, "ll"))
	fmt.Println(strings.Count(a, "l"))
	fmt.Println(strings.Compare(a, "HELLO WORLD"))
	fmt.Println(strings.HasPrefix(a, "hel"))
	fmt.Println(strings.HasSuffix(a, "rld"))
	fmt.Println(strings.Index(a, "l"))
	fmt.Println(strings.Join([]string{"he", "llo"}, "-"))
	fmt.Println(strings.Repeat(a, 2))
	fmt.Println(strings.Replace(a, "e", "E", -1))
	fmt.Println(strings.Split("a-v-b-v", "-"))
	fmt.Println(strings.ToLower("ABSAD"))
	fmt.Println(strings.ToUpper("asdas"))
	fmt.Println(len("你好"))
}
