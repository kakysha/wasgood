// This file is automatically generated by qtc from "tag.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line wasgood/templates/admintemplates/tag.qtpl:1
package admintemplates

//line wasgood/templates/admintemplates/tag.qtpl:1
import (
	"github.com/gin-gonic/gin"
	"wasgood/models"
)

//line wasgood/templates/admintemplates/tag.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line wasgood/templates/admintemplates/tag.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line wasgood/templates/admintemplates/tag.qtpl:6
func StreamTag(qw422016 *qt422016.Writer, c *gin.Context, tag *models.Tag) {
	//line wasgood/templates/admintemplates/tag.qtpl:6
	qw422016.N().S(`

`)
	//line wasgood/templates/admintemplates/tag.qtpl:8
	streamheader(qw422016, c)
	//line wasgood/templates/admintemplates/tag.qtpl:8
	qw422016.N().S(`

<form data-type="json" class="form-horizontal" method="POST" enctype="multipart/form-data" action="`)
	//line wasgood/templates/admintemplates/tag.qtpl:10
	streamurl(qw422016, "/admin/tags")
	//line wasgood/templates/admintemplates/tag.qtpl:10
	qw422016.N().S(`">
	<fieldset>
		<legend>
			Tag
		</legend>
		<div class="span7">
			<div class="control-group">
				<label class="control-label">
					ID
				</label>
				<div class="controls">
					<input class="input-xlarge disabled" type="text" placeholder="`)
	//line wasgood/templates/admintemplates/tag.qtpl:21
	qw422016.N().D(tag.ID)
	//line wasgood/templates/admintemplates/tag.qtpl:21
	qw422016.N().S(`" disabled=""/>
					<input name="id" type="hidden" value="`)
	//line wasgood/templates/admintemplates/tag.qtpl:22
	qw422016.N().D(tag.ID)
	//line wasgood/templates/admintemplates/tag.qtpl:22
	qw422016.N().S(`"/>
				</div>
			</div>
			<div class="control-group">
				<label class="control-label">
					Name
				</label>
				<div class="controls">
					<input name="name" type="text" class="input-xlarge" value="`)
	//line wasgood/templates/admintemplates/tag.qtpl:30
	qw422016.E().S(tag.Name)
	//line wasgood/templates/admintemplates/tag.qtpl:30
	qw422016.N().S(`"/>
				</div>
			</div>
			<div class="control-group">
				<label class="control-label">
					Slug
				</label>
				<div class="controls">
					<input name="slug" type="text" class="input-xlarge" value="`)
	//line wasgood/templates/admintemplates/tag.qtpl:38
	qw422016.E().S(tag.Slug)
	//line wasgood/templates/admintemplates/tag.qtpl:38
	qw422016.N().S(`"/>
				</div>
			</div>
			<div class="control-group">
				<label class="control-label">
					Description
				</label>
				<div class="controls">
					<textarea name="description" id="cleditor">
						`)
	//line wasgood/templates/admintemplates/tag.qtpl:47
	qw422016.E().S(tag.Description)
	//line wasgood/templates/admintemplates/tag.qtpl:47
	qw422016.N().S(`
					</textarea>
				</div>
			</div>
			<div class="control-group images-block">
				<label class="control-label">
					Logo
				</label>
				<div class="controls span6" id="uniform-fileInput">
					<input class="input-file uniform_on" id="images" type="file" size="19"/>
					<button type="button" class="btn btn-primary upload" data-post-url="`)
	//line wasgood/templates/admintemplates/tag.qtpl:57
	streamurl(qw422016, "/admin/tags/")
	//line wasgood/templates/admintemplates/tag.qtpl:57
	qw422016.N().D(tag.ID)
	//line wasgood/templates/admintemplates/tag.qtpl:57
	qw422016.N().S(`/images?field=logo">
						Upload
					</button>
				</div>
				`)
	//line wasgood/templates/admintemplates/tag.qtpl:61
	if tag.Logo != "" {
		//line wasgood/templates/admintemplates/tag.qtpl:61
		qw422016.N().S(`
				<span class="img span2" data-src="`)
		//line wasgood/templates/admintemplates/tag.qtpl:62
		qw422016.E().S(tag.Logo)
		//line wasgood/templates/admintemplates/tag.qtpl:62
		qw422016.N().S(`">
					<img src="`)
		//line wasgood/templates/admintemplates/tag.qtpl:63
		qw422016.E().S(tag.Logo)
		//line wasgood/templates/admintemplates/tag.qtpl:63
		qw422016.N().S(`"/>
					<a class="remove" href="`)
		//line wasgood/templates/admintemplates/tag.qtpl:64
		streamurl(qw422016, "/admin/tags/")
		//line wasgood/templates/admintemplates/tag.qtpl:64
		qw422016.N().D(tag.ID)
		//line wasgood/templates/admintemplates/tag.qtpl:64
		qw422016.N().S(`/images/`)
		//line wasgood/templates/admintemplates/tag.qtpl:64
		qw422016.N().U(tag.Logo)
		//line wasgood/templates/admintemplates/tag.qtpl:64
		qw422016.N().S(`?field=logo">X</a>
				</span>
				`)
		//line wasgood/templates/admintemplates/tag.qtpl:66
	}
	//line wasgood/templates/admintemplates/tag.qtpl:66
	qw422016.N().S(`
			</div>
		</div>
		<div class="form-actions span11">
			<div class="span7">
				<button type="submit" class="btn btn-primary">
					Save changes
				</button>
			</div>
		</div>
	</fieldset>
</form>

`)
	//line wasgood/templates/admintemplates/tag.qtpl:79
	streamfooter(qw422016, c)
	//line wasgood/templates/admintemplates/tag.qtpl:79
	qw422016.N().S(`

`)
//line wasgood/templates/admintemplates/tag.qtpl:81
}

//line wasgood/templates/admintemplates/tag.qtpl:81
func WriteTag(qq422016 qtio422016.Writer, c *gin.Context, tag *models.Tag) {
	//line wasgood/templates/admintemplates/tag.qtpl:81
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/admintemplates/tag.qtpl:81
	StreamTag(qw422016, c, tag)
	//line wasgood/templates/admintemplates/tag.qtpl:81
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/admintemplates/tag.qtpl:81
}

//line wasgood/templates/admintemplates/tag.qtpl:81
func Tag(c *gin.Context, tag *models.Tag) string {
	//line wasgood/templates/admintemplates/tag.qtpl:81
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/admintemplates/tag.qtpl:81
	WriteTag(qb422016, c, tag)
	//line wasgood/templates/admintemplates/tag.qtpl:81
	qs422016 := string(qb422016.B)
	//line wasgood/templates/admintemplates/tag.qtpl:81
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/admintemplates/tag.qtpl:81
	return qs422016
//line wasgood/templates/admintemplates/tag.qtpl:81
}