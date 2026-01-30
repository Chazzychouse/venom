package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

// ProgressModel wraps a progress bar with label
type ProgressModel struct {
	progress progress.Model
	label    string
	percent  float64
	done     bool
}

// ProgressUpdate updates the progress bar percentage (0.0 to 1.0)
type ProgressUpdate float64

// ProgressDone signals completion
type ProgressDone struct{}

// NewProgress creates a styled progress bar
// Usage:
//
//	p := tea.NewProgram(ui.NewProgress("Scanning files..."))
//	go func() {
//	    for i := 0; i <= 100; i++ {
//	        p.Send(ui.ProgressUpdate(float64(i) / 100))
//	        time.Sleep(50 * time.Millisecond)
//	    }
//	    p.Send(ui.ProgressDone{})
//	}()
//	p.Run()
func NewProgress(label string) ProgressModel {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
	)
	return ProgressModel{
		progress: p,
		label:    label,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return nil
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ProgressUpdate:
		m.percent = float64(msg)
		return m, nil
	case ProgressDone:
		m.done = true
		m.percent = 1.0
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case progress.FrameMsg:
		pm, cmd := m.progress.Update(msg)
		m.progress = pm.(progress.Model)
		return m, cmd
	}
	return m, nil
}

func (m ProgressModel) View() string {
	if m.done {
		return Success.Render("✓") + " " + m.label + " - Complete\n"
	}
	pct := fmt.Sprintf(" %3.0f%%", m.percent*100)
	return m.label + "\n" + m.progress.ViewAs(m.percent) + pct
}

// TrackProgress runs a function and tracks progress via a channel
// Send values 0.0-1.0 to the channel, close when done
//
// Usage:
//
//	progress := make(chan float64)
//	go func() {
//	    for i, file := range files {
//	        processFile(file)
//	        progress <- float64(i+1) / float64(len(files))
//	    }
//	    close(progress)
//	}()
//	ui.TrackProgress("Processing files...", progress)
func TrackProgress(label string, updates <-chan float64) error {
	m := NewProgress(label)
	p := tea.NewProgram(m)

	go func() {
		for pct := range updates {
			p.Send(ProgressUpdate(pct))
		}
		p.Send(ProgressDone{})
	}()

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("progress error: %w", err)
	}
	return nil
}
