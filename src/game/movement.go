package game

import (
	"math/rand"
	"uuid"
)

type Move = string

const (
	MoveOneDiagonal Move = "one_diagonal"
	MoveTwoDiagonal Move = "two_diagonal"
	MoveOneStraight Move = "one_straight"
	MoveTwoStraight Move = "two_straight"
	MoveOneL        Move = "one_l"
	MoveTwoL        Move = "two_l"
	MirrorOneL      Move = "mirror_one_l"
	MirrorTwoL      Move = "mirror_two_l"
)

var Moves = []Move{
	MoveOneDiagonal,
	MoveTwoDiagonal,
	MoveOneStraight,
	MoveTwoStraight,
	MoveOneL,
	MoveTwoL,
	MirrorOneL,
	MirrorTwoL,
}

type MovementCard struct {
	Id   uuid.UUID `json:"id"`
	Move Move      `json:"move"`
}

func NewMovementCard() *MovementCard {
	index := rand.Intn(len(Moves))
	m := Moves[index]
	return &MovementCard{Id: uuid.New(), Move: m}
}

func (mc *MovementCard) ValidSwap(start, end Position) bool {
	switch mc.Move {
	case MoveOneDiagonal:
		return isOneDiagonal(start, end)
	case MoveTwoDiagonal:
		return isTwoDiagonal(start, end)
	case MoveOneStraight:
		return isOneStraight(start, end)
	case MoveTwoStraight:
		return isTwoStraight(start, end)
	case MoveOneL:
		return isOneL(start, end)
	case MoveTwoL:
		return isTwoL(start, end)
	case MirrorOneL:
		return isMirrorOneL(start, end)
	case MirrorTwoL:
		return isMirrorTwoL(start, end)
	default:
		return false
	}
}

func isOneStraight(start, end Position) bool {
	dx, dy := delta(start, end)
	return (dx == 0 && abs(dy) == 1) || (dy == 0 && abs(dx) == 1)
}

func isTwoStraight(start, end Position) bool {
	dx, dy := delta(start, end)
	return (dx == 0 && abs(dy) == 2) || (dy == 0 && abs(dx) == 2)
}

func isOneDiagonal(start, end Position) bool {
	dx, dy := delta(start, end)
	return abs(dx) == 1 && abs(dy) == 1
}

func isTwoDiagonal(start, end Position) bool {
	dx, dy := delta(start, end)
	return abs(dx) == 2 && abs(dy) == 2
}

// oneLOffsets are the 4 rotations of a (1,2) knight move; mirrorLOffsets
// are the 4 rotations of its reflection, i.e. the other knight-move
// chirality. twoLOffsets/mirrorTwoLOffsets are the same two shapes scaled
// by 2, matching how OneStraight/TwoStraight and OneDiagonal/TwoDiagonal
// scale the same direction by 2.
var oneLOffsets = []Position{
	{X: 1, Y: 2},
	{X: -2, Y: 1},
	{X: -1, Y: -2},
	{X: 2, Y: -1},
}

var mirrorOneLOffsets = []Position{
	{X: -1, Y: 2},
	{X: -2, Y: -1},
	{X: 1, Y: -2},
	{X: 2, Y: 1},
}

var twoLOffsets = []Position{
	{X: 2, Y: 4},
	{X: -4, Y: 2},
	{X: -2, Y: -4},
	{X: 4, Y: -2},
}

var mirrorTwoLOffsets = []Position{
	{X: -2, Y: 4},
	{X: -4, Y: -2},
	{X: 2, Y: -4},
	{X: 4, Y: 2},
}

func isOneL(start, end Position) bool {
	return matchesOffset(start, end, oneLOffsets)
}

func isMirrorOneL(start, end Position) bool {
	return matchesOffset(start, end, mirrorOneLOffsets)
}

func isTwoL(start, end Position) bool {
	return matchesOffset(start, end, twoLOffsets)
}

func isMirrorTwoL(start, end Position) bool {
	return matchesOffset(start, end, mirrorTwoLOffsets)
}

func matchesOffset(start, end Position, offsets []Position) bool {
	dx, dy := delta(start, end)
	for _, o := range offsets {
		if o.X == dx && o.Y == dy {
			return true
		}
	}
	return false
}

func delta(start, end Position) (int, int) {
	return end.X - start.X, end.Y - start.Y
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
