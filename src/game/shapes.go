package game

type Shape = string

type Position struct {
	X, Y int
}

const (
	ShapeSquare  Shape = "square"
	ShapeL       Shape = "L"
	ShapeMirrorL Shape = "MirrorL"
	ShapeT       Shape = "T"
	ShapeS       Shape = "S"
	ShapeMirrorS Shape = "MirrorS"
	ShapeC       Shape = "C"
	ShapeCross   Shape = "cross"
)

var Shapes = make(map[Shape][]Position)

func init() {
	Shapes[ShapeSquare] = SquarePositions
	Shapes[ShapeL] = LPositions
	Shapes[ShapeMirrorL] = MirrorLPositions
	Shapes[ShapeT] = TPositions
	Shapes[ShapeS] = SPositions
	Shapes[ShapeMirrorS] = MirrorSPositions
	Shapes[ShapeC] = CPositions
	Shapes[ShapeCross] = CrossPositions
}

var SquarePositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
}

var LPositions = []Position{
	{X: 0, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: 2},
	{X: 1, Y: 2},
}

var MirrorLPositions = []Position{
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 0, Y: 2},
	{X: 1, Y: 2},
}

var TPositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 1, Y: 1},
}

var SPositions = []Position{
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
}

var MirrorSPositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 2, Y: 1},
}

var CPositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: 2},
	{X: 1, Y: 2},
}

var CrossPositions = []Position{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 2, Y: 1},
	{X: 1, Y: 2},
}
