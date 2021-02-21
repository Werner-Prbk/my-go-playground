package main

import (
	"fmt"

	"github.com/Werner-Prbk/my-go-playground/pkg/aoc"
)

type position struct {
	row int
	col int
}

type seat struct {
	adjacentIndices []position
	state           byte
	nextState       byte
}

const (
	seatOccupied byte = '#'
	seatEmpty    byte = 'L'
	seatIsFloor  byte = '.'
)

func getSeatMap(rows []string) [][]seat {
	var rowCnt = len(rows)
	var colCnt = len(rows[0])

	var sm = make([][]seat, rowCnt)
	for i := range sm {
		sm[i] = make([]seat, colCnt)
	}

	for r, row := range rows {
		for c, value := range row {
			sm[r][c].state = byte(value)
			sm[r][c].nextState = sm[r][c].state
			sm[r][c].adjacentIndices = getAdjacentIndices(r, c, rowCnt, colCnt)
		}
	}

	return sm
}

func getAdjacentIndices(r int, c int, rowCnt int, colCnt int) []position {
	var res = make([]position, 0, 8)

	if c > 0 {
		res = append(res, position{col: c - 1, row: r})
	}
	if c < (colCnt - 1) {
		res = append(res, position{col: c + 1, row: r})
	}
	if r > 0 {
		res = append(res, position{col: c, row: r - 1})

		if c > 0 {
			res = append(res, position{col: c - 1, row: r - 1})
		}
		if c < (colCnt - 1) {
			res = append(res, position{col: c + 1, row: r - 1})
		}
	}
	if r < (rowCnt - 1) {
		res = append(res, position{col: c, row: r + 1})

		if c > 0 {
			res = append(res, position{col: c - 1, row: r + 1})
		}
		if c < (colCnt - 1) {
			res = append(res, position{col: c + 1, row: r + 1})
		}
	}

	return res
}

func applyChanges(sm [][]seat) bool {
	var changed = false
	iterateOverSeatMap(sm, func(s *seat) {
		if s.state != s.nextState {
			changed = true
			s.state = s.nextState
		}
	})
	return changed
}

func countSeatsWithState(sm [][]seat, toCheck []position, state byte) (cnt int) {
	for _, p := range toCheck {
		if sm[p.row][p.col].state == state {
			cnt++
		}
	}
	return cnt
}

func simulateUntilStable(sm [][]seat) (iterations int) {
	for {
		iterateOverSeatMap(sm, func(s *seat) {
			if s.state == seatEmpty {
				if countSeatsWithState(sm, s.adjacentIndices, seatOccupied) == 0 {
					s.nextState = seatOccupied
				}
			} else if s.state == seatOccupied {
				if countSeatsWithState(sm, s.adjacentIndices, seatOccupied) >= 4 {
					s.nextState = seatEmpty
				}
			}
		})

		if applyChanges(sm) == false {
			return iterations
		}
		iterations++
	}
}

func iterateOverSeatMap(sm [][]seat, f func(s *seat)) {
	for ri, row := range sm {
		for ci := range row {
			f(&sm[ri][ci])
		}
	}
}

func main() {
	var lines, err = aoc.ReadTextFileAllLines("input.txt")
	if err != nil {
		fmt.Println("Can not read input. exit")
		return
	}

	var sm = getSeatMap(lines)

	var iterations = simulateUntilStable(sm)

	var occupiedSeats = 0

	iterateOverSeatMap(sm, func(s *seat) {
		if s.state == seatOccupied {
			occupiedSeats++
		}
	})

	fmt.Printf("After %v iterations it got stable with %v occupied seats\n", iterations, occupiedSeats)
}
