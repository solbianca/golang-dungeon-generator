package resources

import (
	"image"

	"github.com/nfnt/resize"
)

func rect(x0, y0, x1, y1 int) image.Rectangle {
	return image.Rect(x0, y0, x1, y1)
}

func resizeTileImage(targetSize uint, targetImage image.Image) image.Image {
	return resize.Resize(targetSize, targetSize, targetImage, resize.Lanczos3)
}

func resizeImage(targetWidth, targetHeight uint, targetImage image.Image) image.Image {
	return resize.Resize(targetWidth, targetHeight, targetImage, resize.Lanczos3)
}
