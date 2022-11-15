// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgdoc

import (
	"bytes"
	"fmt"
	"golang.org/x/website/internal/backport/go/parser"
	"golang.org/x/website/internal/backport/go/token"
	"strings"
	"testing"

	"golang.org/x/website/internal/backport/html/template"
)

func TestSrcPosLink(t *testing.T) {
	for _, tc := range []struct {
		src  string
		line int
		low  int
		high int
		want template.HTML
	}{
		{"/src/fmt/print.go", 42, 30, 50, "/src/fmt/print.go?s=30:50#L32"},
		{"/src/fmt/print.go", 2, 1, 5, "/src/fmt/print.go?s=1:5#L1"},
		{"/src/fmt/print.go", 2, 0, 0, "/src/fmt/print.go#L2"},
		{"/src/fmt/print.go", 0, 0, 0, "/src/fmt/print.go"},
		{"/src/fmt/print.go", 0, 1, 5, "/src/fmt/print.go?s=1:5#L1"},
		{"fmt/print.go", 0, 0, 0, "/src/fmt/print.go"},
		{"fmt/print.go", 0, 1, 5, "/src/fmt/print.go?s=1:5#L1"},
	} {
		if got := srcPosLink(tc.src, tc.line, tc.low, tc.high); got != tc.want {
			t.Errorf("srcPosLink(%v, %v, %v, %v) = %v; want %v", tc.src, tc.line, tc.low, tc.high, got, tc.want)
		}
	}
}

func TestSanitize(t *testing.T) {
	for _, tc := range []struct {
		src  template.HTML
		want template.HTML
	}{
		{},
		{"foo", "foo"},
		{"func   f()", "func f()"},
		{"func f(a int,)", "func f(a int)"},
		{"func f(a int,\n)", "func f(a int)"},
		{"func f(\n\ta int,\n\tb int,\n\tc int,\n)", "func f(a int, b int, c int)"},
		{"  (   a,   b,  c  )  ", "(a, b, c)"},
		{"(  a,  b, c    int, foo   bar  ,  )", "(a, b, c int, foo bar)"},
		{"{   a,   b}", "{a, b}"},
		{"[   a,   b]", "[a, b]"},
	} {
		if got := sanitize(tc.src); got != tc.want {
			t.Errorf("sanitize(%v) = %v; want %v", tc.src, got, tc.want)
		}
	}
}

// Test that we add <span id="StructName.FieldName"> elements
// to the HTML of struct fields.
func TestStructFieldsIDAttributes(t *testing.T) {
	got := linkifySource(t, []byte(`
package foo

type T struct {
	NoDoc string

	// Doc has a comment.
	Doc string

	// Opt, if non-nil, is an option.
	Opt *int

	// Опция - другое поле.
	Опция bool
}
`))
	want := `type T struct {
<span id="T.NoDoc"></span>    NoDoc <a href="/pkg/builtin/#string">string</a>

<span id="T.Doc"></span>    <span class="comment">// Doc has a comment.</span>
    Doc <a href="/pkg/builtin/#string">string</a>

<span id="T.Opt"></span>    <span class="comment">// Opt, if non-nil, is an option.</span>
    Opt *<a href="/pkg/builtin/#int">int</a>

<span id="T.Опция"></span>    <span class="comment">// Опция - другое поле.</span>
    Опция <a href="/pkg/builtin/#bool">bool</a>
}`
	if got != want {
		t.Errorf("got: %s\n\nwant: %s\n", got, want)
	}
}

// Test that we add <span id="ConstName"> elements to the HTML
// of definitions in const and var specs.
func TestValueSpecIDAttributes(t *testing.T) {
	got := linkifySource(t, []byte(`
package foo

const (
	NoDoc string = "NoDoc"

	// Doc has a comment
	Doc = "Doc"

	NoVal
)`))
	want := `const (
    <span id="NoDoc">NoDoc</span> <a href="/pkg/builtin/#string">string</a> = &#34;NoDoc&#34;

    <span class="comment">// Doc has a comment</span>
    <span id="Doc">Doc</span> = &#34;Doc&#34;

    <span id="NoVal">NoVal</span>
)`
	if got != want {
		t.Errorf("got: %s\n\nwant: %s\n", got, want)
	}
}

func TestCompositeLitLinkFields(t *testing.T) {
	got := linkifySource(t, []byte(`
package foo

type T struct {
	X int
}

var S T = T{X: 12}`))
	want := `type T struct {
<span id="T.X"></span>    X <a href="/pkg/builtin/#int">int</a>
}
var <span id="S">S</span> <a href="#T">T</a> = <a href="#T">T</a>{<a href="#T.X">X</a>: 12}`
	if got != want {
		t.Errorf("got: %s\n\nwant: %s\n", got, want)
	}
}

func TestFuncDeclNotLink(t *testing.T) {
	// Function.
	got := linkifySource(t, []byte(`
package http

func Get(url string) (resp *Response, err error)`))
	want := `func Get(url <a href="/pkg/builtin/#string">string</a>) (resp *<a href="#Response">Response</a>, err <a href="/pkg/builtin/#error">error</a>)`
	if got != want {
		t.Errorf("got: %s\n\nwant: %s\n", got, want)
	}

	// Method.
	got = linkifySource(t, []byte(`
package http

func (h Header) Get(key string) string`))
	want = `func (h <a href="#Header">Header</a>) Get(key <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a>`
	if got != want {
		t.Errorf("got: %s\n\nwant: %s\n", got, want)
	}
}

func linkifySource(t *testing.T, src []byte) string {
	fset := token.NewFileSet()
	af, err := parser.ParseFile(fset, "foo.go", src, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	pi := &Page{
		fset: fset,
	}
	sep := ""
	for _, decl := range af.Decls {
		buf.WriteString(sep)
		sep = "\n"
		buf.WriteString(string(pi.Node(decl)))
	}
	return buf.String()
}

func TestReplaceLeadingIndentation(t *testing.T) {
	oldIndent := strings.Repeat(" ", 2)
	newIndent := strings.Repeat(" ", 4)
	tests := []struct {
		src, want string
	}{
		{"  foo\n    bar\n  baz", "    foo\n      bar\n    baz"},
		{"  '`'\n  '`'\n", "    '`'\n    '`'\n"},
		{"  '\\''\n  '`'\n", "    '\\''\n    '`'\n"},
		{"  \"`\"\n  \"`\"\n", "    \"`\"\n    \"`\"\n"},
		{"  `foo\n  bar`", "    `foo\n      bar`"},
		{"  `foo\\`\n  bar", "    `foo\\`\n    bar"},
		{"  '\\`'`foo\n  bar", "    '\\`'`foo\n      bar"},
		{
			"  if true {\n    foo := `One\n    \tTwo\nThree`\n  }\n",
			"    if true {\n      foo := `One\n        \tTwo\n    Three`\n    }\n",
		},
	}
	for _, tc := range tests {
		if got := replaceLeadingIndentation(tc.src, oldIndent, newIndent); got != tc.want {
			t.Errorf("replaceLeadingIndentation:\n%v\n---\nhave:\n%v\n---\nwant:\n%v\n",
				tc.src, got, tc.want)
		}
	}
}

func TestFilterOutBuildAnnotations(t *testing.T) {
	// TODO: simplify this by using a multiline string once we stop
	// using go vet from 1.10 on the build dashboard.
	// https://golang.org/issue/26627
	src := []byte("// +build !foo\n" +
		"// +build !anothertag\n" +
		"\n" +
		"// non-tag comment\n" +
		"\n" +
		"package foo\n" +
		"\n" +
		"func bar() int {\n" +
		"	return 42\n" +
		"}\n")

	fset := token.NewFileSet()
	af, err := parser.ParseFile(fset, "foo.go", src, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	var found bool
	for _, cg := range af.Comments {
		if strings.HasPrefix(cg.Text(), "+build ") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("TestFilterOutBuildAnnotations is broken: missing build tag in test input")
	}

	found = false
	for _, cg := range filterOutBuildAnnotations(af.Comments) {
		if strings.HasPrefix(cg.Text(), "+build ") {
			t.Errorf("filterOutBuildAnnotations failed to filter build tag")
		}

		if strings.Contains(cg.Text(), "non-tag comment") {
			found = true
		}
	}
	if !found {
		t.Errorf("filterOutBuildAnnotations should not remove non-build tag comment")
	}
}

// Verify that scanIdentifier isn't quadratic.
// This doesn't actually measure and fail on its own, but it was previously
// very obvious when running by hand.
//
// TODO: if there's a reliable and non-flaky way to test this, do so.
// Maybe count user CPU time instead of wall time? But that's not easy
// to do portably in Go.
func TestStructField(t *testing.T) {
	for _, n := range []int{10, 100, 1000, 10000} {
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			var buf bytes.Buffer
			fmt.Fprintf(&buf, "package foo\n\ntype T struct {\n")
			for i := 0; i < n; i++ {
				fmt.Fprintf(&buf, "\t// Field%d is foo.\n\tField%d int\n\n", i, i)
			}
			fmt.Fprintf(&buf, "}\n")
			linkifySource(t, buf.Bytes())
		})
	}
}
