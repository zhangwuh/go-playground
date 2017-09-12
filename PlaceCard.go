package main

import (
	"fmt"
)

const count = 3

var book [count+1]int
var data [count+1]int
func main()  {
	fmt.Println("-----------")
	dfs(1)
}

func dfs(box int) {
	if (box > count) {
		fmt.Println(data[1:])
		return
	}
	for i := 1; i <= count; i++ {
		if (book[i] == 0) {
			fmt.Printf("put %v to box %v\n", i, box)
			data[box] = i
			book[i] = 1
			dfs(box + 1)
			book[i] = 0
			fmt.Printf("pop %v\n", i)
		}
	}
}

