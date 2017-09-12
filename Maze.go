package main

import "fmt"

const n  = 5

var dimension = [5][5]int{ {0,0,0,1,0},
	                       {1,0,0,0,1},
						   {1,0,0,1,0},
	                       {0,1,0,0,0},
						   {0,0,0,0,0} }

var direction = [4][2]int{{0,1},{1,0},{0,-1},{-1,0}}

var dest  = Point{4,2}

var route = []Point{}

var results = [][]Point{}
type Point struct{
	x int
	y int
}

func (p Point) move(x,y int) Point {
	return Point{p.x + x,p.y+y}
}

func main()  {
	sp := Point{0,0}
	route = append(route, sp)
	mazeDfs(sp)
	for _ , element := range results {
		fmt.Println(element)
	}
}

func mazeDfs(p Point) {
	if(p == dest){
		fmt.Printf("re:%v\n",results)
		_route := make([]Point,len(route))
		copy(_route,route)
		results = append(results,_route)
		return
	}
	for i:=0; i<4; i++ {
		element := direction[i]
		next := p.move(element[0],element[1])
		if(next.x < 0 || next.x >= n || next.y <0 || next.y >= n ){
			continue
		}
		if(isInRoute(next) || dimension[next.x][next.y] == 1) {
			fmt.Printf("%v in route or blocked\n",next)
			continue;
		}
		route = append(route, next)
		mazeDfs(next)
		pop(next)
	}
}

func isInRoute(p Point) bool {
	for _, r := range route {
		if(r == p) {
			return true
		}
	}
	return false
}

func pop(p Point) {
	for i, r := range route {
		if(r == p) {
			route = append(route[:i], route[i+1:]...)
			fmt.Printf("pop %v\n",p)
			return
		}
	}
}