package tiles

import (
	"fmt"
)

type Collection struct {
	columns, rows int
	tileMap       map[Address]*Tile

	tilesListCache []*Tile
}

func NewEmptyTileCollection(columns, rows int) *Collection {
	tileCollection := &Collection{
		columns: columns,
		rows:    rows,
		tileMap: map[Address]*Tile{},
	}

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			tileCollection.Set(NewTile(column, row, EmptyClass))
		}
	}

	tileCollection.FindNeighborsForTiles()

	return tileCollection
}

func (c *Collection) AsList() []*Tile {
	if c.tilesListCache != nil {
		return c.tilesListCache
	}

	tiles := make([]*Tile, 0, len(c.tileMap))

	for column := 0; column < c.columns; column++ {
		for row := 0; row < c.rows; row++ {
			if tile, err := c.Get(column, row); err == nil {
				tiles = append(tiles, tile)
			}
		}
	}

	if c.tilesListCache == nil {
		c.tilesListCache = tiles
	}

	return tiles
}

func (c *Collection) Has(column int, row int) bool {
	if _, ok := c.tileMap[NewAddress(column, row)]; ok {
		return true
	}

	return false
}

func (c *Collection) Get(column, row int) (*Tile, error) {
	if tile, ok := c.tileMap[NewAddress(column, row)]; ok {
		return tile, nil
	}

	return nil, fmt.Errorf("can'tileMap find tile by Address Column [%d] and Row [%d]", column, row)
}

func (c *Collection) Set(tile *Tile) {
	c.tileMap[tile.address] = tile
}

func (c *Collection) FindNeighborsForTiles() {
	for _, tile := range c.tileMap {
		c.findNeighborsForTile(tile)
	}
}

func (c *Collection) findNeighborsForTile(tile *Tile) {
	column, row := tile.GetAddress()
	neighbors := map[Addressed]*Tile{}

	if c.Has(column+1, row) {
		tile, _ := c.Get(column+1, row)
		neighbors[tile] = tile
	}
	if c.Has(column-1, row) {
		tile, _ := c.Get(column-1, row)
		neighbors[tile] = tile
	}
	if c.Has(column, row+1) {
		tile, _ := c.Get(column, row+1)
		neighbors[tile] = tile
	}
	if c.Has(column, row-1) {
		tile, _ := c.Get(column, row-1)
		neighbors[tile] = tile
	}
	if c.Has(column+1, row+1) {
		tile, _ := c.Get(column+1, row+1)
		neighbors[tile] = tile
	}
	if c.Has(column+1, row-1) {
		tile, _ := c.Get(column+1, row-1)
		neighbors[tile] = tile
	}
	if c.Has(column-1, row-1) {
		tile, _ := c.Get(column-1, row-1)
		neighbors[tile] = tile
	}
	if c.Has(column-1, row+1) {
		tile, _ := c.Get(column-1, row+1)
		neighbors[tile] = tile
	}

	tile.neighbors = neighbors
}
