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
	Strikethrough(true).
	MarginLeft(1).
	MarginRight(1).
	Inherit(Unsolved)

var Category lipgloss.Style = lipgloss.NewStyle().
	Bold(false).
	MarginLeft(1).
	MarginRight(1).
	Inherit(Unsolved)

var Header lipgloss.Style = lipgloss.NewStyle().
	Bold(true).
	Inherit(Unsolved)

var Id lipgloss.Style = lipgloss.NewStyle().
	Inherit(Category)

var First lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#EBCB8B")).
	Bold(true)

var Second lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#D8DEE9")).
	Bold(true)

var Third lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#D08770")).
	Bold(true)
