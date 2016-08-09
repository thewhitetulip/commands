// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"sort"
	"strings"

	"github.com/limetext/backend"
	"github.com/limetext/text"
)

type (
	//Copy command will be used to copy
	//the selected lines of text to the clipboard
	Copy struct {
		backend.DefaultCommand
	}
	//Cut command will be used to cut the selected lines
	//from the view
	Cut struct {
		backend.DefaultCommand
	}
	//Paste command will be used to paste the clipboard
	//contents into a view
	Paste struct {
		backend.DefaultCommand
	}
)

func getRegions(v *backend.View, cut bool) *text.RegionSet {
	rs := &text.RegionSet{}
	regions := v.Sel().Regions()
	sort.Sort(regionSorter(regions))
	rs.AddAll(regions)

	he, ae := rs.HasEmpty(), !rs.HasNonEmpty() || cut
	for _, r := range rs.Regions() {
		if ae && r.Empty() {
			rs.Add(v.FullLineR(r))
		} else if he && r.Empty() {
			rs.Subtract(r)
		}
	}

	return rs
}

func getSelSubstrs(v *backend.View, rs *text.RegionSet) []string {
	var add, s1 string
	s := make([]string, len(rs.Regions()))
	for i, r := range rs.Regions() {
		add = ""
		s1 = v.Substr(r)
		if !v.Sel().HasNonEmpty() && !strings.HasSuffix(s1, "\n") {
			add = "\n"
		}
		s[i] = s1 + add
	}
	return s
}

//Run will execute the Copy command when initialized
func (c *Copy) Run(v *backend.View, e *backend.Edit) error {
	rs := getRegions(v, false)
	s := getSelSubstrs(v, rs)

	backend.GetEditor().SetClipboard(strings.Join(s, "\n"))

	return nil
}

//Run will execute the Cut command when initialized
func (c *Cut) Run(v *backend.View, e *backend.Edit) error {
	s := getSelSubstrs(v, getRegions(v, false))

	rs := getRegions(v, true)
	regions := rs.Regions()
	sort.Sort(sort.Reverse(regionSorter(regions)))
	for _, r := range regions {
		v.Erase(e, r)
	}

	backend.GetEditor().SetClipboard(strings.Join(s, "\n"))

	return nil
}

//Run will execute the Paste command when initialized
func (c *Paste) Run(v *backend.View, e *backend.Edit) error {
	// TODO: Paste the entire line on the line before the cursor if a
	//		 line was autocopied.

	ed := backend.GetEditor()

	rs := &text.RegionSet{}
	regions := v.Sel().Regions()
	sort.Sort(sort.Reverse(regionSorter(regions)))
	rs.AddAll(regions)
	for _, r := range rs.Regions() {
		v.Replace(e, r, ed.GetClipboard())
	}

	return nil
}

func init() {
	register([]backend.Command{
		&Copy{},
		&Cut{},
		&Paste{},
	})
}
