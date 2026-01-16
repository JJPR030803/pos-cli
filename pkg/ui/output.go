package ui

import (
	"fmt"
	"strings"
)

// Ouptut provides formatted output utilities

type Output struct {
	Styles  *Styles
	Verbose bool
}

// Kind of like a contructor
func NewOutput(verbose bool) *Output {
	return &Output{
		Styles:  DefaultStyles(),
		Verbose: verbose,
	}
}

// Print title
func (o *Output) PrintTitle(title string) {
	fmt.Println(o.Styles.Title.Render(title))
}

//etc

func (o *Output) PrintSubtitle(subtitle string) {
	fmt.Println(o.Styles.Subtitle.Render(subtitle))
}

func (o *Output) PrintSection(section string) {
	fmt.Println(o.Styles.SectionHeader.Render(section))
}

// PrintSuccess prints a success message
func (o *Output) PrintSuccess(msg string) {
	fmt.Println(o.Styles.SuccessMsg(msg))
}

// PrintError prints an error message
func (o *Output) PrintError(msg string) {
	fmt.Println(o.Styles.ErrorMsg(msg))
}

// PrintWarning prints a warning message
func (o *Output) PrintWarning(msg string) {
	fmt.Println(o.Styles.WarningMsg(msg))
}

// PrintInfo prints an info message
func (o *Output) PrintInfo(msg string) {
	fmt.Println(o.Styles.InfoMsg(msg))
}

// Verbose muted not usually important
func (o *Output) PrintVerbose(msg string) {
	if o.Verbose {
		fmt.Println(o.Styles.Muted.Render("$ " + msg))
	}
}

// PrintCommand prints a command being executed
// Shows commands in a distinct "code" style
func (o *Output) PrintCommand(cmd string) {
	fmt.Println(o.Styles.Code.Render("$ " + cmd))
}

// PrintBox prints content in a bordered box
func (o *Output) PrintBox(content string) {
	fmt.Println(o.Styles.Box.Render(content))
}

// PrintList prints a bulleted list
func (o *Output) PrintList(items []string) {
	for _, item := range items {
		fmt.Println(o.Styles.List.Render(Icons.Bullet + " " + item))
	}
}

// PrintSeparator prints a horizontal separator
// Default width is 50, but you can customize by calling o.Styles.HSeparator(width)
func (o *Output) PrintSeparator() {
	fmt.Println(o.Styles.HSeparator(50))
}
func (o *Output) PrintStatus(service string, running bool) {
	if running {
		fmt.Printf("%s %s\n",
			o.Styles.StatusRunning.Render(Icons.Running),
			service)
	} else {
		fmt.Printf("%s %s\n",
			o.Styles.StatusStopped.Render(Icons.Stopped),
			service)
	}
}

// PrintKeyValue prints a key-value pair
// Example: "Port: → 3000"
func (o *Output) PrintKeyValue(key, value string) {
	fmt.Printf("%s %s %s\n",
		o.Styles.Muted.Render(key+":"),
		o.Styles.Info.Render(Icons.Arrow),
		value)
}

// PrintTable prints a simple two-column table
// Pass in rows as [][2]string where each row is [key, value]
func (o *Output) PrintTable(rows [][2]string) {
	maxKeyLen := 0
	for _, row := range rows {
		if len(row[0]) > maxKeyLen {
			maxKeyLen = len(row[0])
		}
	}

	for _, row := range rows {
		key := row[0]
		value := row[1]
		padding := strings.Repeat(" ", maxKeyLen-len(key))

		fmt.Printf("%s%s %s %s\n",
			o.Styles.Muted.Render(key),
			padding,
			o.Styles.Info.Render(Icons.Arrow),
			value)
	}
}

// StartSpinner prints a "loading" message
// For now, just prints the message. Can be enhanced with actual spinners later
func (o *Output) StartSpinner(msg string) {
	fmt.Print(o.Styles.Info.Render(Icons.Hourglass + " " + msg + "..."))
}

// StopSpinner completes a spinner with success or error
func (o *Output) StopSpinner(success bool, msg string) {
	fmt.Print("\r") // Clear line (carriage return)
	if success {
		o.PrintSuccess(msg)
	} else {
		o.PrintError(msg)
	}
}

// PrintBadge prints a badge-style message
// Useful for highlighting important info
func (o *Output) PrintBadge(label string) {
	fmt.Println(o.Styles.Badge.Render(label))
}

// Newline prints a blank line
// Simple helper for spacing
func (o *Output) Newline() {
	fmt.Println()
}
