package color

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"
)

type RGB struct {
	Red   int
	Green int
	Blue  int
}

// RandomRGB Returns a random RGB struct thar represent color
func RandomRGB() RGB {
	rand.Seed(time.Now().UnixNano())
	Red := rand.Intn(255)
	Green := rand.Intn(255)
	blue := rand.Intn(255)
	c := RGB{Red, Green, blue}
	return c
}

func HexToColor(h string) color.Color {
	u, err := strconv.ParseUint(h, 16, 0)
	if err != nil {
		panic(err)
	}

	return color.RGBA{
		R: uint8(u & 0xff0000 >> 16),
		G: uint8(u & 0xff00 >> 8),
		B: uint8(u & 0xff),
		A: 255,
	}
}

func HexToColorWithAlpha(h string, alpha uint8) color.Color {
	u, err := strconv.ParseUint(h, 16, 0)
	if err != nil {
		panic(err)
	}

	return color.RGBA{
		R: uint8(u & 0xff0000 >> 16),
		G: uint8(u & 0xff00 >> 8),
		B: uint8(u & 0xff),
		A: alpha,
	}
}
