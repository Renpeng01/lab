package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// now := time.Now()
	// fmt.Println(now.Format("200601021503"))

	// fmt.Println(time.Unix(1667797464, 0).Format("2006-01-02 15:04:05"))
	var a []int
	b := make([]int, 0, 0)

	sa, _ := json.Marshal(a)
	sb, _ := json.Marshal(b)
	fmt.Println(string(sa))
	fmt.Println(string(sb))

}
