package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	pressedKeys               = pressedKeysCollection{keys: map[ebiten.Key]ebiten.Key{}}
	isLeftMouseButtonPressed  bool
	isRightMouseButtonPressed bool
)

func IsRightMouseButtonPressed() bool {
	return isRightMouseButtonPressed
}

func IsLeftMouseButtonPressed() bool {
	return isLeftMouseButtonPressed
}

func Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && isLeftMouseButtonPressed == false {
		isLeftMouseButtonPressed = true
		return
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && isRightMouseButtonPressed == false {
		isRightMouseButtonPressed = true
		return
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && isLeftMouseButtonPressed {
		isLeftMouseButtonPressed = false
		return
	} else if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && isRightMouseButtonPressed {
		isRightMouseButtonPressed = false
		return
	}
}

func IsKeyPressed(key ebiten.Key) bool {
	if ebiten.IsKeyPressed(key) && !pressedKeys.has(key) {
		pressedKeys.set(key)

		return true
	}

	if inpututil.IsKeyJustReleased(key) {
		pressedKeys.remove(key)
	}

	return false
}

type pressedKeysCollection struct {
	keys map[ebiten.Key]ebiten.Key
}

func (p *pressedKeysCollection) has(key ebiten.Key) bool {
	if _, ok := p.keys[key]; ok {
		return true
	}

	return false
}

func (p *pressedKeysCollection) set(key ebiten.Key) {
	p.keys[key] = key
}

func (p *pressedKeysCollection) remove(key ebiten.Key) {
	delete(p.keys, key)
}
