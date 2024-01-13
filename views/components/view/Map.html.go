// Code generated by qtc from "Map.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Map.html:2
package view

//line views/components/view/Map.html:2
import (
	"fmt"

	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/filter"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
)

//line views/components/view/Map.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Map.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Map.html:11
func StreamMapArray(qw422016 *qt422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/view/Map.html:12
	if len(maps) == 0 {
//line views/components/view/Map.html:12
		qw422016.N().S(`<em>no results</em>`)
//line views/components/view/Map.html:14
	} else {
//line views/components/view/Map.html:14
		qw422016.N().S(`<div class="overflow full-width"><table><thead><tr>`)
//line views/components/view/Map.html:19
		for _, k := range maps[0].Keys() {
//line views/components/view/Map.html:20
			components.StreamTableHeaderSimple(qw422016, "map", k, k, "", params, nil, ps)
//line views/components/view/Map.html:21
		}
//line views/components/view/Map.html:21
		qw422016.N().S(`</tr></thead><tbody>`)
//line views/components/view/Map.html:25
		for _, m := range maps {
//line views/components/view/Map.html:25
			qw422016.N().S(`<tr>`)
//line views/components/view/Map.html:27
			for _, k := range m.Keys() {
//line views/components/view/Map.html:29
				res := ""
				switch t := m[k].(type) {
				case string:
					res = t
				case []byte:
					res = string(t)
				default:
					res = fmt.Sprint(m[k])
				}

//line views/components/view/Map.html:39
				if preserveWhitespace {
//line views/components/view/Map.html:39
					qw422016.N().S(`<td class="prews">`)
//line views/components/view/Map.html:40
					qw422016.E().S(res)
//line views/components/view/Map.html:40
					qw422016.N().S(`</td>`)
//line views/components/view/Map.html:41
				} else {
//line views/components/view/Map.html:41
					qw422016.N().S(`<td>`)
//line views/components/view/Map.html:42
					qw422016.E().S(res)
//line views/components/view/Map.html:42
					qw422016.N().S(`</td>`)
//line views/components/view/Map.html:43
				}
//line views/components/view/Map.html:44
			}
//line views/components/view/Map.html:44
			qw422016.N().S(`</tr>`)
//line views/components/view/Map.html:46
		}
//line views/components/view/Map.html:46
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/view/Map.html:50
	}
//line views/components/view/Map.html:51
}

//line views/components/view/Map.html:51
func WriteMapArray(qq422016 qtio422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/view/Map.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Map.html:51
	StreamMapArray(qw422016, maps, params, preserveWhitespace, ps)
//line views/components/view/Map.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Map.html:51
}

//line views/components/view/Map.html:51
func MapArray(maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) string {
//line views/components/view/Map.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Map.html:51
	WriteMapArray(qb422016, maps, params, preserveWhitespace, ps)
//line views/components/view/Map.html:51
	qs422016 := string(qb422016.B)
//line views/components/view/Map.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Map.html:51
	return qs422016
//line views/components/view/Map.html:51
}
