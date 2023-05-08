package main

import "fmt"

func main() {
	var ans = 0
	var rAns = add(0, 10, ans)
	fmt.Println(rAns)
}
func add(start int, end int, ans int) int {
	if start > end{
		return ans
	}
	return add(start+1,end,ans+start)
}