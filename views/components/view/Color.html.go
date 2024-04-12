// Code generated by qtc from "Color.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Color.html:2
package view

//line views/components/view/Color.html:2
import (
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/theme"
)

//line views/components/view/Color.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Color.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Color.html:7
func StreamColor(qw422016 *qt422016.Writer, clr string, cls string, ps *cutil.PageState) {
//line views/components/view/Color.html:7
	qw422016.N().S(`<div style="background-color:`)
//line views/components/view/Color.html:8
	qw422016.E().S(clr)
//line views/components/view/Color.html:8
	qw422016.N().S(`; color:`)
//line views/components/view/Color.html:8
	qw422016.E().S(theme.TextColorFor(clr))
//line views/components/view/Color.html:8
	qw422016.N().S(`" class="`)
//line views/components/view/Color.html:8
	qw422016.E().S(cls)
//line views/components/view/Color.html:8
	qw422016.N().S(`">`)
//line views/components/view/Color.html:8
	qw422016.E().V(clr)
//line views/components/view/Color.html:8
	qw422016.N().S(`</div>`)
//line views/components/view/Color.html:9
}

//line views/components/view/Color.html:9
func WriteColor(qq422016 qtio422016.Writer, clr string, cls string, ps *cutil.PageState) {
//line views/components/view/Color.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Color.html:9
	StreamColor(qw422016, clr, cls, ps)
//line views/components/view/Color.html:9
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Color.html:9
}

//line views/components/view/Color.html:9
func Color(clr string, cls string, ps *cutil.PageState) string {
//line views/components/view/Color.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Color.html:9
	WriteColor(qb422016, clr, cls, ps)
//line views/components/view/Color.html:9
	qs422016 := string(qb422016.B)
//line views/components/view/Color.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Color.html:9
	return qs422016
//line views/components/view/Color.html:9
}
