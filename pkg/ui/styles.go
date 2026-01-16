package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// TODO check pallete
var (
	ColorPrimary   = lipgloss.Color("#7D56F4") // Purple accent
	ColorSecondary = lipgloss.Color("#00D9FF") // Cyan accent
	ColorSuccess   = lipgloss.Color("#00FF87") // Green
	ColorWarning   = lipgloss.Color("#FFD700") // Yellow
	ColorError     = lipgloss.Color("#FF5F87") // Red
	ColorMuted     = lipgloss.Color("#6C7086") // Gray
	ColorBorder    = lipgloss.Color("#45475A") // Dark gray
)

type Styles struct {
	//Headers
	Title         lipgloss.Style
	Subtitle      lipgloss.Style
	SectionHeader lipgloss.Style

	// Content
	Success lipgloss.Style
	Error   lipgloss.Style
	Warning lipgloss.Style
	Info    lipgloss.Style
	Muted   lipgloss.Style

	// UI elements
	Box       lipgloss.Style
	List      lipgloss.Style
	Code      lipgloss.Style
	Highlight lipgloss.Style
	Badge     lipgloss.Style
	Separator lipgloss.Style

	//Status Indicators
	StatusRunning lipgloss.Style
	StatusStopped lipgloss.Style
	StatusSuccess lipgloss.Style
	StatusError   lipgloss.Style
}

//DefaultStyle (main constructor)

func DefaultStyles() *Styles {
	return &Styles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			MarginTop(1).
			MarginBottom(1),

		Subtitle: lipgloss.NewStyle().
			Foreground(ColorSecondary).
			MarginBottom(1),
		SectionHeader: lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			MarginTop(1),

		// Status messages
		Success: lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true),

		Error: lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true),

		Warning: lipgloss.NewStyle().
			Foreground(ColorWarning).
			Bold(true),

		Info: lipgloss.NewStyle().
			Foreground(ColorSecondary),

		Muted: lipgloss.NewStyle().
			Foreground(ColorMuted),

		// UI Elements
		Box: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(1, 2),

		List: lipgloss.NewStyle().
			MarginLeft(2),

		Code: lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Background(lipgloss.Color("#1E1E2E")).
			Padding(0, 1),

		Highlight: lipgloss.NewStyle().
			Background(ColorPrimary).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1),

		Badge: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(ColorPrimary).
			Padding(0, 1).
			Bold(true),

		Separator: lipgloss.NewStyle().
			Foreground(ColorBorder),

		// Status indicators with icons
		StatusRunning: lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true),

		StatusStopped: lipgloss.NewStyle().
			Foreground(ColorWarning).
			Bold(true),

		StatusSuccess: lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true),

		StatusError: lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true),
	}
}

var Icons = struct {
	Success   string
	Error     string
	Warning   string
	Info      string
	Running   string
	Stopped   string
	Arrow     string
	Bullet    string
	Check     string
	Cross     string
	Hourglass string
}{
	Success:   "✓",
	Error:     "✗",
	Warning:   "⚠",
	Info:      "ℹ",
	Running:   "●",
	Stopped:   "○",
	Arrow:     "→",
	Bullet:    "•",
	Check:     "✓",
	Cross:     "✗",
	Hourglass: "⏳",
}

// Common message formatters
func (s *Styles) SuccessMsg(msg string) string {
	return s.Success.Render(Icons.Success + " " + msg)
}

func (s *Styles) ErrorMsg(msg string) string {
	return s.Error.Render(Icons.Error + " " + msg)
}

func (s *Styles) WarningMsg(msg string) string {
	return s.Warning.Render(Icons.Warning + " " + msg)
}

func (s *Styles) InfoMsg(msg string) string {
	return s.Info.Render(Icons.Info + " " + msg)
}

func (s *Styles) RunningMsg(msg string) string {
	return s.StatusRunning.Render(Icons.Running + " " + msg)
}

func (s *Styles) StoppedMsg(msg string) string {
	return s.StatusStopped.Render(Icons.Stopped + " " + msg)
}

//Horizontal separator

func (s *Styles) HSeparator(width int) string {
	line := ""
	for i := 0; i < width; i++ {
		line += "-"
	}
	return s.Separator.Render(line)
}
