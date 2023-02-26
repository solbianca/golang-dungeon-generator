package tiles

const (
	EmptyClass = "empty"
	FloorClass = "floor"
	WallClass  = "wall"
)

type Tile struct {
	address   Address
	neighbors map[Addressed]*Tile

	Class string
}

func NewTile(column, row int, class string) *Tile {
	return &Tile{
		address:   NewAddress(column, row),
		neighbors: map[Addressed]*Tile{},
		Class:     class,
	}
}

func (t *Tile) GetNeighbors() map[Addressed]*Tile {
	return t.neighbors
}

func (t *Tile) GetAddress() (column, row int) {
	return t.address.GetAddress()
}
