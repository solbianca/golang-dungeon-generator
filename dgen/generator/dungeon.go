package generator

import (
	"dgen/core/resources"
	"dgen/core/tiles"
	"dgen/core/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Dungeon struct {
	Tiles    *tiles.Collection
	TileSize int

	img *ebiten.Image
}

func (d *Dungeon) Update() error {
	return nil
}

func (d *Dungeon) Draw(screen *ebiten.Image) {
	if d.img != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 0)
		screen.DrawImage(d.img, op)

		return
	}

	img := ebiten.NewImageFromImage(screen)

	for _, tile := range d.Tiles.AsList() {
		op := &ebiten.DrawImageOptions{}
		column, row := tile.GetAddress()
		x, y := utils.ConvertAddressToCoordinates(column, row, d.TileSize)
		op.GeoM.Translate(x, y)

		var tileImg *ebiten.Image
		switch tile.Class {
		case tiles.FloorClass:
			tileImg = resources.GetSpriteSheetOrPanic("dungeon").Sprite(0, 0).Original()
		case tiles.WallClass:
			tileImg = resources.GetSpriteSheetOrPanic("dungeon").Sprite(1, 0).Original()
		}

		if tileImg != nil {
			img.DrawImage(tileImg, op)
		}
	}

	d.img = img
}
