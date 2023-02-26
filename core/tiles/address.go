package tiles

type Addressed interface {
	GetAddress() (int, int)
}

type Address struct {
	Column, Row int
}

func (a Address) GetAddress() (int, int) {
	return a.Column, a.Row
}

func NewAddress(column int, row int) Address {
	return Address{Column: column, Row: row}
}
