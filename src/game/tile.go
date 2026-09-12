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
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
}

func NewTile(color GameColor) *Tile {
	return &Tile{Color: color}
}

func (t *Tile) GetColor() GameColor {
	return t.Color
}

func (t *Tile) GetPosition() (int, int) {
	return t.Position.X, t.Position.Y
}

func (t *Tile) SetPosition(x, y int) {
	t.Position.X = x
	t.Position.Y = y
}
