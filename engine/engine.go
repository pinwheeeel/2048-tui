package engine

import (
	"math/rand"
)

type Engine struct {
	grid       *Grid
	score      int
	bestTile   int
	newestTile *coordinate
	rng        *rand.Rand
	status     GameStatus
}

type GameStatus int

const (
	Playing GameStatus = iota
	Won
	Lost
)

func New() *Engine {
	return NewWithSeed(rand.Int63())
}

func NewWithSeed(seed int64) *Engine {
	rng := rand.New(rand.NewSource(seed))
	grid, newestTile := initGrid(rng)

	return &Engine{
		grid:       grid,
		score:      0,
		bestTile:   2,
		newestTile: newestTile,
		rng:        rng,
		status:     Playing,
	}
}

func (e *Engine) GridAt(row, col int) (int, error) {
	return e.grid.at(row, col)
}

func (e *Engine) Score() int {
	return e.score
}

func (e *Engine) Status() GameStatus {
	return e.status
}

func (e *Engine) NewestTile() coordinate {
	return *e.newestTile
}
