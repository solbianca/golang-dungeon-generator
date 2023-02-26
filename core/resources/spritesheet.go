package resources

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	spritesheet    string
	originalWidth  int
	originalHeight int
	image          *ebiten.Image
	cache          map[string]*Sprite
}

func NewSpritesheet(alias string, originalWidth, originalHeight int, image *ebiten.Image) *SpriteSheet {
	return &SpriteSheet{
		spritesheet:    alias,
		originalWidth:  originalWidth,
		originalHeight: originalHeight,
		image:          image,
		cache:          map[string]*Sprite{},
	}
}

func (s *SpriteSheet) Original() *ebiten.Image {
	return s.image
}

func (s *SpriteSheet) Sprite(column, row int) *Sprite {
	return s.getSprite(column, row, s.originalWidth, s.originalHeight)
}

func (s *SpriteSheet) Resized(column, row, targetWidth, targetHeight int) *Sprite {
	return s.getSprite(column, row, targetWidth, targetHeight)
}

func (s *SpriteSheet) getSprite(column, row, targetWidth, targetHeight int) *Sprite {
	key := fmt.Sprintf("%d:%d:%d:%d", column, row, targetWidth, targetHeight)

	if sprite, ok := s.cache[key]; ok == true {
		return sprite
	}

	originalWidth := s.originalWidth
	originalHeight := s.originalHeight
	x0 := originalWidth * column
	y0 := originalHeight * row
	x1 := x0 + originalWidth
	y1 := y0 + originalHeight

	subImage := s.image.SubImage(rect(x0, y0, x1, y1))

	if s.originalWidth != targetWidth || s.originalHeight != targetHeight {
		resizedImage := ebiten.NewImageFromImage(resizeImage(uint(targetWidth), uint(targetHeight), subImage))
		s.cache[key] = NewSprite("", resizedImage)
	} else {
		s.cache[key] = NewSprite("", ebiten.NewImageFromImage(subImage))
	}

	return s.cache[key]
}
