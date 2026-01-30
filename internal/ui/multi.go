package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Task represents a single task in a multi-task display
type Task struct {
	ID      string
	Label   string
	Status  TaskStatus
	Percent float64 // For progress tasks (0.0 to 1.0)
}

type TaskStatus int

const (
	TaskPending TaskStatus = iota
	TaskRunning
	TaskComplete
	TaskFailed
)

// MultiModel displays multiple concurrent tasks
type MultiModel struct {
	tasks    []Task
	spinner  spinner.Model
	progress progress.Model
	done     bool
}

// TaskUpdate updates a specific task's status
type TaskUpdate struct {
	ID      string
	Status  TaskStatus
	Percent float64
}

// MultiDone signals all tasks complete
type MultiDone struct{}

// NewMulti creates a multi-task display
// Usage:
//
//	tasks := []ui.Task{
//	    {ID: "scan", Label: "Scanning files"},
//	    {ID: "analyze", Label: "Analyzing audio"},
//	    {ID: "write", Label: "Writing metadata"},
//	}
//	m := ui.NewMulti(tasks)
//	p := tea.NewProgram(m)
//	go func() {
//	    p.Send(ui.TaskUpdate{ID: "scan", Status: ui.TaskRunning})
//	    // do work...
//	    p.Send(ui.TaskUpdate{ID: "scan", Status: ui.TaskComplete})
//	    // etc
//	    p.Send(ui.MultiDone{})
//	}()
//	p.Run()
func NewMulti(tasks []Task) MultiModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = Spinner

	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(20),
		progress.WithoutPercentage(),
	)

	return MultiModel{
		tasks:    tasks,
		spinner:  s,
		progress: p,
	}
}

func (m MultiModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m MultiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TaskUpdate:
		for i := range m.tasks {
			if m.tasks[i].ID == msg.ID {
				m.tasks[i].Status = msg.Status
				m.tasks[i].Percent = msg.Percent
				break
			}
		}
		return m, nil
	case MultiDone:
		m.done = true
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m MultiModel) View() string {
	var b strings.Builder

	for _, task := range m.tasks {
		var icon string
		var style lipgloss.Style

		switch task.Status {
		case TaskPending:
			icon = Subtle.Render("○")
			style = Subtle
		case TaskRunning:
			icon = m.spinner.View()
			style = lipgloss.NewStyle()
		case TaskComplete:
			icon = Success.Render("✓")
			style = Success
		case TaskFailed:
			icon = ErrorText.Render("✗")
			style = ErrorText
		}

		line := icon + " " + style.Render(task.Label)

		// Show progress bar for running tasks with percent > 0
		if task.Status == TaskRunning && task.Percent > 0 {
			line += " " + m.progress.ViewAs(task.Percent)
		}

		b.WriteString(line + "\n")
	}

	return b.String()
}
