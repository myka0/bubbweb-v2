//go:build !js
// +build !js

package bubbweb

import (
	tea "github.com/charmbracelet/bubbletea/v2"
)

var NewProgram = tea.NewProgram
