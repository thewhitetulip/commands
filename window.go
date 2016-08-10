// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"github.com/limetext/backend"
)

type (
	// NewWindow command lets us open a new window
	// of lime editor
	NewWindow struct {
		backend.DefaultCommand
	}

	// CloseAll command lets us close all the
	// open views inside the current window
	CloseAll struct {
		backend.DefaultCommand
	}

	// CloseWindow command lets us close the current window
	CloseWindow struct {
		backend.DefaultCommand
	}

	// NewWindowApp will create a new window and open an
	// empty editor and set it as active
	NewWindowApp struct {
		backend.DefaultCommand
	}

	// CloseWindowApp command will close the windows
	// which are active
	CloseWindowApp struct {
		backend.DefaultCommand
	}
)

// Run executes the NewWindow command
func (c *NewWindow) Run(w *backend.Window) error {
	ed := backend.GetEditor()
	ed.SetActiveWindow(ed.NewWindow())
	return nil
}

// Run executes the CloseAll command
func (c *CloseAll) Run(w *backend.Window) error {
	w.CloseAllViews()
	return nil
}

// Run executes the CloseWindow command
func (c *CloseWindow) Run(w *backend.Window) error {
	ed := backend.GetEditor()
	ed.ActiveWindow().Close()
	return nil
}

// Run executes the NewWindowApp command
func (c *NewWindowApp) Run() error {
	ed := backend.GetEditor()
	ed.SetActiveWindow(ed.NewWindow())
	return nil
}

// Run executes the CloseWindowApp command
func (c *CloseWindowApp) Run() error {
	ed := backend.GetEditor()
	ed.ActiveWindow().Close()
	return nil
}

func (c *NewWindowApp) IsChecked() bool {
	return false
}

func (c *CloseWindowApp) IsChecked() bool {
	return false
}

func init() {
	register([]backend.Command{
		&NewWindow{},
		&CloseAll{},
		&CloseWindow{},
	})

	registerByName([]namedCmd{
		{"new_window", &NewWindowApp{}},
		{"close_window", &CloseWindowApp{}},
	})
}
