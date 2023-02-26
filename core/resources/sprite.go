package resources

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	alias          string
	originalWidth  int
	originalHeight int
	image          *ebiten.Image
	cache          map[string]*ebiten.Image
}

func NewSprite(alias string, image *ebiten.Image) *Sprite {
	width, height := image.Size()

	return &Sprite{
		alias:          alias,
		originalWidth:  width,
		originalHeight: height,
		image:          image,
		cache:          map[string]*ebiten.Image{},
	}
}

func (s *Sprite) Original() *ebiten.Image {
	return s.image
}

func (s *Sprite) Resized(width, height int) *ebiten.Image {
	key := fmt.Sprintf("%d:%d", width, height)
	if img, ok := s.cache[key]; ok == true {
		return img
	}

	if s.originalWidth != width || s.originalHeight != height {
		s.cache[key] = ebiten.NewImageFromImage(resizeImage(uint(width), uint(height), s.image))
	} else {
		s.cache[key] = s.image
	}

	return s.cache[key]
}
