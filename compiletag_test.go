package compiler

import (
	p "github.com/realistschuckle/gohaml/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagWithOnlyTagName(t *testing.T) {
	opts := CompilerOpts{}
	node := &p.TagNode{}
	node.Name = "p"
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<p></p>")
}

func TestAutoClosingTagInXhtml(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br />")
}

func TestAutoClosingTagInHtml4(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html4"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br>")
}

func TestAutoClosingTagInHtml5(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html5"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br>")
}

func TestSelfClosingTagInXhtml(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	node := &p.TagNode{}
	node.Name = "whatever"
	node.Close = true
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<whatever />")
}

func TestSelfClosingTagInHtml5(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html5"
	node := &p.TagNode{}
	node.Name = "whatever"
	node.Close = true
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<whatever>")
}

func TestTagWithSingleClassNameOutputsClassAttribute(t *testing.T) {
	opts := CompilerOpts{}
	node := &p.TagNode{}
	node.Name = "h1"
	node.Classes = []string{"ui-helper-hidden"}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<h1 class='ui-helper-hidden'></h1>")
}

func TestAutoClosingTagInXhtmlWithClass(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Classes = []string{"ui-helper-hidden"}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br class='ui-helper-hidden' />")
}

func TestAutoClosingTagInHtml4WithClass(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html4"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Classes = []string{"ui-state-highlight"}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br class='ui-state-highlight'>")
}

func TestAutoClosingTagInHtml5WithClass(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html5"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Classes = []string{"ui-state-default"}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br class='ui-state-default'>")
}

func TestAutoClosingTagInXhtmlWithId(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Id = "marklar"
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br id='marklar' />")
}

func TestAutoClosingTagInHtml4WithId(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html4"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Id = "marklar"
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br id='marklar'>")
}

func TestAutoClosingTagInHtml5WithId(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html5"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Id = "marklar"
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br id='marklar'>")
}

func TestTagWithIdOuptutsId(t *testing.T) {
	opts := CompilerOpts{}
	node := &p.TagNode{}
	node.Name = "h1"
	node.Id = "marklar"
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<h1 id='marklar'></h1>")
}

func TestAutoClosingTagInXhtmlWithIdAndClassOutputsClassesFirst(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Id = "marklar"
	node.Classes = []string{"one", "two", "three"}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br class='one two three' id='marklar' />")
}

func TestAutoClosingTagInHtml4WithIdAndClassesOutputsClassesFirst(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html4"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Id = "marklar"
	node.Classes = []string{"one", "two", "three"}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br class='one two three' id='marklar'>")
}

func TestAutoClosingTagInHtml5WithIdAndClassesOutputsClassesFirst(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html5"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Id = "marklar"
	node.Classes = []string{"one", "two", "three"}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br class='one two three' id='marklar'>")
}

func TestTagWithIdAndClassesOuptutsClassesFirst(t *testing.T) {
	opts := CompilerOpts{}
	node := &p.TagNode{}
	node.Name = "h1"
	node.Id = "marklar"
	node.Classes = []string{"one", "two", "three"}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<h1 class='one two three' id='marklar'></h1>")
}

func TestTagWithChild(t *testing.T) {
	opts := CompilerOpts{}
	a := &p.TagNode{}
	a.Name = "a"
	a.Indent = "  "
	a.LineBreak = "\n"
	div := &p.TagNode{}
	div.Name = "div"
	div.Children = []p.Node{a}
	div.LineBreak = "\n"
	nodes := []p.Node{div}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)
	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<div>\n  <a></a>\n</div>\n")
}

func TestTagWithIdAndClassesWithChildWithIdAndClasses(t *testing.T) {
	opts := CompilerOpts{}
	a := &p.TagNode{}
	a.Name = "a"
	a.Indent = "  "
	a.LineBreak = "\n"
	a.Id = "my_a"
	a.Classes = []string{"foo", "bar"}
	div := &p.TagNode{}
	div.Name = "div"
	div.Children = []p.Node{a}
	div.LineBreak = "\n"
	div.Id = "my_div"
	div.Classes = []string{"baz", "boo"}
	nodes := []p.Node{div}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)
	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<div class='baz boo' id='my_div'>\n  <a class='foo bar' id='my_a'></a>\n</div>\n")
}

func TestTagWithContent(t *testing.T) {
	opts := CompilerOpts{}
	s := &p.StaticNode{}
	s.Content = "This is my weapon; this is my div."
	div := &p.TagNode{}
	div.Name = "div"
	div.Children = []p.Node{s}
	div.LineBreak = ""
	nodes := []p.Node{div}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)
	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<div>This is my weapon; this is my div.</div>")
}

func TestTagWithContentAndIdAndClasses(t *testing.T) {
	opts := CompilerOpts{}
	s := &p.StaticNode{}
	s.Content = "This is my weapon; this is my div."
	div := &p.TagNode{}
	div.Name = "div"
	div.Children = []p.Node{s}
	div.LineBreak = ""
	div.Id = "mother"
	div.Indent = "    "
	div.Classes = []string{"father brother sister"}
	nodes := []p.Node{div}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)
	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "    <div class='father brother sister' id='mother'>This is my weapon; this is my div.</div>")
}

func TestTagWithChildContentAndIdAndClasses(t *testing.T) {
	opts := CompilerOpts{}
	s := &p.StaticLineNode{}
	s.Content = "This is my weapon; this is my div."
	s.Indent = "      "
	div := &p.TagNode{}
	div.Name = "div"
	div.Children = []p.Node{s}
	div.LineBreak = "\n"
	div.Id = "mother"
	div.Indent = "    "
	div.Classes = []string{"father brother sister"}
	nodes := []p.Node{div}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)
	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "    <div class='father brother sister' id='mother'>\n      This is my weapon; this is my div.\n    </div>\n")
}

func TestTagWithValueAndAttributes(t *testing.T) {
	opts := CompilerOpts{}
	s := &p.StaticNode{}
	s.Content = "This is my weapon; this is my div."
	div := &p.TagNode{}
	div.Name = "div"
	div.Children = []p.Node{s}
	div.LineBreak = "\n"
	div.Id = "mother"
	div.Indent = "    "
	div.Classes = []string{"father brother sister"}
	div.Attrs = []p.Attribute{p.Attribute{"uno", &p.StaticNode{"dos"}}}
	nodes := []p.Node{div}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)
	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, "    <div class='father brother sister' id='mother' uno='dos'>This is my weapon; this is my div.</div>\n", output.Content)
}

func TestTagWithChildValueAndAttributes(t *testing.T) {
	opts := CompilerOpts{}
	s := &p.StaticLineNode{}
	s.Content = "This is my weapon; this is my div."
	s.Indent = "    "
	div := &p.TagNode{}
	div.Name = "div"
	div.Children = []p.Node{s}
	div.LineBreak = "\n"
	div.Id = "mother"
	div.Indent = "  "
	div.Classes = []string{"father brother sister"}
	div.Attrs = []p.Attribute{p.Attribute{"uno", &p.StaticNode{"dos"}}}
	nodes := []p.Node{div}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)
	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, "  <div class='father brother sister' id='mother' uno='dos'>\n    This is my weapon; this is my div.\n  </div>\n", output.Content)
}

func TestAutoClosingTagInXhtmlWithAttributes(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Attrs = []p.Attribute{p.Attribute{"uno", &p.StaticNode{"dos"}}}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br uno='dos' />")
}

func TestAutoClosingTagInHtml4WithAttributes(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html4"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Attrs = []p.Attribute{p.Attribute{"uno", &p.StaticNode{"dos"}}}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br uno='dos'>")
}

func TestAutoClosingTagInHtml5WithAttributes(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html5"
	opts.Autoclose = []string{"br"}
	node := &p.TagNode{}
	node.Name = "br"
	node.Attrs = []p.Attribute{p.Attribute{"uno", &p.StaticNode{"dos"}}}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<br uno='dos'>")
}

func TestTagWithOnlyTagNameAndAttributes(t *testing.T) {
	opts := CompilerOpts{}
	node := &p.TagNode{}
	node.Name = "p"
	node.Attrs = []p.Attribute{p.Attribute{"uno", &p.StaticNode{"dos"}}}
	nodes := []p.Node{node}
	pdoc := p.ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	if ok := assert.Nil(t, e); !ok {
		return
	}
	if ok := assert.NotNil(t, cdoc); !ok {
		return
	}
	if ok := assert.Equal(t, 1, len(cdoc.Outputs)); !ok {
		return
	}

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<p uno='dos'></p>")
}
