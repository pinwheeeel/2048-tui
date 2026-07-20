package engine

import "testing"

func TestMoveLeft(t *testing.T) {
	grid := &Grid{
		{2, 0, 0, 0},
		{0, 4, 0, 0},
		{0, 0, 8, 0},
		{0, 0, 0, 16},
	}

	grid.moveTiles(left)
	expected := &Grid{
		{2, 0, 0, 0},
		{4, 0, 0, 0},
		{8, 0, 0, 0},
		{16, 0, 0, 0},
	}

	if !grid.equals(expected) {
		t.Errorf("grid mismatch after moving left")
	}
}

func TestMoveDown(t *testing.T) {
	grid := &Grid{
		{2, 0, 0, 0},
		{0, 4, 0, 0},
		{0, 0, 8, 0},
		{0, 0, 0, 16},
	}

	grid.moveTiles(down)
	expected := &Grid{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 4, 8, 16},
	}

	if !grid.equals(expected) {
		t.Errorf("grid mismatch after moving down")
	}
}

func TestMoveUp(t *testing.T) {
	grid := &Grid{
		{2, 0, 0, 0},
		{0, 4, 0, 0},
		{0, 0, 8, 0},
		{0, 0, 0, 16},
	}

	grid.moveTiles(up)
	expected := &Grid{
		{2, 4, 8, 16},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	if !grid.equals(expected) {
		t.Errorf("grid mismatch after moving up")
	}
}

func TestMoveRight(t *testing.T) {
	grid := &Grid{
		{2, 0, 0, 0},
		{0, 4, 0, 0},
		{0, 0, 8, 0},
		{0, 0, 0, 16},
	}

	grid.moveTiles(right)
	expected := &Grid{
		{0, 0, 0, 2},
		{0, 0, 0, 4},
		{0, 0, 0, 8},
		{0, 0, 0, 16},
	}

	if !grid.equals(expected) {
		t.Errorf("grid mismatch after moving right")
	}
}

func TestNoChange(t *testing.T) {
	testGrid := Grid{
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 2},
	}

	for _, dir := range []Direction{left, down, up, right} {
		grid := testGrid
		if changed, _, bestTile := grid.moveTiles(dir); changed || bestTile != 4 {
			t.Errorf("grid should not change")
		}
		if grid.canMove() {
			t.Errorf("grid should not have legal moves")
		}
	}
}

func TestMerge(t *testing.T) {
	testGrid := Grid{
		{4, 4, 0, 0},
		{4, 0, 4, 0},
		{0, 4, 0, 4},
		{0, 0, 4, 4},
	}

	grid := testGrid
	grid.moveTiles(left)
	expected := &Grid{
		{8, 0, 0, 0},
		{8, 0, 0, 0},
		{8, 0, 0, 0},
		{8, 0, 0, 0},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging left")
	}

	grid = testGrid
	grid.moveTiles(down)
	expected = &Grid{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{8, 8, 8, 8},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging down")
	}

	grid = testGrid
	grid.moveTiles(up)
	expected = &Grid{
		{8, 8, 8, 8},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging up")
	}

	grid = testGrid
	grid.moveTiles(right)
	expected = &Grid{
		{0, 0, 0, 8},
		{0, 0, 0, 8},
		{0, 0, 0, 8},
		{0, 0, 0, 8},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging right")
	}
}

func TestDoubleMerge(t *testing.T) {
	testGrid := Grid{
		{4, 4, 8, 8},
		{4, 0, 0, 8},
		{16, 0, 0, 32},
		{16, 16, 32, 32},
	}

	grid := testGrid
	grid.moveTiles(left)
	expected := &Grid{
		{8, 16, 0, 0},
		{4, 8, 0, 0},
		{16, 32, 0, 0},
		{32, 64, 0, 0},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging left")
	}

	grid = testGrid
	grid.moveTiles(down)
	expected = &Grid{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{8, 4, 8, 16},
		{32, 16, 32, 64},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging down")
	}

	grid = testGrid
	grid.moveTiles(up)
	expected = &Grid{
		{8, 4, 8, 16},
		{32, 16, 32, 64},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging up")
	}

	grid = testGrid
	grid.moveTiles(right)
	expected = &Grid{
		{0, 0, 8, 16},
		{0, 0, 4, 8},
		{0, 0, 16, 32},
		{0, 0, 32, 64},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging right")
	}
}

func TestWeirdMerge(t *testing.T) {
	testGrid := Grid{
		{2, 2, 2, 0},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{0, 2, 2, 2},
	}

	grid := testGrid
	grid.moveTiles(left)
	expected := &Grid{
		{4, 2, 0, 0},
		{4, 4, 0, 0},
		{4, 4, 0, 0},
		{4, 2, 0, 0},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging left")
	}

	grid = testGrid
	grid.moveTiles(down)
	expected = &Grid{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 4, 4, 2},
		{4, 4, 4, 4},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging down")
	}

	grid = testGrid
	grid.moveTiles(up)
	expected = &Grid{
		{4, 4, 4, 4},
		{2, 4, 4, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging up")
	}

	grid = testGrid
	grid.moveTiles(right)
	expected = &Grid{
		{0, 0, 2, 4},
		{0, 0, 4, 4},
		{0, 0, 4, 4},
		{0, 0, 2, 4},
	}
	if !grid.equals(expected) {
		t.Errorf("grid mismatch after merging right")
	}
}

func TestEngineSimulateGoodMove(t *testing.T) {
	e := NewWithSeed(0)
	e.grid = &Grid{
		{0, 0, 32, 32},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	oldScore := e.score
	changed, err := e.MoveLeft()
	if !changed || err != nil {
		t.Errorf("expected a changed grid and no error")
	}
	if e.score != oldScore+64 {
		t.Errorf("score should increase")
	}

	newTiles := 0
	for i := range 4 {
		for j := range 4 {
			if i == 0 && j == 0 {
				continue
			}
			tileValue := e.grid[i][j]
			if tileValue != 0 {
				if tileValue != 2 {
					t.Errorf("new tile should be 2, got %d", tileValue)
				}
				newTiles++
			}
		}
	}
	if newTiles != 1 {
		t.Errorf("expected 1 new tile, got %d", newTiles)
	}
}

func TestEngineSimulateNoopMove(t *testing.T) {
	e := NewWithSeed(0)
	e.grid = &Grid{
		{32, 64, 32, 64},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	oldScore := e.score
	changed, err := e.MoveLeft()
	if changed || err != nil {
		t.Errorf("expected no change and no error")
	}
	if e.score != oldScore {
		t.Errorf("score should not increase")
	}

	newTiles := 0
	for i := range 4 {
		for range 4 {
			if i == 0 {
				continue
			}
		}
	}
	if newTiles != 0 {
		t.Errorf("expected 0 new tiles, got %d", newTiles)
	}
}

func TestEngineSimulateWin(t *testing.T) {
	e := NewWithSeed(0)
	e.grid = &Grid{
		{1024, 1024, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	e.MoveLeft()
	if e.score != 2048 {
		t.Errorf("expected score to be 2048, got %d", e.score)
	}
	if e.status != Won {
		t.Errorf("expected status to be Won, got %d", e.status)
	}

	_, err := e.MoveRight()
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestEngineSimulateLoss(t *testing.T) {
	e := NewWithSeed(0)
	e.grid = &Grid{
		{8, 16, 8, 16},
		{32, 64, 32, 64},
		{8, 16, 8, 16},
		{32, 64, 32, 0},
	}

	if !e.grid.canMove() {
		t.Errorf("expected grid to be moveable, got false")
	}

	_, err := e.MoveRight()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if e.status != Lost {
		t.Errorf("expected status to be Lost, got %d", e.status)
	}

	if e.grid.canMove() {
		t.Errorf("expected grid to be unmoveable, got true")
	}
}
