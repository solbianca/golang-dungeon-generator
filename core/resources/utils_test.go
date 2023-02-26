package resources

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestImage() image.Image {
	return image.NewRGBA(image.Rect(0, 0, 1000, 1000))
}

func TestRect(t *testing.T) {
	rect := rect(0, 0, 10, 10)

	assert.Equal(t, 0, rect.Min.X)
	assert.Equal(t, 0, rect.Min.Y)
	assert.Equal(t, 10, rect.Max.X)
	assert.Equal(t, 10, rect.Max.Y)
}

func TestResizeImage(t *testing.T) {
	img := createTestImage()

	assert.Equal(t, 0, img.Bounds().Min.X)
	assert.Equal(t, 0, img.Bounds().Min.Y)
	assert.Equal(t, 1000, img.Bounds().Max.X)
	assert.Equal(t, 1000, img.Bounds().Max.Y)

	resizedImage := resizeTileImage(16, img)

	assert.Equal(t, 0, resizedImage.Bounds().Min.X)
	assert.Equal(t, 0, resizedImage.Bounds().Min.Y)
	assert.Equal(t, 16, resizedImage.Bounds().Max.X)
	assert.Equal(t, 16, resizedImage.Bounds().Max.Y)
}
