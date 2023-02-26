package resources

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

func createTestSprite() *Sprite {
	img := ebiten.NewImage(32, 32)
	return NewSprite("test", img)
}

func TestNewSprite(t *testing.T) {
	sprite := createTestSprite()

	expected := &Sprite{
		alias:          "test",
		originalWidth:  32,
		originalHeight: 32,
		image:          ebiten.NewImage(32, 32),
		cache:          nil,
	}

	assert.Equal(t, expected, sprite)
}
