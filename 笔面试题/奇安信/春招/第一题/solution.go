package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 990 {
		fmt.Printf("%.6f\n", 1)
		return
	}
	var c1, c2 float64
	c1, c2 = 990, 1000
	for i := 1; i < n; i++ {
		c1 *= float64(990 - i)
		c2 *= float64(1000 - i)
	}
	val := fmt.Sprintf("%.12f", 1-c1/c2)
	fmt.Println(val[:8])
}

// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	if n > 990 {
// 		fmt.Printf("%.6f\n", 1)
// 		return
// 	}
// 	var all float64
// 	all = 1000
// 	var result float64
// 	result = 1
// 	for i := 1; i <= n; i++ {
// 		result *= (all - 10) / all
//         //all--
// 	}
// 	val := fmt.Sprintf("%.12f", 1-result)
// 	fmt.Println(val[:8])
// }
