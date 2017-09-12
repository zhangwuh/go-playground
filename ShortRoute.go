package main

import "fmt"

var dimension = [4][4]int{ {0,2,6,4},
	{999,0,3,999},
	{7,999,0,1},
	{5,999,12,0},
}

func main() {
	for i:=0;i<4;i++{
		for j:=0;j<4;j++{
			cal(i,j)
		}
	}

	fmt.Println(dimension)
}

func cal(s int,d int)  {
	for i:=0;i<4;i++{
		if (dist(s, i) + dist(i, d) < dist(s,d)) {
			dimension[s][d] = dist(s, i) + dist(i, d)
		}
	}
}

func dist(s int,d int) int{
	if(s == d) {
		return 0
	}
	return dimension[s][d]
}
