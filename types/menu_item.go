package types

type MenuItem struct {
	Title    string
	URL      string
	Target   string
	Icon     string
	Sequence int
	IsActive bool
	Children []MenuItem
}
