package engine

import "fmt"

type Direction int

const (
	left Direction = iota
	down
	up
	right
)

func (e *Engine) MoveLeft() (changed bool, err error)  { return e.applyMove(left) }
func (e *Engine) MoveDown() (changed bool, err error)  { return e.applyMove(down) }
func (e *Engine) MoveUp() (changed bool, err error)    { return e.applyMove(up) }
func (e *Engine) MoveRight() (changed bool, err error) { return e.applyMove(right) }

func (e *Engine) applyMove(dir Direction) (changed bool, err error) {
	if e.status != Playing {
		return false, fmt.Errorf("Game is over")
	}

	grid := e.grid
	changed, scoreDelta, bestTile := grid.moveTiles(dir)

	if !changed {
		return false, nil
	}

	e.score += scoreDelta

	grid.spawnRandomTile(e.rng)

	if bestTile == 2048 {
		e.status = Won
	} else if grid.isFilled() && !grid.canMove() {
		e.status = Lost
	}

	return true, nil
}

func (g Grid) canMove() bool {
	for _, dir := range [4]Direction{left, down, up, right} {
		tmp := g
		if canMove, _, _ := tmp.moveTiles(dir); canMove {
			return true
		}
	}

	return false
}

func (grid *Grid) moveTiles(dir Direction) (changed bool, scoreDelta int, bestTile int) {
	changed = false
	scoreDelta = 0
	bestTile = 0

	var js [4]int
	if dir == left || dir == up {
		js = [4]int{0, 1, 2, 3}
	} else {
		js = [4]int{3, 2, 1, 0}
	}

	for i := range 4 {
		k := 0
		for _, j := range js[1:] {
			p := js[k]

			bestTile = max(bestTile, grid[i][p], grid[i][j])

			var u, v *int
			if dir == left || dir == right {
				u, v = &grid[i][p], &grid[i][j]
			} else {
				u, v = &grid[p][i], &grid[j][i]
			}

			// cases: 0 0, 0 x, x 0, x x, x y
			if *v == 0 { // 0 0, x 0
				continue
			}

			if *u == *v { // x x
				*u *= 2
				*v = 0
				k++
				changed = true
				scoreDelta += *u
				bestTile = max(bestTile, *u)
				continue
			}

			if *u != *v { // 0 x, x y
				if *u == 0 { // 0 x
					*u = *v
					*v = 0
					changed = true
					continue
				}

				// x y
				k++
				p = js[k]
				if dir == left || dir == right {
					u = &grid[i][p]
				} else {
					u = &grid[p][i]
				}

				if u == v { // neighbours, noop
					continue
				}

				// treat as 0 x
				*u = *v
				*v = 0
				changed = true
				continue
			}
		}
	}

	return changed, scoreDelta, bestTile
}
