package ui

import "github.com/charmbracelet/lipgloss"

// Color palette - adjust to your taste
var (
	Primary   = lipgloss.Color("#7C3AED") // Purple
	Secondary = lipgloss.Color("#10B981") // Green
	Muted     = lipgloss.Color("#6B7280") // Gray
	Error     = lipgloss.Color("#EF4444") // Red
	Warning   = lipgloss.Color("#F59E0B") // Amber
)

// Reusable styles
var (
	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(Primary)

	Subtle = lipgloss.NewStyle().
		Foreground(Muted)

	Success = lipgloss.NewStyle().
		Foreground(Secondary)

	ErrorText = lipgloss.NewStyle().
			Foreground(Error)

	// Box for displaying metadata, results, etc.
	Box = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Primary).
		Padding(0, 1)
)

// TODO: Add more styles as needed
// - Table styles for metadata display
// - Progress bar styling
// - Spectrum bar colors (gradient from green to red)
//
// Lipgloss docs: https://github.com/charmbracelet/lipgloss
