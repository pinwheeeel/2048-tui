package main

import (
	"charm.land/lipgloss/v2"
)

type styles struct {
	board lipgloss.Style
	tiles map[int]lipgloss.Style
}

func initStyles() *styles {
	s := styles{
		board: lipgloss.NewStyle().Margin(1),
		tiles: initTileStyles(),
	}
	return &s
}

func initTileStyles() map[int]lipgloss.Style {
	s := make(map[int]lipgloss.Style, 11)

	genericTileStyle := func() lipgloss.Style {
		return lipgloss.NewStyle().
			Width(6).
			Height(3).
			Align(lipgloss.Center, lipgloss.Center).
			Bold(true)
	}

	bgHexes := [11]string{
		"#EEE4DA", "#EDE0C8", "#F2B179",
		"#F59563", "#F67C5F", "#F65E3B",
		"#EDCF72", "#EDCC61", "#EDC850",
		"#EDC53F", "#EDC22E",
	}

	i := 0
	for n := 2; n <= 2048; n *= 2 {
		bgHex := bgHexes[i]
		i++

		textHex := "#776E65"
		if n >= 8 {
			textHex = "#F9F6F2"
		}

		s[n] = genericTileStyle().
			Foreground(lipgloss.Color(textHex)).
			Background(lipgloss.Color(bgHex))
	}

	s[0] = genericTileStyle().
		Foreground(lipgloss.Color("#CDC1B4")).
		Background(lipgloss.Color("#CDC1B4"))

	return s
}

func toNewTileStyle(s lipgloss.Style) lipgloss.Style {
	return s.Foreground(lipgloss.Color("#CC9E00"))
}
