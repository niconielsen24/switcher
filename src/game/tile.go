package game

type GameColor string

const (
	Red    GameColor = "red"
	Blue   GameColor = "blue"
	Green  GameColor = "green"
	Yellow GameColor = "yellow"
)

type Tile struct {
	Color    GameColor `json:"color"`
	Position Position  `json:"position"`
}

func NewTile(color GameColor) *Tile {
	return &Tile{Color: color}
}

func (t *Tile) SetPosition(x, y int) {
	t.Position.X = x
	t.Position.Y = y
}
