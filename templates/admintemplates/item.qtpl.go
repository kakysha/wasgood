// This file is automatically generated by qtc from "item.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line wasgood/templates/admintemplates/item.qtpl:1
package admintemplates

//line wasgood/templates/admintemplates/item.qtpl:1
import (
	"github.com/gin-gonic/gin"
	"wasgood/models"
)

//line wasgood/templates/admintemplates/item.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line wasgood/templates/admintemplates/item.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line wasgood/templates/admintemplates/item.qtpl:6
func streamitemPart(qw422016 *qt422016.Writer, item *models.Item) {
	//line wasgood/templates/admintemplates/item.qtpl:6
	qw422016.N().S(`
<!-- images -->
`)
	//line wasgood/templates/admintemplates/item.qtpl:8
	if item.ID != 0 {
		//line wasgood/templates/admintemplates/item.qtpl:8
		qw422016.N().S(`
<div class="span4 pull-right images-block">
	<div class="control-group">
		<div class="images sortable" data-sort-url="`)
		//line wasgood/templates/admintemplates/item.qtpl:11
		streamurl(qw422016, "/admin/items/")
		//line wasgood/templates/admintemplates/item.qtpl:11
		qw422016.N().D(item.ID)
		//line wasgood/templates/admintemplates/item.qtpl:11
		qw422016.N().S(`/images/sort">
			`)
		//line wasgood/templates/admintemplates/item.qtpl:12
		for _, src := range item.Images {
			//line wasgood/templates/admintemplates/item.qtpl:12
			qw422016.N().S(`
				<span class="img" data-src="`)
			//line wasgood/templates/admintemplates/item.qtpl:13
			qw422016.E().S(src)
			//line wasgood/templates/admintemplates/item.qtpl:13
			qw422016.N().S(`">
					<img src="`)
			//line wasgood/templates/admintemplates/item.qtpl:14
			qw422016.E().S(src)
			//line wasgood/templates/admintemplates/item.qtpl:14
			qw422016.N().S(`"/>
					<a class="remove" href="`)
			//line wasgood/templates/admintemplates/item.qtpl:15
			streamurl(qw422016, "/admin/items")
			//line wasgood/templates/admintemplates/item.qtpl:15
			qw422016.N().S(`/`)
			//line wasgood/templates/admintemplates/item.qtpl:15
			qw422016.N().D(item.ID)
			//line wasgood/templates/admintemplates/item.qtpl:15
			qw422016.N().S(`/images/`)
			//line wasgood/templates/admintemplates/item.qtpl:15
			qw422016.N().U(src)
			//line wasgood/templates/admintemplates/item.qtpl:15
			qw422016.N().S(`">X</a>
				</span>
			`)
			//line wasgood/templates/admintemplates/item.qtpl:17
		}
		//line wasgood/templates/admintemplates/item.qtpl:17
		qw422016.N().S(`
		</div>
		<div class="uploader" id="uniform-fileInput">
			<input class="input-file uniform_on" id="images" type="file" size="19" multiple="multiple"/>
		</div>
		<button type="button" class="btn btn-primary upload" data-post-url="`)
		//line wasgood/templates/admintemplates/item.qtpl:22
		streamurl(qw422016, "/admin/items/")
		//line wasgood/templates/admintemplates/item.qtpl:22
		qw422016.N().D(item.ID)
		//line wasgood/templates/admintemplates/item.qtpl:22
		qw422016.N().S(`/images">
				Upload
		</button>
	</div>
</div>
`)
		//line wasgood/templates/admintemplates/item.qtpl:27
	}
	//line wasgood/templates/admintemplates/item.qtpl:27
	qw422016.N().S(`
<div class="span7">
	<div class="control-group">
		<label class="control-label">ID</label>
		<div class="controls">
			<input class="input-xlarge disabled" type="text" placeholder="`)
	//line wasgood/templates/admintemplates/item.qtpl:32
	qw422016.N().D(item.ID)
	//line wasgood/templates/admintemplates/item.qtpl:32
	qw422016.N().S(`" disabled="">
			<input name="id" type="hidden" value="`)
	//line wasgood/templates/admintemplates/item.qtpl:33
	qw422016.N().D(item.ID)
	//line wasgood/templates/admintemplates/item.qtpl:33
	qw422016.N().S(`" >
		</div>
	</div>
			
	<div class="control-group">
		<label class="control-label">
			<a href="`)
	//line wasgood/templates/admintemplates/item.qtpl:39
	streamurl(qw422016, "/admin/tags")
	//line wasgood/templates/admintemplates/item.qtpl:39
	qw422016.N().S(`">Brand</a>
		</label>
		<div class="controls">
			<select data-rel="chosen" name="brand" class="span7">
				`)
	//line wasgood/templates/admintemplates/item.qtpl:43
	for id, brand := range models.GetAllTagsForField("brand") {
		//line wasgood/templates/admintemplates/item.qtpl:43
		qw422016.N().S(`
				<option value="`)
		//line wasgood/templates/admintemplates/item.qtpl:44
		qw422016.N().D(id)
		//line wasgood/templates/admintemplates/item.qtpl:44
		qw422016.N().S(`"
				`)
		//line wasgood/templates/admintemplates/item.qtpl:45
		if item.Brand.ID == id {
			//line wasgood/templates/admintemplates/item.qtpl:45
			qw422016.N().S(`selected`)
			//line wasgood/templates/admintemplates/item.qtpl:45
		}
		//line wasgood/templates/admintemplates/item.qtpl:45
		qw422016.N().S(`
				>`)
		//line wasgood/templates/admintemplates/item.qtpl:46
		qw422016.E().S(brand.Name)
		//line wasgood/templates/admintemplates/item.qtpl:46
		qw422016.N().S(`
				</option>
				`)
		//line wasgood/templates/admintemplates/item.qtpl:48
	}
	//line wasgood/templates/admintemplates/item.qtpl:48
	qw422016.N().S(`
			</select>
		</div>
	</div>
	
	<div class="control-group">
		<label class="control-label">Name</label>
		<div class="controls">
			<input name="name:string" type="text" class="input-xlarge" value="`)
	//line wasgood/templates/admintemplates/item.qtpl:56
	qw422016.E().S(item.Name)
	//line wasgood/templates/admintemplates/item.qtpl:56
	qw422016.N().S(`" required>
		</div>
	</div>
		

	<div class="control-group">
		<label class="control-label">Description</label>
		<div class="controls">
			<textarea name="description" id="cleditor">`)
	//line wasgood/templates/admintemplates/item.qtpl:64
	qw422016.E().S(item.Description)
	//line wasgood/templates/admintemplates/item.qtpl:64
	qw422016.N().S(`</textarea>
		</div>
	</div>

	<div class="control-group">
		<label class="control-label">Hidden</label>
		<div class="controls">
			<input name="hidden" `)
	//line wasgood/templates/admintemplates/item.qtpl:71
	if item.Hidden {
		//line wasgood/templates/admintemplates/item.qtpl:71
		qw422016.N().S(`checked`)
		//line wasgood/templates/admintemplates/item.qtpl:71
	}
	//line wasgood/templates/admintemplates/item.qtpl:71
	qw422016.N().S(` data-no-uniform="true" type="checkbox" class="iphone-toggle" value="true"> <!-- value is just a text that will be sent as param value when this is checked -->
		</div>
	</div>
</div>
`)
//line wasgood/templates/admintemplates/item.qtpl:75
}

//line wasgood/templates/admintemplates/item.qtpl:75
func writeitemPart(qq422016 qtio422016.Writer, item *models.Item) {
	//line wasgood/templates/admintemplates/item.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/admintemplates/item.qtpl:75
	streamitemPart(qw422016, item)
	//line wasgood/templates/admintemplates/item.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/admintemplates/item.qtpl:75
}

//line wasgood/templates/admintemplates/item.qtpl:75
func itemPart(item *models.Item) string {
	//line wasgood/templates/admintemplates/item.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/admintemplates/item.qtpl:75
	writeitemPart(qb422016, item)
	//line wasgood/templates/admintemplates/item.qtpl:75
	qs422016 := string(qb422016.B)
	//line wasgood/templates/admintemplates/item.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/admintemplates/item.qtpl:75
	return qs422016
//line wasgood/templates/admintemplates/item.qtpl:75
}

//line wasgood/templates/admintemplates/item.qtpl:77
func StreamItem(qw422016 *qt422016.Writer, c *gin.Context, item *models.Item) {
	//line wasgood/templates/admintemplates/item.qtpl:77
	qw422016.N().S(`

`)
	//line wasgood/templates/admintemplates/item.qtpl:79
	streamheader(qw422016, c)
	//line wasgood/templates/admintemplates/item.qtpl:79
	qw422016.N().S(`

<form data-type="json" class="form-horizontal" method="POST" enctype="multipart/form-data">
	<fieldset>
		<legend>
			Item
		</legend>
		`)
	//line wasgood/templates/admintemplates/item.qtpl:86
	streamitemPart(qw422016, item)
	//line wasgood/templates/admintemplates/item.qtpl:86
	qw422016.N().S(`
		
		<div class="form-actions span11">
			<div class="span7">
				<button type="submit" class="btn btn-primary">
					Save changes
				</button>
			</div>
			<div class="span2">
				<span class="icon-arrow-left">
				</span>
				<a href="`)
	//line wasgood/templates/admintemplates/item.qtpl:97
	streamurl(qw422016, "/admin/items/")
	//line wasgood/templates/admintemplates/item.qtpl:97
	qw422016.N().D(item.ID - 1)
	//line wasgood/templates/admintemplates/item.qtpl:97
	qw422016.N().S(`">
					Prev
				</a>
			</div>
			<div class="span2">
				<a href="`)
	//line wasgood/templates/admintemplates/item.qtpl:102
	streamurl(qw422016, "/admin/items/")
	//line wasgood/templates/admintemplates/item.qtpl:102
	qw422016.N().D(item.ID + 1)
	//line wasgood/templates/admintemplates/item.qtpl:102
	qw422016.N().S(`">
					Next
				</a>
				<span class="icon-arrow-right">
				</span>
			</div>
		</div>
	</fieldset>
</form>

`)
	//line wasgood/templates/admintemplates/item.qtpl:112
	streamfooter(qw422016, c)
	//line wasgood/templates/admintemplates/item.qtpl:112
	qw422016.N().S(`

`)
//line wasgood/templates/admintemplates/item.qtpl:114
}

//line wasgood/templates/admintemplates/item.qtpl:114
func WriteItem(qq422016 qtio422016.Writer, c *gin.Context, item *models.Item) {
	//line wasgood/templates/admintemplates/item.qtpl:114
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/admintemplates/item.qtpl:114
	StreamItem(qw422016, c, item)
	//line wasgood/templates/admintemplates/item.qtpl:114
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/admintemplates/item.qtpl:114
}

//line wasgood/templates/admintemplates/item.qtpl:114
func Item(c *gin.Context, item *models.Item) string {
	//line wasgood/templates/admintemplates/item.qtpl:114
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/admintemplates/item.qtpl:114
	WriteItem(qb422016, c, item)
	//line wasgood/templates/admintemplates/item.qtpl:114
	qs422016 := string(qb422016.B)
	//line wasgood/templates/admintemplates/item.qtpl:114
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/admintemplates/item.qtpl:114
	return qs422016
//line wasgood/templates/admintemplates/item.qtpl:114
}
