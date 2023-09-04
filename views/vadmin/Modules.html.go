// Code generated by qtc from "Modules.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Modules.html:2
package vadmin

//line views/vadmin/Modules.html:2
import (
	"runtime/debug"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/layout"
)

//line views/vadmin/Modules.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Modules.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Modules.html:11
type Modules struct {
	layout.Basic
	Modules []*debug.Module
}

//line views/vadmin/Modules.html:16
func (p *Modules) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Modules.html:16
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vadmin/Modules.html:18
	qw422016.E().S(util.AppName)
//line views/vadmin/Modules.html:18
	qw422016.N().S(` v`)
//line views/vadmin/Modules.html:18
	qw422016.E().S(as.BuildInfo.Version)
//line views/vadmin/Modules.html:18
	qw422016.N().S(`</div>
    <h3>Go Modules</h3>
    `)
//line views/vadmin/Modules.html:20
	streammoduleTable(qw422016, p.Modules)
//line views/vadmin/Modules.html:20
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/Modules.html:22
}

//line views/vadmin/Modules.html:22
func (p *Modules) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Modules.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Modules.html:22
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Modules.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Modules.html:22
}

//line views/vadmin/Modules.html:22
func (p *Modules) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Modules.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Modules.html:22
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Modules.html:22
	qs422016 := string(qb422016.B)
//line views/vadmin/Modules.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Modules.html:22
	return qs422016
//line views/vadmin/Modules.html:22
}

//line views/vadmin/Modules.html:24
func streammoduleTable(qw422016 *qt422016.Writer, mods []*debug.Module) {
//line views/vadmin/Modules.html:24
	qw422016.N().S(`
  <table class="mt">
    <thead>
      <tr>
        <th>Name</th>
        <th>Version</th>
      </tr>
    </thead>
    <tbody>
`)
//line views/vadmin/Modules.html:33
	for _, m := range mods {
//line views/vadmin/Modules.html:33
		qw422016.N().S(`      <tr>
        <td><a target="_blank" rel="noopener noreferrer" href="https://`)
//line views/vadmin/Modules.html:35
		qw422016.E().S(m.Path)
//line views/vadmin/Modules.html:35
		qw422016.N().S(`">`)
//line views/vadmin/Modules.html:35
		qw422016.E().S(m.Path)
//line views/vadmin/Modules.html:35
		qw422016.N().S(`</a></td>
        <td title="`)
//line views/vadmin/Modules.html:36
		qw422016.E().S(m.Sum)
//line views/vadmin/Modules.html:36
		qw422016.N().S(`">`)
//line views/vadmin/Modules.html:36
		qw422016.E().S(m.Version)
//line views/vadmin/Modules.html:36
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vadmin/Modules.html:38
	}
//line views/vadmin/Modules.html:38
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vadmin/Modules.html:41
}

//line views/vadmin/Modules.html:41
func writemoduleTable(qq422016 qtio422016.Writer, mods []*debug.Module) {
//line views/vadmin/Modules.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Modules.html:41
	streammoduleTable(qw422016, mods)
//line views/vadmin/Modules.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Modules.html:41
}

//line views/vadmin/Modules.html:41
func moduleTable(mods []*debug.Module) string {
//line views/vadmin/Modules.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Modules.html:41
	writemoduleTable(qb422016, mods)
//line views/vadmin/Modules.html:41
	qs422016 := string(qb422016.B)
//line views/vadmin/Modules.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Modules.html:41
	return qs422016
//line views/vadmin/Modules.html:41
}
