package manager

type Location struct {
	X int
	Y int
}

func getFreshLocation(y, x int) Location {
	var result Location

	result.X = x
	result.Y = y

	return result
}
