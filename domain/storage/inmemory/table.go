package inmemory

type Table struct {
	records   []Record
	currentID uint
}
