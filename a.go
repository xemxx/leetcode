package main

func main() {

}
func solve(n int) int {
	f := make([]int, n+1)
	f[0], f[1], f[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
