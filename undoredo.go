//  Copyright 2013 The lime Authors.
//  Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"github.com/limetext/backend"
)

type (

	// Undo command will revert the last change
	Undo struct {
		backend.BypassUndoCommand
		hard bool
	}

	// Redo command will redo the last chnage
	Redo struct {
		backend.BypassUndoCommand
		hard bool
	}
)

// Run will execute the Undo command
func (c *Undo) Run(v *backend.View, e *backend.Edit) error {
	v.UndoStack().Undo(c.hard)
	return nil
}

// Run will execute the Redo command
func (c *Redo) Run(v *backend.View, e *backend.Edit) error {
	v.UndoStack().Redo(c.hard)
	return nil
}

func init() {
	register([]backend.Command{
		&Undo{hard: true},
		&Redo{hard: true},
	})

	registerByName([]namedCmd{
		{"soft_undo", &Undo{}},
		{"soft_redo", &Redo{}},
	})
}
