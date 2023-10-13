package inmemory

type Record interface {
	ID() string
	ObjectValueMap() map[string]any
}
