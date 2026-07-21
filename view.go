package main

import (
	"fmt"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/pinwheeeel/2048-tui/engine"
)

func (m *model) View() tea.View {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		m.renderScore(),
		m.renderBoard(),
		m.renderFooter(),
	)

	box := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
	return tea.NewView(box)
}

func (m *model) renderBoard() string {
	rows := make([]string, 4)

	for i := range 4 {
		cols := make([]string, 4)

		for j := range 4 {
			v, _ := m.engine.GridAt(i, j)

			newestTile := m.engine.NewestTile()
			isNewTile := newestTile.X == j && newestTile.Y == i
			cols[j] = m.renderTile(v, isNewTile)
		}

		rows[i] = lipgloss.JoinHorizontal(lipgloss.Top, cols...)
	}

	board := lipgloss.JoinVertical(lipgloss.Left, rows...)

	boardStyle := m.styles.board
	return boardStyle.Render(board)
}

func (m *model) renderTile(v int, isNewTile bool) string {
	text := " "
	if v != 0 {
		text = strconv.Itoa(v)
	}

	tileStyle := m.styles.tiles[v]
	if isNewTile {
		tileStyle = toNewTileStyle(tileStyle)
	}
	return tileStyle.Render(text)
}

func (m *model) renderScore() string {
	return fmt.Sprintf("Score: %d", m.engine.Score())
}

func (m *model) renderFooter() string {
	switch m.engine.Status() {
	case engine.Lost:
		return "You lose!"
	case engine.Won:
		return "You win!"
	default:
		return "Move ←↑↓→\tQuit q"
	}
}
