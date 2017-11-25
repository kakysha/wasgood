// This file is automatically generated by qtc from "rating.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line wasgood/templates/rating.qtpl:1
package templates

//line wasgood/templates/rating.qtpl:1
import (
	"github.com/gin-gonic/gin"
	"wasgood/models"
)

//line wasgood/templates/rating.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line wasgood/templates/rating.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line wasgood/templates/rating.qtpl:7
type RatingPage struct {
	Page
	Items      []*models.Item
	PageNum    int
	IsLastPage bool
}

//line wasgood/templates/rating.qtpl:15
func streamitemRatingBlock(qw422016 *qt422016.Writer, item *models.Item, class string) {
	//line wasgood/templates/rating.qtpl:15
	qw422016.N().S(`
	<div class="rating uk-width-1-6 uk-flex `)
	//line wasgood/templates/rating.qtpl:16
	qw422016.E().S(class)
	//line wasgood/templates/rating.qtpl:16
	qw422016.N().S(`" data-uk-margin data-vote-url="`)
	//line wasgood/templates/rating.qtpl:16
	streamurl(qw422016, "/items/")
	//line wasgood/templates/rating.qtpl:16
	qw422016.N().D(item.ID)
	//line wasgood/templates/rating.qtpl:16
	qw422016.N().S(`/vote/">
		<div class="minus uk-text-center">
			<button class="btn uk-button uk-icon-minus" `)
	//line wasgood/templates/rating.qtpl:18
	if item.UserVoice < 0 {
		//line wasgood/templates/rating.qtpl:18
		qw422016.N().S(`disabled`)
		//line wasgood/templates/rating.qtpl:18
	}
	//line wasgood/templates/rating.qtpl:18
	qw422016.N().S(`></button>
			<div class="count uk-text-muted uk-text-small">`)
	//line wasgood/templates/rating.qtpl:19
	qw422016.N().D(item.PlusCount - item.Rating)
	//line wasgood/templates/rating.qtpl:19
	qw422016.N().S(`</div>
		</div>
		<div class="total uk-margin-small-left
			`)
	//line wasgood/templates/rating.qtpl:22
	if item.Rating > 0 {
		//line wasgood/templates/rating.qtpl:22
		qw422016.N().S(`uk-text-success`)
		//line wasgood/templates/rating.qtpl:22
	} else if item.Rating < 0 {
		//line wasgood/templates/rating.qtpl:22
		qw422016.N().S(`uk-text-danger`)
		//line wasgood/templates/rating.qtpl:22
	}
	//line wasgood/templates/rating.qtpl:22
	qw422016.N().S(` uk-h3 uk-text-bold">`)
	//line wasgood/templates/rating.qtpl:22
	qw422016.N().D(item.Rating)
	//line wasgood/templates/rating.qtpl:22
	qw422016.N().S(`</div>
		<div class="plus uk-text-center">
			<button class="btn uk-button uk-icon-plus uk-margin-small-left" `)
	//line wasgood/templates/rating.qtpl:24
	if item.UserVoice > 0 {
		//line wasgood/templates/rating.qtpl:24
		qw422016.N().S(`disabled`)
		//line wasgood/templates/rating.qtpl:24
	}
	//line wasgood/templates/rating.qtpl:24
	qw422016.N().S(`></button>
			<div class="count uk-text-muted uk-text-small">`)
	//line wasgood/templates/rating.qtpl:25
	qw422016.N().D(item.PlusCount)
	//line wasgood/templates/rating.qtpl:25
	qw422016.N().S(`</div>
		</div>
	</div>
`)
//line wasgood/templates/rating.qtpl:28
}

//line wasgood/templates/rating.qtpl:28
func writeitemRatingBlock(qq422016 qtio422016.Writer, item *models.Item, class string) {
	//line wasgood/templates/rating.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/rating.qtpl:28
	streamitemRatingBlock(qw422016, item, class)
	//line wasgood/templates/rating.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/rating.qtpl:28
}

//line wasgood/templates/rating.qtpl:28
func itemRatingBlock(item *models.Item, class string) string {
	//line wasgood/templates/rating.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/rating.qtpl:28
	writeitemRatingBlock(qb422016, item, class)
	//line wasgood/templates/rating.qtpl:28
	qs422016 := string(qb422016.B)
	//line wasgood/templates/rating.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/rating.qtpl:28
	return qs422016
//line wasgood/templates/rating.qtpl:28
}

//line wasgood/templates/rating.qtpl:30
func streamitemBlock(qw422016 *qt422016.Writer, item *models.Item) {
	//line wasgood/templates/rating.qtpl:30
	qw422016.N().S(`
<div class="uk-grid uk-grid-small uk-flex-middle uk-overflow-hidden item `)
	//line wasgood/templates/rating.qtpl:31
	if item.Hidden {
		//line wasgood/templates/rating.qtpl:31
		qw422016.N().S(`hidden`)
		//line wasgood/templates/rating.qtpl:31
	}
	//line wasgood/templates/rating.qtpl:31
	qw422016.N().S(`" data-item-id="`)
	//line wasgood/templates/rating.qtpl:31
	qw422016.N().D(item.ID)
	//line wasgood/templates/rating.qtpl:31
	qw422016.N().S(`">
	<div class="uk-width-1-6 thumbnail">
		<a href="`)
	//line wasgood/templates/rating.qtpl:33
	streamurl(qw422016, "/liquids/")
	//line wasgood/templates/rating.qtpl:33
	qw422016.N().D(item.ID)
	//line wasgood/templates/rating.qtpl:33
	qw422016.N().S(`-`)
	//line wasgood/templates/rating.qtpl:33
	qw422016.E().S(item.Slug())
	//line wasgood/templates/rating.qtpl:33
	qw422016.N().S(`" title="`)
	//line wasgood/templates/rating.qtpl:33
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:33
	qw422016.N().S(` / `)
	//line wasgood/templates/rating.qtpl:33
	qw422016.E().S(item.Name)
	//line wasgood/templates/rating.qtpl:33
	qw422016.N().S(`">
			<img src="`)
	//line wasgood/templates/rating.qtpl:34
	if len(item.Images) > 0 {
		//line wasgood/templates/rating.qtpl:34
		qw422016.E().S(item.Images[0])
		//line wasgood/templates/rating.qtpl:34
	} else {
		//line wasgood/templates/rating.qtpl:34
		streamurl(qw422016, "/images/bottle-200x200.png")
		//line wasgood/templates/rating.qtpl:34
	}
	//line wasgood/templates/rating.qtpl:34
	qw422016.N().S(`" alt="`)
	//line wasgood/templates/rating.qtpl:34
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:34
	qw422016.N().S(` / `)
	//line wasgood/templates/rating.qtpl:34
	qw422016.E().S(item.Name)
	//line wasgood/templates/rating.qtpl:34
	qw422016.N().S(`" class="uk-border-rounded" />
		</a>
	</div>
	<div class="uk-width-3-6">
		<div>
			<a href="`)
	//line wasgood/templates/rating.qtpl:39
	streamurl(qw422016, "/liquids/")
	//line wasgood/templates/rating.qtpl:39
	qw422016.N().D(item.ID)
	//line wasgood/templates/rating.qtpl:39
	qw422016.N().S(`-`)
	//line wasgood/templates/rating.qtpl:39
	qw422016.E().S(item.Slug())
	//line wasgood/templates/rating.qtpl:39
	qw422016.N().S(`" title="`)
	//line wasgood/templates/rating.qtpl:39
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:39
	qw422016.N().S(` / `)
	//line wasgood/templates/rating.qtpl:39
	qw422016.E().S(item.Name)
	//line wasgood/templates/rating.qtpl:39
	qw422016.N().S(`" class="uk-h4 uk-text-bold no-decoration name">
				`)
	//line wasgood/templates/rating.qtpl:40
	qw422016.E().S(item.Name)
	//line wasgood/templates/rating.qtpl:40
	qw422016.N().S(`
			</a>
		</div>
		<div>
			<span class="uk-text-muted">от</span>
			<a class="brand" href="`)
	//line wasgood/templates/rating.qtpl:45
	streamurl(qw422016, "/brand/")
	//line wasgood/templates/rating.qtpl:45
	qw422016.E().S(item.Brand.Slug)
	//line wasgood/templates/rating.qtpl:45
	qw422016.N().S(`" title="`)
	//line wasgood/templates/rating.qtpl:45
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:45
	qw422016.N().S(`">
				`)
	//line wasgood/templates/rating.qtpl:46
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:46
	qw422016.N().S(`
			</a>
		</div>
		<div class="properties uk-text-muted uk-text-small">
			<a href="`)
	//line wasgood/templates/rating.qtpl:50
	streamurl(qw422016, "/liquids/")
	//line wasgood/templates/rating.qtpl:50
	qw422016.N().D(item.ID)
	//line wasgood/templates/rating.qtpl:50
	qw422016.N().S(`-`)
	//line wasgood/templates/rating.qtpl:50
	qw422016.E().S(item.Slug())
	//line wasgood/templates/rating.qtpl:50
	qw422016.N().S(`"><i class="uk-icon-commenting-o"></i> `)
	//line wasgood/templates/rating.qtpl:50
	qw422016.N().D(len(item.Reviews))
	//line wasgood/templates/rating.qtpl:50
	qw422016.N().S(`</a>
		</div>
	</div>
	<div class="uk-width-1-6"></div>
	`)
	//line wasgood/templates/rating.qtpl:54
	streamitemRatingBlock(qw422016, item, "uk-flex-top uk-flex-right")
	//line wasgood/templates/rating.qtpl:54
	qw422016.N().S(`
</div>
`)
//line wasgood/templates/rating.qtpl:56
}

//line wasgood/templates/rating.qtpl:56
func writeitemBlock(qq422016 qtio422016.Writer, item *models.Item) {
	//line wasgood/templates/rating.qtpl:56
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/rating.qtpl:56
	streamitemBlock(qw422016, item)
	//line wasgood/templates/rating.qtpl:56
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/rating.qtpl:56
}

//line wasgood/templates/rating.qtpl:56
func itemBlock(item *models.Item) string {
	//line wasgood/templates/rating.qtpl:56
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/rating.qtpl:56
	writeitemBlock(qb422016, item)
	//line wasgood/templates/rating.qtpl:56
	qs422016 := string(qb422016.B)
	//line wasgood/templates/rating.qtpl:56
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/rating.qtpl:56
	return qs422016
//line wasgood/templates/rating.qtpl:56
}

//line wasgood/templates/rating.qtpl:58
func streamdetailedItemBlock(qw422016 *qt422016.Writer, item *models.Item) {
	//line wasgood/templates/rating.qtpl:58
	qw422016.N().S(`
<div class="uk-grid uk-grid-small uk-flex-middle uk-overflow-hidden item `)
	//line wasgood/templates/rating.qtpl:59
	if item.Hidden {
		//line wasgood/templates/rating.qtpl:59
		qw422016.N().S(`hidden`)
		//line wasgood/templates/rating.qtpl:59
	}
	//line wasgood/templates/rating.qtpl:59
	qw422016.N().S(`" data-item-id="`)
	//line wasgood/templates/rating.qtpl:59
	qw422016.N().D(item.ID)
	//line wasgood/templates/rating.qtpl:59
	qw422016.N().S(`">
	<div class="uk-width-1-6 thumbnail">
		<a href="`)
	//line wasgood/templates/rating.qtpl:61
	streamurl(qw422016, "/liquids/")
	//line wasgood/templates/rating.qtpl:61
	qw422016.N().D(item.ID)
	//line wasgood/templates/rating.qtpl:61
	qw422016.N().S(`-`)
	//line wasgood/templates/rating.qtpl:61
	qw422016.E().S(item.Slug())
	//line wasgood/templates/rating.qtpl:61
	qw422016.N().S(`" title="`)
	//line wasgood/templates/rating.qtpl:61
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:61
	qw422016.N().S(` / `)
	//line wasgood/templates/rating.qtpl:61
	qw422016.E().S(item.Name)
	//line wasgood/templates/rating.qtpl:61
	qw422016.N().S(`">
			<img src="`)
	//line wasgood/templates/rating.qtpl:62
	if len(item.Images) > 0 {
		//line wasgood/templates/rating.qtpl:62
		qw422016.E().S(item.Images[0])
		//line wasgood/templates/rating.qtpl:62
	} else {
		//line wasgood/templates/rating.qtpl:62
		streamurl(qw422016, "/images/bottle-200x200.png")
		//line wasgood/templates/rating.qtpl:62
	}
	//line wasgood/templates/rating.qtpl:62
	qw422016.N().S(`" alt="`)
	//line wasgood/templates/rating.qtpl:62
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:62
	qw422016.N().S(` / `)
	//line wasgood/templates/rating.qtpl:62
	qw422016.E().S(item.Name)
	//line wasgood/templates/rating.qtpl:62
	qw422016.N().S(`" class="uk-border-rounded" />
		</a>
	</div>
	<div class="uk-width-4-6 uk-width-medium-2-6">
		<div>
			<a href="`)
	//line wasgood/templates/rating.qtpl:67
	streamurl(qw422016, "/liquids/")
	//line wasgood/templates/rating.qtpl:67
	qw422016.N().D(item.ID)
	//line wasgood/templates/rating.qtpl:67
	qw422016.N().S(`-`)
	//line wasgood/templates/rating.qtpl:67
	qw422016.E().S(item.Slug())
	//line wasgood/templates/rating.qtpl:67
	qw422016.N().S(`" title="`)
	//line wasgood/templates/rating.qtpl:67
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:67
	qw422016.N().S(` / `)
	//line wasgood/templates/rating.qtpl:67
	qw422016.E().S(item.Name)
	//line wasgood/templates/rating.qtpl:67
	qw422016.N().S(`" class="uk-h4 uk-text-bold no-decoration name">
				`)
	//line wasgood/templates/rating.qtpl:68
	qw422016.E().S(item.Name)
	//line wasgood/templates/rating.qtpl:68
	qw422016.N().S(`
			</a>
		</div>
		<div>
			<span class="uk-text-muted">от</span>
			<a class="brand" href="`)
	//line wasgood/templates/rating.qtpl:73
	streamurl(qw422016, "/brand/")
	//line wasgood/templates/rating.qtpl:73
	qw422016.E().S(item.Brand.Slug)
	//line wasgood/templates/rating.qtpl:73
	qw422016.N().S(`" title="`)
	//line wasgood/templates/rating.qtpl:73
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:73
	qw422016.N().S(` / `)
	//line wasgood/templates/rating.qtpl:73
	qw422016.E().S(item.Name)
	//line wasgood/templates/rating.qtpl:73
	qw422016.N().S(`">
				`)
	//line wasgood/templates/rating.qtpl:74
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/rating.qtpl:74
	qw422016.N().S(`
			</a>
		</div>
		<div class="properties uk-text-muted uk-text-small">
			<a href="`)
	//line wasgood/templates/rating.qtpl:78
	streamurl(qw422016, "/liquids/")
	//line wasgood/templates/rating.qtpl:78
	qw422016.N().D(item.ID)
	//line wasgood/templates/rating.qtpl:78
	qw422016.N().S(`-`)
	//line wasgood/templates/rating.qtpl:78
	qw422016.E().S(item.Slug())
	//line wasgood/templates/rating.qtpl:78
	qw422016.N().S(`"><i class="uk-icon-commenting-o"></i> `)
	//line wasgood/templates/rating.qtpl:78
	qw422016.N().D(len(item.Reviews))
	//line wasgood/templates/rating.qtpl:78
	qw422016.N().S(`</a>
		</div>
	</div>
	<div class="uk-width-2-6 uk-hidden-small">
		<div class="description">
			`)
	//line wasgood/templates/rating.qtpl:83
	qw422016.N().S(item.Description)
	//line wasgood/templates/rating.qtpl:83
	qw422016.N().S(`
		</div>
	</div>
	`)
	//line wasgood/templates/rating.qtpl:86
	streamitemRatingBlock(qw422016, item, "uk-flex-top uk-flex-right")
	//line wasgood/templates/rating.qtpl:86
	qw422016.N().S(`
</div>
`)
//line wasgood/templates/rating.qtpl:88
}

//line wasgood/templates/rating.qtpl:88
func writedetailedItemBlock(qq422016 qtio422016.Writer, item *models.Item) {
	//line wasgood/templates/rating.qtpl:88
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/rating.qtpl:88
	streamdetailedItemBlock(qw422016, item)
	//line wasgood/templates/rating.qtpl:88
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/rating.qtpl:88
}

//line wasgood/templates/rating.qtpl:88
func detailedItemBlock(item *models.Item) string {
	//line wasgood/templates/rating.qtpl:88
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/rating.qtpl:88
	writedetailedItemBlock(qb422016, item)
	//line wasgood/templates/rating.qtpl:88
	qs422016 := string(qb422016.B)
	//line wasgood/templates/rating.qtpl:88
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/rating.qtpl:88
	return qs422016
//line wasgood/templates/rating.qtpl:88
}

//line wasgood/templates/rating.qtpl:90
func StreamRating(qw422016 *qt422016.Writer, c *gin.Context, p *RatingPage) {
	//line wasgood/templates/rating.qtpl:90
	qw422016.N().S(`
`)
	//line wasgood/templates/rating.qtpl:91
	streamheader(qw422016, c, &p.Page)
	//line wasgood/templates/rating.qtpl:91
	qw422016.N().S(`
<div class="content">
	<div class="uk-panel uk-panel-box uk-panel-header rating-list">
		<div class="uk-panel-title">Рейтинг жидкостей</div>
		`)
	//line wasgood/templates/rating.qtpl:95
	for _, item := range p.Items {
		//line wasgood/templates/rating.qtpl:95
		qw422016.N().S(`
			`)
		//line wasgood/templates/rating.qtpl:96
		streamdetailedItemBlock(qw422016, item)
		//line wasgood/templates/rating.qtpl:96
		qw422016.N().S(`
		`)
		//line wasgood/templates/rating.qtpl:97
	}
	//line wasgood/templates/rating.qtpl:97
	qw422016.N().S(`
	</div>
	<div class="pagination uk-width-1-1 uk-button-group uk-margin-large" data-pagenum="`)
	//line wasgood/templates/rating.qtpl:99
	qw422016.N().D(p.PageNum)
	//line wasgood/templates/rating.qtpl:99
	qw422016.N().S(`">
		<button class="uk-button uk-width-1-3 prev" `)
	//line wasgood/templates/rating.qtpl:100
	if p.PageNum == 1 {
		//line wasgood/templates/rating.qtpl:100
		qw422016.N().S(`disabled`)
		//line wasgood/templates/rating.qtpl:100
	}
	//line wasgood/templates/rating.qtpl:100
	qw422016.N().S(`><i class="uk-icon-arrow-left"></i> предыдущая</button>
		<button class="uk-button uk-width-1-3" disabled><span class="uk-hidden-small">страница </span>`)
	//line wasgood/templates/rating.qtpl:101
	qw422016.N().D(p.PageNum)
	//line wasgood/templates/rating.qtpl:101
	qw422016.N().S(` из 224</button>
		<button class="uk-button uk-width-1-3 next" `)
	//line wasgood/templates/rating.qtpl:102
	if p.IsLastPage {
		//line wasgood/templates/rating.qtpl:102
		qw422016.N().S(`disabled`)
		//line wasgood/templates/rating.qtpl:102
	}
	//line wasgood/templates/rating.qtpl:102
	qw422016.N().S(`>следующая <i class="uk-icon-arrow-right"></i></button>
	</div>
</div>
`)
	//line wasgood/templates/rating.qtpl:105
	streamfooter(qw422016, c, &p.Page)
	//line wasgood/templates/rating.qtpl:105
	qw422016.N().S(`
`)
//line wasgood/templates/rating.qtpl:106
}

//line wasgood/templates/rating.qtpl:106
func WriteRating(qq422016 qtio422016.Writer, c *gin.Context, p *RatingPage) {
	//line wasgood/templates/rating.qtpl:106
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/rating.qtpl:106
	StreamRating(qw422016, c, p)
	//line wasgood/templates/rating.qtpl:106
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/rating.qtpl:106
}

//line wasgood/templates/rating.qtpl:106
func Rating(c *gin.Context, p *RatingPage) string {
	//line wasgood/templates/rating.qtpl:106
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/rating.qtpl:106
	WriteRating(qb422016, c, p)
	//line wasgood/templates/rating.qtpl:106
	qs422016 := string(qb422016.B)
	//line wasgood/templates/rating.qtpl:106
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/rating.qtpl:106
	return qs422016
//line wasgood/templates/rating.qtpl:106
}