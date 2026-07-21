package engine

import (
	"testing"
)

func TestEngineConstruction(t *testing.T) {
	e := NewWithSeed(0)

	if e == nil {
		t.Errorf("engine is nil")
	}
	if e.grid == nil {
		t.Errorf("grid is nil")
	}
	if e.score != 0 {
		t.Errorf("score is not 0")
	}
	if e.bestTile != 2 && e.bestTile != 4 {
		t.Errorf("bestTile is not 2 or 4")
	}
	if e.status != Playing {
		t.Errorf("status is not Playing")
	}
}

func TestGridAt(t *testing.T) {
	e := NewWithSeed(0)

	for i := range 4 {
		for j := range 4 {
			gridAt, err := e.GridAt(i, j)
			if err != nil {
				t.Errorf("GridAt(%d, %d) returned error: %v", i, j, err)
			}

			if e.grid[i][j] != gridAt {
				t.Errorf("grid[%d][%d] != GridAt(%d, %d)", i, j, i, j)
			}
		}
	}
}

func TestGridOutOfBounds(t *testing.T) {
	e := NewWithSeed(0)

	a := [][2]int{{-1, 0}, {0, -1}, {4, 0}, {0, 4}}

	for _, coords := range a {
		row, col := coords[0], coords[1]
		_, err := e.GridAt(row, col)
		if err == nil {
			t.Errorf("GridAt(%d, %d) did not return error", row, col)
		}
	}
}

func TestNewGame(t *testing.T) {
	e := NewWithSeed(0)

	startingTiles := 0
	for i := range 4 {
		for j := range 4 {
			tileValue := e.grid[i][j]
			if tileValue != 0 {
				startingTiles++

				if tileValue != 2 && tileValue != 4 {
					t.Errorf("tile value is not 2 or 4")
				}
			}
		}
	}

	if startingTiles != 1 {
		t.Errorf("must start with exactly one tile")
	}
}

func TestSeedReproducibility(t *testing.T) {
	for seed := range 16 {
		seed := int64(seed)
		e1 := NewWithSeed(seed)
		e2 := NewWithSeed(seed)

		for i := range 4 {
			for j := range 4 {
				if e1.grid[i][j] != e2.grid[i][j] {
					t.Errorf("seed %d: grid mismatch at (%d, %d)", seed, i, j)
				}
			}
		}
	}
}
