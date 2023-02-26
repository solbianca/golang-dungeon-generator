package utils

func ConvertAddressToCoordinates(column, row, tileSize int) (float64, float64) {
	x := float64(column * tileSize)
	y := float64(row * tileSize)

	return x, y
}

func ConvertCoordinatesToAddress(x, y, tileSize int) (int, int) {
	column := x / tileSize
	row := y / tileSize

	return column, row
}
