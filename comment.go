// Copyright 2015 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"strings"
	"unicode"

	"github.com/limetext/backend"
)

type (
	// ToggleComment converts the current selection into a comment,
	// or, if it is already a comment, converts it back into plain text
	// it'll add the comment symbol at the first character of the selection
	// regardless if the line is already commented.
	// While removing a comment during a toggle, it'll follow the same logic as above
	// and it'll only remove one layer of comments
	ToggleComment struct {
		backend.DefaultCommand
	}
)

// Run will execute the ToggleComment command
func (c *ToggleComment) Run(v *backend.View, e *backend.Edit) error {
	// TODO: Comment the line if we only have a cursor.
	// TODO: Expand the selection after altering it.
	// TODO: Align the comment characters for multiline selection.
	// TODO: Get the comment value from the Textmate files.
	comm := "//"

	for _, r := range v.Sel().Regions() {
		if r.Size() != 0 {
			t := v.Substr(r)

			trim := strings.TrimLeftFunc(t, unicode.IsSpace)
			if strings.HasPrefix(trim, comm) {
				repl := comm
				if strings.HasPrefix(trim, comm+" ") {
					repl += " "
				}

				t = strings.Replace(t, repl, "", 1)
			} else {
				t = strings.Replace(t, trim, comm+" "+trim, 1)
			}

			v.Replace(e, r, t)
		}
	}

	return nil
}

func init() {
	register([]backend.Command{
		&ToggleComment{},
	})
}
