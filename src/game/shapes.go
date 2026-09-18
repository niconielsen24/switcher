package game

type Shape = string

type Position struct {
	X, Y int
}

const (
	// Tetrominoes (4 tiles) - the 7 one-sided tetrominoes.
	ShapeSquare  Shape = "square"
	ShapeLine    Shape = "line"
	ShapeL       Shape = "L"
	ShapeMirrorL Shape = "MirrorL"
	ShapeT       Shape = "T"
	ShapeS       Shape = "S"
	ShapeMirrorS Shape = "MirrorS"

	// Pentominoes (5 tiles) - the 18 one-sided pentominoes.
	ShapeC        Shape = "C"
	ShapeCross    Shape = "cross"
	ShapeF        Shape = "F"
	ShapeMirrorF  Shape = "MirrorF"
	ShapeLongLine Shape = "longLine"
	ShapeL5       Shape = "L5"
	ShapeMirrorL5 Shape = "MirrorL5"
	ShapeN        Shape = "N"
	ShapeMirrorN  Shape = "MirrorN"
	ShapeP        Shape = "P"
	ShapeMirrorP  Shape = "MirrorP"
	ShapeT5       Shape = "T5"
	ShapeV        Shape = "V"
	ShapeW        Shape = "W"
	ShapeY        Shape = "Y"
	ShapeMirrorY  Shape = "MirrorY"
	ShapeZ        Shape = "Z"
	ShapeMirrorZ  Shape = "MirrorZ"
)

var ShapeMap = make(map[Shape][]Position)
var ShapeList = []Shape{
	ShapeSquare,
	ShapeLine,
	ShapeL,
	ShapeMirrorL,
	ShapeT,
	ShapeS,
	ShapeMirrorS,
	ShapeC,
	ShapeCross,
	ShapeF,
	ShapeMirrorF,
	ShapeLongLine,
	ShapeL5,
	ShapeMirrorL5,
	ShapeN,
	ShapeMirrorN,
	ShapeP,
	ShapeMirrorP,
	ShapeT5,
	ShapeV,
	ShapeW,
	ShapeY,
	ShapeMirrorY,
	ShapeZ,
	ShapeMirrorZ,
}

func init() {
	ShapeMap[ShapeSquare] = SquarePositions
	ShapeMap[ShapeLine] = LinePositions
	ShapeMap[ShapeL] = LPositions
	ShapeMap[ShapeMirrorL] = MirrorLPositions
	ShapeMap[ShapeT] = TPositions
	ShapeMap[ShapeS] = SPositions
	ShapeMap[ShapeMirrorS] = MirrorSPositions
	ShapeMap[ShapeC] = CPositions
	ShapeMap[ShapeCross] = CrossPositions
	ShapeMap[ShapeF] = FPositions
	ShapeMap[ShapeMirrorF] = MirrorFPositions
	ShapeMap[ShapeLongLine] = LongLinePositions
	ShapeMap[ShapeL5] = L5Positions
	ShapeMap[ShapeMirrorL5] = MirrorL5Positions
	ShapeMap[ShapeN] = NPositions
	ShapeMap[ShapeMirrorN] = MirrorNPositions
	ShapeMap[ShapeP] = PPositions
	ShapeMap[ShapeMirrorP] = MirrorPPositions
	ShapeMap[ShapeT5] = T5Positions
	ShapeMap[ShapeV] = VPositions
	ShapeMap[ShapeW] = WPositions
	ShapeMap[ShapeY] = YPositions
	ShapeMap[ShapeMirrorY] = MirrorYPositions
	ShapeMap[ShapeZ] = ZPositions
	ShapeMap[ShapeMirrorZ] = MirrorZPositions
}

var SquarePositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
}

var LinePositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 3, Y: 0},
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

var FPositions = []Position{
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 2},
}

var MirrorFPositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 2, Y: 1},
	{X: 1, Y: 2},
}

var LongLinePositions = []Position{
	{X: 0, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: 2},
	{X: 0, Y: 3},
	{X: 0, Y: 4},
}

var L5Positions = []Position{
	{X: 0, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: 2},
	{X: 0, Y: 3},
	{X: 1, Y: 3},
}

var MirrorL5Positions = []Position{
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 1, Y: 2},
	{X: 0, Y: 3},
	{X: 1, Y: 3},
}

var NPositions = []Position{
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 0, Y: 2},
	{X: 1, Y: 2},
	{X: 0, Y: 3},
}

var MirrorNPositions = []Position{
	{X: 0, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: 2},
	{X: 1, Y: 2},
	{X: 1, Y: 3},
}

var PPositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 0, Y: 2},
}

var MirrorPPositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 2},
}

var T5Positions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 1, Y: 1},
	{X: 1, Y: 2},
}

var VPositions = []Position{
	{X: 0, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: 2},
	{X: 1, Y: 2},
	{X: 2, Y: 2},
}

var WPositions = []Position{
	{X: 0, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 2},
	{X: 2, Y: 2},
}

var YPositions = []Position{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 2},
	{X: 1, Y: 3},
}

var MirrorYPositions = []Position{
	{X: 0, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 0, Y: 2},
	{X: 0, Y: 3},
}

var ZPositions = []Position{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 1, Y: 2},
	{X: 2, Y: 2},
}

var MirrorZPositions = []Position{
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 1, Y: 1},
	{X: 0, Y: 2},
	{X: 1, Y: 2},
}
