package game

import "testing"

func offsetsFrom(start Position, offsets []Position) []Position {
	ends := make([]Position, len(offsets))
	for i, o := range offsets {
		ends[i] = Position{X: start.X + o.X, Y: start.Y + o.Y}
	}
	return ends
}

func testValidSwap(t *testing.T, move Move, start Position, validEnds, invalidEnds []Position) {
	t.Helper()
	card := &MovementCard{Move: move}

	for _, end := range validEnds {
		if !card.ValidSwap(start, end) {
			t.Errorf("%s: expected %v -> %v to be valid, got false", move, start, end)
		}
	}
	for _, end := range invalidEnds {
		if card.ValidSwap(start, end) {
			t.Errorf("%s: expected %v -> %v to be invalid, got true", move, start, end)
		}
	}
}

func TestValidSwapOneStraight(t *testing.T) {
	start := Position{X: 0, Y: 0}
	testValidSwap(t, MoveOneStraight, start,
		[]Position{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}},
		[]Position{{X: 1, Y: 1}, {X: 2, Y: 0}, {X: 0, Y: 0}, {X: 2, Y: 2}},
	)
}

func TestValidSwapTwoStraight(t *testing.T) {
	start := Position{X: 2, Y: 2}
	testValidSwap(t, MoveTwoStraight, start,
		[]Position{{X: 4, Y: 2}, {X: 0, Y: 2}, {X: 2, Y: 4}, {X: 2, Y: 0}},
		[]Position{{X: 3, Y: 2}, {X: 4, Y: 4}, {X: 2, Y: 2}, {X: 1, Y: 1}},
	)
}

func TestValidSwapOneDiagonal(t *testing.T) {
	start := Position{X: 2, Y: 2}
	testValidSwap(t, MoveOneDiagonal, start,
		[]Position{{X: 3, Y: 3}, {X: 3, Y: 1}, {X: 1, Y: 3}, {X: 1, Y: 1}},
		[]Position{{X: 2, Y: 3}, {X: 4, Y: 4}, {X: 2, Y: 2}},
	)
}

func TestValidSwapTwoDiagonal(t *testing.T) {
	start := Position{X: 2, Y: 2}
	testValidSwap(t, MoveTwoDiagonal, start,
		[]Position{{X: 4, Y: 4}, {X: 4, Y: 0}, {X: 0, Y: 4}, {X: 0, Y: 0}},
		[]Position{{X: 3, Y: 3}, {X: 4, Y: 2}, {X: 2, Y: 2}},
	)
}

func TestValidSwapOneL(t *testing.T) {
	start := Position{X: 3, Y: 3}
	testValidSwap(t, MoveOneL, start,
		offsetsFrom(start, oneLOffsets),
		append(offsetsFrom(start, mirrorOneLOffsets), start, Position{X: 4, Y: 4}),
	)
}

func TestValidSwapMirrorOneL(t *testing.T) {
	start := Position{X: 3, Y: 3}
	testValidSwap(t, MirrorOneL, start,
		offsetsFrom(start, mirrorOneLOffsets),
		append(offsetsFrom(start, oneLOffsets), start, Position{X: 4, Y: 4}),
	)
}

func TestValidSwapTwoL(t *testing.T) {
	start := Position{X: 5, Y: 5}
	testValidSwap(t, MoveTwoL, start,
		offsetsFrom(start, twoLOffsets),
		append(offsetsFrom(start, mirrorTwoLOffsets), offsetsFrom(start, oneLOffsets)...),
	)
}

func TestValidSwapMirrorTwoL(t *testing.T) {
	start := Position{X: 5, Y: 5}
	testValidSwap(t, MirrorTwoL, start,
		offsetsFrom(start, mirrorTwoLOffsets),
		append(offsetsFrom(start, twoLOffsets), offsetsFrom(start, mirrorOneLOffsets)...),
	)
}

func TestValidSwapUnknownMove(t *testing.T) {
	card := &MovementCard{Move: "not_a_real_move"}
	if card.ValidSwap(Position{X: 0, Y: 0}, Position{X: 1, Y: 0}) {
		t.Error("expected an unrecognized move to never be valid")
	}
}
