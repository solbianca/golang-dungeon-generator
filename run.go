package main

import (
	"dgen/core/resources"
	game "dgen/dgen"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Dungeon Generator")
	ebiten.SetScreenClearedEveryFrame(true)

	loadResources()

	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func loadResources() {
	resources.LoadSpriteSheet("dungeon", "resources/roguelike_dungeon_tileset.png", 16, 16)
}
