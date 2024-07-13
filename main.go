package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	morestrings "github.com/lich0079/go_proj/mod/morestrings"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
