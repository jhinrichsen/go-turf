package transform

// Transform a position p in range a..b into position p' in range x..y
func transform(p int, a int, b int, x int, y int) int {
	return x + (y-x)*(p-a)/(b-a)
}
