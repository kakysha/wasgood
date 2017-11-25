package models

// Liquid model
type Liquid struct {
	Item
	Flavors []*Tag `json:"flavors"`
}

// NewLiquid constructor, initializes slice to marshal it into '[]' but not into 'null' coz of nil pointer
func NewLiquid() (i *Liquid) {
	i = &Liquid{Item: *NewItem()}
	i.Flavors = make([]*Tag, 0)
	return
}
