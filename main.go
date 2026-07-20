package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/pinwheeeel/2048-tui/engine"
)

type model struct {
	width  int
	height int
	styles *styles

	engine *engine.Engine
}

func initialModel() *model {
	return &model{
		engine: engine.New(),
		styles: initStyles(),
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left":
			m.engine.MoveLeft()
		case "down":
			m.engine.MoveDown()
		case "up":
			m.engine.MoveUp()
		case "right":
			m.engine.MoveRight()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
