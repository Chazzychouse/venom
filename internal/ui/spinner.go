package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// SpinnerModel wraps a spinner with a message
type SpinnerModel struct {
	spinner spinner.Model
	message string
	done    bool
}

// NewSpinner creates a spinner with a message
// Usage:
//
//	s := ui.NewSpinner("Loading...")
//	p := tea.NewProgram(s)
//	go func() {
//	    doWork()
//	    p.Send(ui.SpinnerDone{})
//	}()
//	p.Run()
func NewSpinner(message string) SpinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = Spinner
	return SpinnerModel{
		spinner: s,
		message: message,
	}
}

// SpinnerDone signals the spinner to stop
type SpinnerDone struct{}

func (m SpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case SpinnerDone:
		m.done = true
		return m, tea.Quit
	case tea.KeyMsg:
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m SpinnerModel) View() string {
	if m.done {
		return Success.Render("✓") + " " + m.message + "\n"
	}
	return m.spinner.View() + " " + m.message
}

// RunWithSpinner executes a function while showing a spinner
// This is the simplest way to show loading state
//
// Usage:
//
//	err := ui.RunWithSpinner("Processing...", func() error {
//	    return doExpensiveWork()
//	})
func RunWithSpinner(message string, fn func() error) error {
	m := NewSpinner(message)
	p := tea.NewProgram(m)

	var fnErr error
	go func() {
		fnErr = fn()
		p.Send(SpinnerDone{})
	}()

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("spinner error: %w", err)
	}
	return fnErr
}
