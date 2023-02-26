package game

import (
	"dgen/core/input"
	"dgen/dgen/generator"
	"dgen/dgen/generator/bsptree"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
	TileSize     = 16

	MaxLevel = 4
)

type GameObject interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type Game struct {
	dungeon *generator.Dungeon
}

func NewGame() *Game {
	g := &Game{
		dungeon: bsptree.Generate(&bsptree.Config{
			Width:    1024,
			Height:   1024,
			TileSize: TileSize,
			MaxLevel: MaxLevel,
		}),
	}

	return g
}

func (g *Game) Update() error {
	input.Update()

	if input.IsKeyPressed(ebiten.KeySpace) {
		g.dungeon = bsptree.Generate(&bsptree.Config{
			Width:    1024,
			Height:   1024,
			TileSize: TileSize,
			MaxLevel: MaxLevel,
		})
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.dungeon.Draw(screen)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(bsptree.Img, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f, FPS: %0.2f \nMove (Arrows) Zoom (QE)\nSPACE generate dungeon",
		ebiten.ActualTPS(),
		ebiten.ActualFPS(),
	))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
