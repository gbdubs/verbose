package verbose

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Verbose struct {
	log    bool
	indent int
	writer io.Writer
}

func New() Verbose {
	return Verbose{
		log:    true,
		indent: 0,
	}
}

func Empty() Verbose {
	return Verbose{
		log: false,
	}
}

func NewOrEmpty(b bool) Verbose {
	if b {
		return New()
	}
	return Empty()
}

func NewWithWriter(w io.Writer) Verbose {
	return Verbose{
		log:    true,
		indent: 0,
		writer: w,
	}
}

func (v *Verbose) indentation() string {
	s := ""
	for len(s) < v.indent {
		s += ""
	}
	return s
}

func (v *Verbose) getWriter() io.Writer {
	w := v.writer
	if w == nil {
		w = os.Stdout
	}
	return w
}

func (v *Verbose) VLog(format string, args ...interface{}) {
	if !v.log {
		return
	}
	fmt.Fprintf(v.getWriter(), v.indentation()+format, args...)
}

func (v *Verbose) VDump(i interface{}) {
	if !v.log {
		return
	}
	s := spew.Sdump(i)
	s = strings.ReplaceAll(s, "\n", "\n"+v.indentation())
	s = strings.TrimSpace(s)
	s = v.indentation() + s
	fmt.Fprintf(v.getWriter(), s)
}

func (v *Verbose) VIndent() *Verbose {
	return &Verbose{
		log:    v.log,
		indent: v.indent + 2,
		writer: v.writer,
	}
}
