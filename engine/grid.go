package engine

import (
	"errors"
	"math/rand"
)

type Grid [4][4]int

func initGrid(rng *rand.Rand) *Grid {
	g := &Grid{}
	g.spawnRandomTile(rng)
	return g
}

func (g *Grid) spawnRandomTile(rng *rand.Rand) (ok bool) {
	available := make([][2]int, 0, 16)

	for i := range 4 {
		for j := range 4 {
			if g[i][j] == 0 {
				available = append(available, [2]int{i, j})
			}
		}
	}

	if len(available) == 0 {
		return false
	}

	idx := available[rng.Intn(len(available))]
	i, j := idx[0], idx[1]
	g[i][j] = 2
	return true
}

func (g *Grid) isFilled() bool {
	for i := range 4 {
		for j := range 4 {
			if g[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (g *Grid) at(row, col int) (int, error) {
	if row < 0 || row >= 4 || col < 0 || col >= 4 {
		return 0, errors.New("Index out of bounds")
	}
	return g[row][col], nil
}
