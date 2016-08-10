// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"fmt"

	"github.com/limetext/backend"
)

type (

	// Save command will save the currently
	// opened file to the disk
	Save struct {
		backend.DefaultCommand
	}

	// PromptSaveAs command will let us save
	// the currently opened
	// file with a different file name
	PromptSaveAs struct {
		backend.DefaultCommand
	}

	// SaveAll command will save all the opene files
	SaveAll struct {
		backend.DefaultCommand
	}
)

// Run will execute the Save command
func (c *Save) Run(v *backend.View, e *backend.Edit) error {
	err := v.Save()
	if err != nil {
		backend.GetEditor().Frontend().ErrorMessage(fmt.Sprintf("Failed to save %s:n%s", v.FileName(), err))
		return err
	}
	return nil
}

// Run will execute the PromptSaveAs command
func (c *PromptSaveAs) Run(v *backend.View, e *backend.Edit) error {
	dir := viewDirectory(v)
	fe := backend.GetEditor().Frontend()
	files := fe.Prompt("Save file", dir, backend.PROMPT_SAVE_AS)
	if len(files) == 0 {
		return nil
	}

	name := files[0]
	if err := v.SaveAs(name); err != nil {
		fe.ErrorMessage(fmt.Sprintf("Failed to save as %s:%s", name, err))
		return err
	}
	return nil
}

// Run will execute the SaveAll command
func (c *SaveAll) Run(w *backend.Window) error {
	fe := backend.GetEditor().Frontend()
	for _, v := range w.Views() {
		if err := v.Save(); err != nil {
			fe.ErrorMessage(fmt.Sprintf("Failed to save %s:n%s", v.FileName(), err))
			return err
		}
	}
	return nil
}

func init() {
	register([]backend.Command{
		&Save{},
		&PromptSaveAs{},
		&SaveAll{},
	})
}
