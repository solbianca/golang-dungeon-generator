package bsptree

import (
	"dgen/core/tiles"
	"dgen/core/utils/random"
	"dgen/dgen/generator"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math"
)

const (
	directionVertical   = 0
	directionHorizontal = 1
)

var (
	Img       *ebiten.Image
	leafs     []*leaf
	rooms     []*rectangle
	corridors []*rectangle
	tileSize  int

	defaultMinTiles = 3
	defaultMinSize  = defaultMinTiles * tileSize
	maxLevel        = 4
	defaultMaxLevel = 4
)

type Config struct {
	Width, Height, TileSize int

	MaxLevel int
}

func Generate(config *Config) *generator.Dungeon {
	leafs = []*leaf{}
	rooms = []*rectangle{}
	corridors = []*rectangle{}

	tileSize = config.TileSize

	maxLevel = config.MaxLevel
	if maxLevel == 0 {
		maxLevel = defaultMaxLevel
	}

	Img = ebiten.NewImage(config.Width, config.Height)

	root := newLeaf(0, 0, config.Width, config.Height, 0)
	root.split()
	for _, l := range leafs {
		createRooms(l)
	}
	createCorridors(root)

	//drawLeafs(Img)
	//drawRoom(Img)
	//drawCorridors(Img)

	tilesCollection := tiles.NewEmptyTileCollection(config.Width/config.TileSize, config.Height/config.TileSize)

	for _, r := range rooms {
		for column := r.column; column < r.column+r.columns; column++ {
			for row := r.row; row < r.row+r.rows; row++ {
				tilesCollection.Set(tiles.NewTile(column, row, tiles.FloorClass))
			}
		}
	}

	for _, c := range corridors {
		for column := c.column; column < c.column+c.columns; column++ {
			for row := c.row; row < c.row+c.rows; row++ {
				tilesCollection.Set(tiles.NewTile(column, row, tiles.FloorClass))
			}
		}
	}

	tilesCollection.FindNeighborsForTiles()

	maxColumn := config.Width/config.TileSize - 1
	maxRow := config.Height/config.TileSize - 1

	for _, tile := range tilesCollection.AsList() {
		if tile.Class != tiles.FloorClass {
			continue
		}

		for _, neighbor := range tile.GetNeighbors() {
			if neighbor.Class == tiles.EmptyClass {
				neighbor.Class = tiles.WallClass
			}
		}

		column, row := tile.GetAddress()

		if column == 0 {
			tile.Class = tiles.WallClass
		}
		if row == 0 {
			tile.Class = tiles.WallClass
		}
		if column == maxColumn {
			tile.Class = tiles.WallClass
		}
		if row == maxRow {
			tile.Class = tiles.WallClass
		}
	}

	return &generator.Dungeon{
		Tiles:    tilesCollection,
		TileSize: config.TileSize,
	}
}

type rectangle struct {
	x, y, width, height        int
	column, row, columns, rows int
}

func newRectangle(x, y, width, height int) *rectangle {
	r := &rectangle{x: x, y: y, width: width, height: height}

	r.column = x / tileSize
	r.row = y / tileSize
	r.columns = r.width / tileSize
	r.rows = r.height / tileSize

	return r
}

type leaf struct {
	x, y, width, height, level int
	column, row, columns, rows int
	direction                  int

	left, right *leaf
}

func newLeaf(x, y, width, height, level int) *leaf {
	l := &leaf{}

	l.direction = -1
	l.level = level
	l.x = x
	l.y = y
	l.width = width
	l.height = height

	l.column = x / tileSize
	l.rows = y / tileSize
	l.columns = l.width / tileSize
	l.rows = l.height / tileSize

	return l
}

func (l *leaf) split() {
	if l.left != nil && l.right != nil {
		return
	}

	if l.level == maxLevel {
		leafs = append(leafs, l)
		return
	}

	direction := random.BetweenInt(0, 1)

	if float64(l.width)/float64(l.height) >= 1.25 {
		direction = directionVertical
	} else if float64(l.height)/float64(l.width) >= 1.25 {
		direction = directionHorizontal
	}
	l.direction = direction

	split := 0
	if direction == directionHorizontal {
		split = l.height / 2
	} else {
		split = l.width / 2
	}

	if split < defaultMinSize {
		return
	}

	nextLevel := l.level + 1

	if direction == directionHorizontal {
		l.left = newLeaf(l.x, l.y, l.width, split, nextLevel)
		l.right = newLeaf(l.x, l.y+split, l.width, l.height-split, nextLevel)
	} else {
		l.left = newLeaf(l.x, l.y, split, l.height, nextLevel)
		l.right = newLeaf(l.x+split, l.y, l.width-split, l.height, nextLevel)
	}

	l.left.split()
	l.right.split()
}

func createRooms(l *leaf) {
	minColumns := int(math.Ceil(float64(l.columns) * 0.5))
	minRows := int(math.Ceil(float64(l.rows) * 0.5))

	maxColumns := random.BetweenInt(int(math.Ceil(float64(l.columns)*random.BetweenFloat(0.7, 1.0))), l.columns)
	if minColumns > maxColumns {
		maxColumns = minColumns
	}
	maxRows := random.BetweenInt(int(math.Ceil(float64(l.rows)*random.BetweenFloat(0.7, 1.0))), l.rows)
	if minRows > maxRows {
		maxRows = minRows
	}

	columns, rows := random.BetweenInt(minColumns, maxColumns), random.BetweenInt(minRows, maxRows)

	room := newRectangle(
		l.x,
		l.y,
		columns*tileSize,
		rows*tileSize,
	)
	rooms = append(rooms, room)
}

func createCorridors(l *leaf) {
	if l.level < maxLevel-1 {
		createCorridors(l.left)
		createCorridors(l.right)
	}

	if l.direction == directionHorizontal {
		corridors = append(corridors, newRectangle(
			l.x+l.width/2,
			l.y+l.height/4,
			tileSize,
			l.height/2,
		))
	} else if l.direction == directionVertical {
		corridors = append(corridors, newRectangle(
			l.x+l.width/4,
			l.y+l.height/2,
			l.width/2,
			tileSize,
		))
	}
}

func drawLeafs(screen *ebiten.Image) {
	for _, l := range leafs {
		color := randomColorWithAlfa(120)

		ebitenutil.DrawRect(screen, float64(l.x), float64(l.y), float64(l.width), float64(l.height), color)
	}
}

func drawRoom(screen *ebiten.Image) {
	for _, r := range rooms {
		color := hexToColorWithAlpha("0000FF", 255)

		for column := 0; column < r.columns; column++ {
			for row := 0; row < r.rows; row++ {
				x, y := (r.column+column)*tileSize, (r.row+row)*tileSize
				ebitenutil.DrawRect(screen, float64(x), float64(y), float64(tileSize), float64(tileSize), color)
			}
		}
	}
}

func drawCorridors(screen *ebiten.Image) {
	for _, c := range corridors {
		color := hexToColorWithAlpha("00FF00", 255)

		for column := 0; column < c.columns; column++ {
			for row := 0; row < c.rows; row++ {
				x, y := (c.column+column)*tileSize, (c.row+row)*tileSize
				ebitenutil.DrawRect(screen, float64(x), float64(y), float64(tileSize), float64(tileSize), color)
			}
		}
	}
}
