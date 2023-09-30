// Code generated by qtc from "Menu.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Menu.html:2
package layout

//line views/layout/Menu.html:2
import (
	"strings"

	"github.com/kyleu/npn/app/controller/cmenu"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/menu"
	"github.com/kyleu/npn/views/components"
	"github.com/kyleu/npn/views/vutil"
)

//line views/layout/Menu.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Menu.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Menu.html:12
func StreamMenu(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:13
	if len(ps.Menu) > 0 {
//line views/layout/Menu.html:13
		qw422016.N().S(`<div class="menu-container">`)
//line views/layout/Menu.html:15
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:15
		qw422016.N().S(`<div class="menu">`)
//line views/layout/Menu.html:17
		vutil.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:17
		qw422016.N().S(`<ul class="level-0">`)
//line views/layout/Menu.html:19
		for _, i := range ps.Menu {
//line views/layout/Menu.html:20
			StreamMenuItem(qw422016, i, []string{}, ps.Breadcrumbs, 3, ps)
//line views/layout/Menu.html:21
		}
//line views/layout/Menu.html:22
		vutil.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:22
		qw422016.N().S(`</ul>`)
//line views/layout/Menu.html:24
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:24
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:26
		vutil.StreamIndent(qw422016, true, 1)
//line views/layout/Menu.html:26
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:28
	}
//line views/layout/Menu.html:29
}

//line views/layout/Menu.html:29
func WriteMenu(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:29
	StreamMenu(qw422016, ps)
//line views/layout/Menu.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:29
}

//line views/layout/Menu.html:29
func Menu(ps *cutil.PageState) string {
//line views/layout/Menu.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:29
	WriteMenu(qb422016, ps)
//line views/layout/Menu.html:29
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:29
	return qs422016
//line views/layout/Menu.html:29
}

//line views/layout/Menu.html:31
func StreamMenuItem(qw422016 *qt422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:33
	path = append(path, i.Key)
	active, final := breadcrumbs.Active(i, path)

//line views/layout/Menu.html:36
	if i.Key == "" {
//line views/layout/Menu.html:37
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:37
		qw422016.N().S(`<li class="separator"></li>`)
//line views/layout/Menu.html:39
	} else if len(i.Children) > 0 {
//line views/layout/Menu.html:40
		itemID := strings.Join(path, "--")

//line views/layout/Menu.html:41
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:42
		if active {
//line views/layout/Menu.html:42
			qw422016.N().S(`<li class="active" data-menu-key="`)
//line views/layout/Menu.html:42
			qw422016.E().S(i.Key)
//line views/layout/Menu.html:42
			qw422016.N().S(`">`)
//line views/layout/Menu.html:42
		} else {
//line views/layout/Menu.html:42
			qw422016.N().S(`<li data-menu-key="`)
//line views/layout/Menu.html:42
			qw422016.E().S(i.Key)
//line views/layout/Menu.html:42
			qw422016.N().S(`">`)
//line views/layout/Menu.html:42
		}
//line views/layout/Menu.html:43
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:43
		qw422016.N().S(`<input id="`)
//line views/layout/Menu.html:44
		qw422016.E().S(itemID)
//line views/layout/Menu.html:44
		qw422016.N().S(`-input" type="checkbox"`)
//line views/layout/Menu.html:44
		if active {
//line views/layout/Menu.html:44
			qw422016.N().S(` `)
//line views/layout/Menu.html:44
			qw422016.N().S(`checked="checked"`)
//line views/layout/Menu.html:44
		}
//line views/layout/Menu.html:44
		qw422016.N().S(` `)
//line views/layout/Menu.html:44
		qw422016.N().S(`hidden />`)
//line views/layout/Menu.html:45
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:46
		if final {
//line views/layout/Menu.html:46
			qw422016.N().S(`<label class="final" for="`)
//line views/layout/Menu.html:46
			qw422016.E().S(itemID)
//line views/layout/Menu.html:46
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:46
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:46
			qw422016.N().S(`">`)
//line views/layout/Menu.html:46
		} else {
//line views/layout/Menu.html:46
			qw422016.N().S(`<label for="`)
//line views/layout/Menu.html:46
			qw422016.E().S(itemID)
//line views/layout/Menu.html:46
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:46
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:46
			qw422016.N().S(`">`)
//line views/layout/Menu.html:46
		}
//line views/layout/Menu.html:47
		if i.Route != "" {
//line views/layout/Menu.html:48
			vutil.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:48
			qw422016.N().S(`<a class="label-link" href="`)
//line views/layout/Menu.html:49
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:49
			qw422016.N().S(`">`)
//line views/layout/Menu.html:49
			components.StreamSVGRef(qw422016, `link`, 15, 15, ``, ps)
//line views/layout/Menu.html:49
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:50
		}
//line views/layout/Menu.html:51
		components.StreamExpandCollapse(qw422016, indent+3, ps)
//line views/layout/Menu.html:52
		if i.Badge != "" {
//line views/layout/Menu.html:53
			vutil.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:53
			qw422016.N().S(`<div class="badge">`)
//line views/layout/Menu.html:54
			qw422016.E().S(i.Badge)
//line views/layout/Menu.html:54
			qw422016.N().S(`</div>`)
//line views/layout/Menu.html:55
		}
//line views/layout/Menu.html:56
		vutil.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:57
		if i.Icon != "" {
//line views/layout/Menu.html:58
			components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:59
		}
//line views/layout/Menu.html:60
		if i.Route != "" {
//line views/layout/Menu.html:60
			qw422016.N().S(`<a href="`)
//line views/layout/Menu.html:61
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:61
			qw422016.N().S(`">`)
//line views/layout/Menu.html:61
			qw422016.E().S(i.Title)
//line views/layout/Menu.html:61
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:62
		} else {
//line views/layout/Menu.html:63
			qw422016.E().S(i.Title)
//line views/layout/Menu.html:64
		}
//line views/layout/Menu.html:65
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:65
		qw422016.N().S(`</label>`)
//line views/layout/Menu.html:67
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:67
		qw422016.N().S(`<div class="menu-content level-`)
//line views/layout/Menu.html:68
		qw422016.N().D(len(path))
//line views/layout/Menu.html:68
		qw422016.N().S(`">`)
//line views/layout/Menu.html:69
		vutil.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:69
		qw422016.N().S(`<ul>`)
//line views/layout/Menu.html:71
		for _, i := range i.Children {
//line views/layout/Menu.html:72
			StreamMenuItem(qw422016, i, path, breadcrumbs, indent+3, ps)
//line views/layout/Menu.html:73
		}
//line views/layout/Menu.html:74
		vutil.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:74
		qw422016.N().S(`</ul>`)
//line views/layout/Menu.html:76
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:76
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:78
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:78
		qw422016.N().S(`</li>`)
//line views/layout/Menu.html:80
	} else {
//line views/layout/Menu.html:82
		finalClass := "item"
		if active {
			finalClass += " active"
		}
		if final {
			finalClass += " final"
		}

//line views/layout/Menu.html:90
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:90
		qw422016.N().S(`<li data-menu-key="`)
//line views/layout/Menu.html:91
		qw422016.E().S(i.Key)
//line views/layout/Menu.html:91
		qw422016.N().S(`"><a class="`)
//line views/layout/Menu.html:92
		qw422016.E().S(finalClass)
//line views/layout/Menu.html:92
		qw422016.N().S(`" href="`)
//line views/layout/Menu.html:92
		qw422016.E().S(i.Route)
//line views/layout/Menu.html:92
		qw422016.N().S(`" title="`)
//line views/layout/Menu.html:92
		qw422016.E().S(i.Desc())
//line views/layout/Menu.html:92
		qw422016.N().S(`">`)
//line views/layout/Menu.html:93
		if i.Badge != "" {
//line views/layout/Menu.html:94
			vutil.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:94
			qw422016.N().S(`<div class="badge">`)
//line views/layout/Menu.html:95
			qw422016.E().S(i.Badge)
//line views/layout/Menu.html:95
			qw422016.N().S(`</div>`)
//line views/layout/Menu.html:96
		}
//line views/layout/Menu.html:97
		if i.Icon != "" {
//line views/layout/Menu.html:98
			components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:99
		}
//line views/layout/Menu.html:100
		qw422016.E().S(i.Title)
//line views/layout/Menu.html:100
		qw422016.N().S(`</a></li>`)
//line views/layout/Menu.html:103
	}
//line views/layout/Menu.html:104
}

//line views/layout/Menu.html:104
func WriteMenuItem(qq422016 qtio422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:104
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:104
	StreamMenuItem(qw422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:104
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:104
}

//line views/layout/Menu.html:104
func MenuItem(i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) string {
//line views/layout/Menu.html:104
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:104
	WriteMenuItem(qb422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:104
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:104
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:104
	return qs422016
//line views/layout/Menu.html:104
}
