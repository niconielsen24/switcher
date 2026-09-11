package game

type GameColor string

const (
	Red    GameColor = "red"
	Blue   GameColor = "blue"
	Green  GameColor = "green"
	Yellow GameColor = "yellow"
)

type Tile struct {
	color    GameColor
	position struct {
		x int
		y int
	}
}

func NewTile(color GameColor) *Tile {
	return &Tile{color: color}
}

func (t *Tile) GetColor() GameColor {
	return t.color
}

func (t *Tile) GetPosition() (int, int) {
	return t.position.x, t.position.y
}

func (t *Tile) SetPosition(x, y int) {
	t.position.x = x
	t.position.y = y
}
