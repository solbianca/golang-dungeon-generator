package resources

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"image"
	"io/ioutil"
	"strconv"
)

const dpi = 72

var res *storage

func init() {
	res = newResources()
}

type storage struct {
	sprites      map[string]*Sprite
	spriteSheets map[string]*SpriteSheet

	ttf   map[string]*truetype.Font
	faces map[string]font.Face
}

func newResources() *storage {
	return &storage{
		sprites:      map[string]*Sprite{},
		spriteSheets: map[string]*SpriteSheet{},
		ttf:          map[string]*truetype.Font{},
		faces:        map[string]font.Face{},
	}
}

func GetSpriteOrPanic(alias string) *Sprite {
	sprite, err := GetSprite(alias)

	if err == nil {
		return sprite
	}

	panic(err)
}

func GetSprite(alias string) (*Sprite, error) {
	if sprite, ok := res.sprites[alias]; ok == true {
		return sprite, nil
	}

	return nil, fmt.Errorf("sprite by alias [%s] not exists", alias)
}

func GetSpriteSheetOrPanic(alias string) *SpriteSheet {
	spriteSheet, err := GetSpriteSheet(alias)

	if err == nil {
		return spriteSheet
	}

	panic(err)
}

func GetSpriteSheet(alias string) (*SpriteSheet, error) {
	if spriteSheet, ok := res.spriteSheets[alias]; ok == true {
		return spriteSheet, nil
	}

	return nil, fmt.Errorf("spritesheet by alias [%s] not exists", alias)
}

func GetFontOrPanic(alias string, size int) font.Face {
	fontFace, err := GetFont(alias, size)

	if err != nil {
		panic(err)
	}

	return fontFace
}

func GetFont(alias string, size int) (font.Face, error) {
	ttfFont, ok := res.ttf[alias]

	if !ok {
		return nil, fmt.Errorf("font by alias [%s] not exists", alias)
	}

	fontKey := fmt.Sprintf("%s:%s", alias, strconv.Itoa(size))
	fontFace, ok := res.faces[fontKey]

	if ok {
		return fontFace, nil
	}

	fontFace = truetype.NewFace(
		ttfFont, &truetype.Options{
			Size:    float64(size),
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	)

	res.faces[fontKey] = fontFace

	return fontFace, nil
}

func LoadSprite(alias, path string) {
	img, err := loadImage(path)
	if err != nil {
		panic(err)
	}

	res.sprites[alias] = NewSprite(alias, img)
}

func LoadSpriteSheet(alias, path string, originalWidth, originalHeight int) {
	img, err := loadImage(path)
	if err != nil {
		panic(err)
	}

	res.spriteSheets[alias] = NewSpritesheet(alias, originalWidth, originalHeight, img)
}

func LoadFont(alias, path string) {
	if _, ok := res.ttf[alias]; ok {
		panic(fmt.Errorf("font by alias [%s] already loaded", alias))
	}

	fontData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("font not founded by path [%s], %v", path, err))
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		panic(fmt.Errorf("font not parsed by path [%a], %v", path, err))
	}

	res.ttf[alias] = ttfFont
}

func loadImage(path string) (*ebiten.Image, error) {
	file, err := ebitenutil.OpenFile(path)

	if err != nil {
		return nil, fmt.Errorf("can't open file by path [%s], %v", path, err)
	}
	defer func() {
		_ = file.Close()
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("can't decode file by path [%s], %v", path, err)
	}

	return ebiten.NewImageFromImage(img), nil
}
