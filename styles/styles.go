package styles

import "github.com/charmbracelet/lipgloss"

// The margins shouldn't be repeated but inheritance doesn't seem to be fully working
var Unsolved lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "0", Dark: "15"}).
	Bold(true).
	MarginLeft(1).
	MarginRight(1)

var Solved lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#8FBCBB")).
	MarginLeft(1).
	MarginRight(1).
	Strikethrough(true).
	Inherit(Unsolved)

var Category lipgloss.Style = lipgloss.NewStyle().
	Bold(false).
	MarginLeft(1).
	MarginRight(1).
	Inherit(Unsolved)

var Header lipgloss.Style = lipgloss.NewStyle().
	Bold(true).
	MarginLeft(1).
	MarginRight(1).
	Inherit(Unsolved)

var Id lipgloss.Style = lipgloss.NewStyle().
	Bold(false).
	MarginLeft(1).
	MarginRight(1).
	Inherit(Unsolved)
