// Code generated by qtc from "Authentication.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vauth/Authentication.html:2
package vauth

//line views/vauth/Authentication.html:2
import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/auth"
	"github.com/kyleu/npn/views/components"
)

//line views/vauth/Authentication.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vauth/Authentication.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vauth/Authentication.html:9
func StreamAuthentication(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vauth/Authentication.html:9
	qw422016.N().S(`
`)
//line views/vauth/Authentication.html:11
	prvs, err := as.Auth.Providers(ps.Logger)
	if err != nil {
		panic(err)
	}

//line views/vauth/Authentication.html:15
	qw422016.N().S(`  <div class="card">
    <div class="right"><a href="#modal-available"><button type="button">Available</button></a></div>
    <h3>`)
//line views/vauth/Authentication.html:18
	components.StreamSVGIcon(qw422016, `profile`, ps)
//line views/vauth/Authentication.html:18
	qw422016.N().S(` Authentication</h3>
`)
//line views/vauth/Authentication.html:19
	if len(prvs) == 0 {
//line views/vauth/Authentication.html:19
		qw422016.N().S(`    <em class="mt">no authentication providers configured, why not <a href="#modal-available">add one</a>?</em>
`)
//line views/vauth/Authentication.html:21
	} else {
//line views/vauth/Authentication.html:21
		qw422016.N().S(`    <table class="mt">
`)
//line views/vauth/Authentication.html:23
		for _, prv := range prvs {
//line views/vauth/Authentication.html:23
			qw422016.N().S(`      <tr><td><a href="/auth/`)
//line views/vauth/Authentication.html:24
			qw422016.N().U(prv.ID)
//line views/vauth/Authentication.html:24
			qw422016.N().S(`?refer=`)
//line views/vauth/Authentication.html:24
			qw422016.N().U(`/admin`)
//line views/vauth/Authentication.html:24
			qw422016.N().S(`">`)
//line views/vauth/Authentication.html:24
			qw422016.E().S(auth.AvailableProviderNames[prv.ID])
//line views/vauth/Authentication.html:24
			qw422016.N().S(`</a></td></tr>
`)
//line views/vauth/Authentication.html:25
		}
//line views/vauth/Authentication.html:25
		qw422016.N().S(`    </table>
`)
//line views/vauth/Authentication.html:27
	}
//line views/vauth/Authentication.html:27
	qw422016.N().S(`  </div>

  <div id="modal-available" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Available Authentication Providers</h2>
      </div>
      <div class="modal-body">
        <table>
`)
//line views/vauth/Authentication.html:39
	for _, x := range auth.AvailableProviderKeys {
//line views/vauth/Authentication.html:39
		qw422016.N().S(`          <tr>
`)
//line views/vauth/Authentication.html:41
		if prvs.Contains(x) {
//line views/vauth/Authentication.html:41
			qw422016.N().S(`            <td class="nowrap"><em>`)
//line views/vauth/Authentication.html:42
			qw422016.E().S(auth.AvailableProviderNames[x])
//line views/vauth/Authentication.html:42
			qw422016.N().S(`</em></td>
            <td><em>`)
//line views/vauth/Authentication.html:43
			qw422016.E().S(auth.ProviderUsage(x, prvs.Contains(x)))
//line views/vauth/Authentication.html:43
			qw422016.N().S(`</em></td>
`)
//line views/vauth/Authentication.html:44
		} else {
//line views/vauth/Authentication.html:44
			qw422016.N().S(`            <td class="nowrap">`)
//line views/vauth/Authentication.html:45
			qw422016.E().S(auth.AvailableProviderNames[x])
//line views/vauth/Authentication.html:45
			qw422016.N().S(`</td>
            <td>`)
//line views/vauth/Authentication.html:46
			qw422016.E().S(auth.ProviderUsage(x, prvs.Contains(x)))
//line views/vauth/Authentication.html:46
			qw422016.N().S(`</td>
`)
//line views/vauth/Authentication.html:47
		}
//line views/vauth/Authentication.html:47
		qw422016.N().S(`          </tr>
`)
//line views/vauth/Authentication.html:49
	}
//line views/vauth/Authentication.html:49
	qw422016.N().S(`        </table>
      </div>
    </div>
  </div>
`)
//line views/vauth/Authentication.html:54
}

//line views/vauth/Authentication.html:54
func WriteAuthentication(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vauth/Authentication.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vauth/Authentication.html:54
	StreamAuthentication(qw422016, as, ps)
//line views/vauth/Authentication.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vauth/Authentication.html:54
}

//line views/vauth/Authentication.html:54
func Authentication(as *app.State, ps *cutil.PageState) string {
//line views/vauth/Authentication.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vauth/Authentication.html:54
	WriteAuthentication(qb422016, as, ps)
//line views/vauth/Authentication.html:54
	qs422016 := string(qb422016.B)
//line views/vauth/Authentication.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vauth/Authentication.html:54
	return qs422016
//line views/vauth/Authentication.html:54
}
