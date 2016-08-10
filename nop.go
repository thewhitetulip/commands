// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import "github.com/limetext/backend"

type (
	// NopApplication is
	NopApplication struct {
		backend.BypassUndoCommand
	}
	// NopWindow is
	NopWindow struct {
		backend.BypassUndoCommand
	}
	// NopText is
	NopText struct {
		backend.BypassUndoCommand
	}
)

// Run will execute the NopApplication command
func (c *NopApplication) Run() error {
	return nil
}

// IsChecked will represent if the command will
// contain a checkbox in the UI
func (c *NopApplication) IsChecked() bool {
	return false
}

// Run will execute the NopWindow command
func (c *NopWindow) Run(w *backend.Window) error {
	return nil
}

// Run will execute the NopText command
func (c *NopText) Run(v *backend.View, e *backend.Edit) error {
	return nil
}

func init() {
	registerByName([]namedCmd{
		{"nop", &NopApplication{}},
		{"nop", &NopWindow{}},
		{"nop", &NopText{}},
	})
}
