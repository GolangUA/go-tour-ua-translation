// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comment

import (
	"os/exec"
	"sort"
	"strings"
	"testing"

	"golang.org/x/website/internal/backport/diff"
)

func TestStd(t *testing.T) {
	out, err := exec.Command("go", "list", "std").CombinedOutput()
	if err != nil {
		t.Fatalf("%v\n%s", err, out)
	}

	var list []string
	for _, pkg := range strings.Fields(string(out)) {
		if !strings.Contains(pkg, "/") {
			list = append(list, pkg)
		}
	}
	sort.Strings(list)

	have := strings.Join(stdPkgs, "\n") + "\n"
	want := strings.Join(list, "\n") + "\n"
	if have != want {
		t.Errorf("stdPkgs is out of date: regenerate with 'go generate'\n%s", diff.Diff("stdPkgs", []byte(have), "want", []byte(want)))
	}
}
