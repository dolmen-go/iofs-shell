// Package shell exposes an interactive user interface that allows to browse an [io/fs] filesystem.
//
// See examples:
//   - [github.com/dolmen-go/iofs-shell/examples/browse-zip]
//   - [github.com/dolmen-go/iofs-shell/examples/browse-sqlar]
package shell

import (
	"fmt"
	iofs "io/fs"
	"strings"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

// Browse opens a terminal-based interactive browser that allows to navigate
// the given filesystem.
//
// Keys:
//   - q: quit
//   - escape: leave a subdirectory
//   - right, left arrow: enter/leave a subdirectory
//   - up, down arrow: navigate the file list of the current directory
//   - g, G: go to top/bottom
func Browse(fs iofs.FS, cwd string) error {
	if !iofs.ValidPath(cwd) {
		return &iofs.PathError{Path: cwd, Op: "open", Err: iofs.ErrInvalid}
	}

	fp := filepicker.New()
	fp.FS = fs
	fp.CurrentDirectory = cwd
	p := tea.NewProgram(browser{
		filepicker: fp,
	})
	_, err := p.Run()
	return err
}

type browser struct {
	filepicker   filepicker.Model
	selectedFile string
}

func (b browser) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return b.filepicker.Init()
}

func (b browser) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return b, tea.Quit

		}
	}

	var cmd tea.Cmd
	b.filepicker, cmd = b.filepicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := b.filepicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		b.selectedFile = path
	}

	/*
		// Did the user select a disabled file?
		// This is only necessary to display an error to the user.
		if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
			// Let's clear the selectedFile and display an error.
			m.err = errors.New(path + " is not valid.")
			m.selectedFile = ""
			return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
		}
	*/

	return b, cmd
}

func (b browser) View() string {
	var w strings.Builder
	fmt.Fprintf(&w, "\n%s\n\n", b.filepicker.CurrentDirectory)
	fmt.Fprintln(&w, b.filepicker.View())
	return w.String()
}
