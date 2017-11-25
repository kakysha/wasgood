// This file is automatically generated by qtc from "item.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line wasgood/templates/item.qtpl:1
package templates

//line wasgood/templates/item.qtpl:1
import (
	"github.com/gin-gonic/gin"
	"wasgood/controllers/auth"
	"wasgood/models"
)

//line wasgood/templates/item.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line wasgood/templates/item.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line wasgood/templates/item.qtpl:8
type ItemPage struct {
	Page
	Item *models.Item
}

//line wasgood/templates/item.qtpl:14
func streamitemPart(qw422016 *qt422016.Writer, item *models.Item, c *gin.Context) {
	//line wasgood/templates/item.qtpl:14
	qw422016.N().S(`
<div class="uk-grid uk-grid-small `)
	//line wasgood/templates/item.qtpl:15
	if item.Hidden {
		//line wasgood/templates/item.qtpl:15
		qw422016.N().S(`hidden`)
		//line wasgood/templates/item.qtpl:15
	}
	//line wasgood/templates/item.qtpl:15
	qw422016.N().S(`">
	<ul class="uk-width-1-3 images uk-grid uk-grid-small" data-uk-grid-margin>
		`)
	//line wasgood/templates/item.qtpl:17
	if len(item.Images) > 0 {
		//line wasgood/templates/item.qtpl:17
		qw422016.N().S(`
			`)
		//line wasgood/templates/item.qtpl:18
		for _, image := range item.Images {
			//line wasgood/templates/item.qtpl:18
			qw422016.N().S(`
				<li class="uk-width-1-1 uk-width-medium-1-2"><a href="`)
			//line wasgood/templates/item.qtpl:19
			qw422016.E().S(image)
			//line wasgood/templates/item.qtpl:19
			qw422016.N().S(`" data-uk-lightbox="{group:'item-images'}"><img src="`)
			//line wasgood/templates/item.qtpl:19
			qw422016.E().S(image)
			//line wasgood/templates/item.qtpl:19
			qw422016.N().S(`"/></a></li>
			`)
			//line wasgood/templates/item.qtpl:20
		}
		//line wasgood/templates/item.qtpl:20
		qw422016.N().S(`
		`)
		//line wasgood/templates/item.qtpl:21
	} else if item.Brand.Logo != "" {
		//line wasgood/templates/item.qtpl:21
		qw422016.N().S(`
			<li class="uk-width-1-1 uk-width-medium-1-2"><img src="`)
		//line wasgood/templates/item.qtpl:22
		streamurl(qw422016, item.Brand.Logo)
		//line wasgood/templates/item.qtpl:22
		qw422016.N().S(`"/></li>
		`)
		//line wasgood/templates/item.qtpl:23
	}
	//line wasgood/templates/item.qtpl:23
	qw422016.N().S(`
	</ul>

	<div class="uk-width-3-6">
		<div class="title">
			<h1 class="uk-display-inline-block name">`)
	//line wasgood/templates/item.qtpl:28
	qw422016.E().S(item.Name)
	//line wasgood/templates/item.qtpl:28
	qw422016.N().S(`</h1>
			 от <span class="uk-h3"><a class="brand uk-text-primary" href="`)
	//line wasgood/templates/item.qtpl:29
	streamurl(qw422016, "/brand/")
	//line wasgood/templates/item.qtpl:29
	qw422016.E().S(item.Brand.Slug)
	//line wasgood/templates/item.qtpl:29
	qw422016.N().S(`">`)
	//line wasgood/templates/item.qtpl:29
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/item.qtpl:29
	qw422016.N().S(`</a></span>
			 `)
	//line wasgood/templates/item.qtpl:30
	u := models.GetUserFromContext(c)

	//line wasgood/templates/item.qtpl:30
	qw422016.N().S(`
			 `)
	//line wasgood/templates/item.qtpl:31
	if u != nil && (u.IsAdmin || auth.HasACLRights(c, item)) {
		//line wasgood/templates/item.qtpl:31
		qw422016.N().S(`[
			 	<span><a href="`)
		//line wasgood/templates/item.qtpl:32
		streamurl(qw422016, "/admin/liquids/")
		//line wasgood/templates/item.qtpl:32
		qw422016.N().D(item.ID)
		//line wasgood/templates/item.qtpl:32
		qw422016.N().S(`" target="_blank">Edit</a></span>
			 ]`)
		//line wasgood/templates/item.qtpl:33
	}
	//line wasgood/templates/item.qtpl:33
	qw422016.N().S(`
		</div>
		<div class="description">
			<p>
				`)
	//line wasgood/templates/item.qtpl:37
	if item.Description != "" {
		//line wasgood/templates/item.qtpl:37
		qw422016.N().S(`
					`)
		//line wasgood/templates/item.qtpl:38
		qw422016.N().S(item.Description)
		//line wasgood/templates/item.qtpl:38
		qw422016.N().S(`
				`)
		//line wasgood/templates/item.qtpl:39
	} else {
		//line wasgood/templates/item.qtpl:39
		qw422016.N().S(`
					`)
		//line wasgood/templates/item.qtpl:40
		qw422016.N().S(item.Brand.Description)
		//line wasgood/templates/item.qtpl:40
		qw422016.N().S(`
				`)
		//line wasgood/templates/item.qtpl:41
	}
	//line wasgood/templates/item.qtpl:41
	qw422016.N().S(`
			</p>
		</div>
	</div>

	`)
	//line wasgood/templates/item.qtpl:46
	streamitemRatingBlock(qw422016, item, "uk-flex-middle uk-flex-center")
	//line wasgood/templates/item.qtpl:46
	qw422016.N().S(`
</div>
`)
//line wasgood/templates/item.qtpl:48
}

//line wasgood/templates/item.qtpl:48
func writeitemPart(qq422016 qtio422016.Writer, item *models.Item, c *gin.Context) {
	//line wasgood/templates/item.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/item.qtpl:48
	streamitemPart(qw422016, item, c)
	//line wasgood/templates/item.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/item.qtpl:48
}

//line wasgood/templates/item.qtpl:48
func itemPart(item *models.Item, c *gin.Context) string {
	//line wasgood/templates/item.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/item.qtpl:48
	writeitemPart(qb422016, item, c)
	//line wasgood/templates/item.qtpl:48
	qs422016 := string(qb422016.B)
	//line wasgood/templates/item.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/item.qtpl:48
	return qs422016
//line wasgood/templates/item.qtpl:48
}

//line wasgood/templates/item.qtpl:50
func streamsquareItemBlock(qw422016 *qt422016.Writer, item *models.Item) {
	//line wasgood/templates/item.qtpl:50
	qw422016.N().S(`
<div class="uk-width-medium-1-6 uk-width-1-3 item `)
	//line wasgood/templates/item.qtpl:51
	if item.Hidden {
		//line wasgood/templates/item.qtpl:51
		qw422016.N().S(`hidden`)
		//line wasgood/templates/item.qtpl:51
	}
	//line wasgood/templates/item.qtpl:51
	qw422016.N().S(`" data-item-id="`)
	//line wasgood/templates/item.qtpl:51
	qw422016.N().D(item.ID)
	//line wasgood/templates/item.qtpl:51
	qw422016.N().S(`">
	<figure class="uk-overlay uk-overlay-hover">
		<img class="uk-overlay-scale" src="`)
	//line wasgood/templates/item.qtpl:53
	if len(item.Images) > 0 {
		//line wasgood/templates/item.qtpl:53
		qw422016.E().S(item.Images[0])
		//line wasgood/templates/item.qtpl:53
	} else {
		//line wasgood/templates/item.qtpl:53
		streamurl(qw422016, "/images/bottle-200x200.png")
		//line wasgood/templates/item.qtpl:53
	}
	//line wasgood/templates/item.qtpl:53
	qw422016.N().S(`" alt="`)
	//line wasgood/templates/item.qtpl:53
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/item.qtpl:53
	qw422016.N().S(` / `)
	//line wasgood/templates/item.qtpl:53
	qw422016.E().S(item.Name)
	//line wasgood/templates/item.qtpl:53
	qw422016.N().S(`">
		<div class="uk-overlay-panel uk-overlay-background uk-overlay-bottom uk-ignore">
			<p>`)
	//line wasgood/templates/item.qtpl:55
	qw422016.E().S(item.Name)
	//line wasgood/templates/item.qtpl:55
	qw422016.N().S(`: `)
	//line wasgood/templates/item.qtpl:55
	qw422016.N().D(item.Rating)
	//line wasgood/templates/item.qtpl:55
	qw422016.N().S(`</p>
		</div>
		<a class="uk-position-cover" href="`)
	//line wasgood/templates/item.qtpl:57
	streamurl(qw422016, "/liquids/")
	//line wasgood/templates/item.qtpl:57
	qw422016.N().D(item.ID)
	//line wasgood/templates/item.qtpl:57
	qw422016.N().S(`-`)
	//line wasgood/templates/item.qtpl:57
	qw422016.E().S(item.Slug())
	//line wasgood/templates/item.qtpl:57
	qw422016.N().S(`" title="`)
	//line wasgood/templates/item.qtpl:57
	qw422016.E().S(item.Brand.Name)
	//line wasgood/templates/item.qtpl:57
	qw422016.N().S(` / `)
	//line wasgood/templates/item.qtpl:57
	qw422016.E().S(item.Name)
	//line wasgood/templates/item.qtpl:57
	qw422016.N().S(`"></a>
	</figure>
</div>
`)
//line wasgood/templates/item.qtpl:60
}

//line wasgood/templates/item.qtpl:60
func writesquareItemBlock(qq422016 qtio422016.Writer, item *models.Item) {
	//line wasgood/templates/item.qtpl:60
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/item.qtpl:60
	streamsquareItemBlock(qw422016, item)
	//line wasgood/templates/item.qtpl:60
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/item.qtpl:60
}

//line wasgood/templates/item.qtpl:60
func squareItemBlock(item *models.Item) string {
	//line wasgood/templates/item.qtpl:60
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/item.qtpl:60
	writesquareItemBlock(qb422016, item)
	//line wasgood/templates/item.qtpl:60
	qs422016 := string(qb422016.B)
	//line wasgood/templates/item.qtpl:60
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/item.qtpl:60
	return qs422016
//line wasgood/templates/item.qtpl:60
}

//line wasgood/templates/item.qtpl:62
func streamreviewsBlock(qw422016 *qt422016.Writer, c *gin.Context, item *models.Item) {
	//line wasgood/templates/item.qtpl:62
	qw422016.N().S(`
	`)
	//line wasgood/templates/item.qtpl:63
	user := models.GetUserFromContext(c)

	//line wasgood/templates/item.qtpl:63
	qw422016.N().S(`
	<!-- COMMENT LIST start here 
	================================================== -->
	
	<div class="uk-panel-title">
		`)
	//line wasgood/templates/item.qtpl:68
	qw422016.N().D(len(item.Reviews))
	//line wasgood/templates/item.qtpl:68
	qw422016.N().S(` отзывов
	</div>

	<ul class="uk-comment-list">
	`)
	//line wasgood/templates/item.qtpl:72
	for _, review := range item.Reviews {
		//line wasgood/templates/item.qtpl:72
		qw422016.N().S(`
		<li>
		<article class="uk-comment `)
		//line wasgood/templates/item.qtpl:74
		if !review.Approved {
			//line wasgood/templates/item.qtpl:74
			qw422016.N().S(`unapproved`)
			//line wasgood/templates/item.qtpl:74
		}
		//line wasgood/templates/item.qtpl:74
		qw422016.N().S(` review" id="review-`)
		//line wasgood/templates/item.qtpl:74
		qw422016.N().D(review.ID)
		//line wasgood/templates/item.qtpl:74
		qw422016.N().S(`">
			<header class="uk-comment-header">
				<img class="uk-comment-avatar avatar" src="`)
		//line wasgood/templates/item.qtpl:76
		qw422016.E().S(review.Author.Photo)
		//line wasgood/templates/item.qtpl:76
		qw422016.N().S(`" alt="">
				<h4 class="uk-comment-title username">`)
		//line wasgood/templates/item.qtpl:77
		qw422016.E().S(review.Author.Name)
		//line wasgood/templates/item.qtpl:77
		qw422016.N().S(`</h4>
				<div class="uk-comment-meta">
					`)
		//line wasgood/templates/item.qtpl:79
		qw422016.E().V(review.Timestamp.Format("2006-01-02 15:04"))
		//line wasgood/templates/item.qtpl:79
		qw422016.N().S(`
					`)
		//line wasgood/templates/item.qtpl:80
		if user != nil && user.IsAdmin {
			//line wasgood/templates/item.qtpl:80
			qw422016.N().S(`
					<span class="actions">
						<a href="`)
			//line wasgood/templates/item.qtpl:82
			streamurl(qw422016, "/admin/reviews/")
			//line wasgood/templates/item.qtpl:82
			qw422016.N().D(review.ID)
			//line wasgood/templates/item.qtpl:82
			qw422016.N().S(`" class="delete">удалить</a>
						`)
			//line wasgood/templates/item.qtpl:83
			if !review.Approved {
				//line wasgood/templates/item.qtpl:83
				qw422016.N().S(`<a href="`)
				//line wasgood/templates/item.qtpl:83
				streamurl(qw422016, "/admin/reviews/")
				//line wasgood/templates/item.qtpl:83
				qw422016.N().D(review.ID)
				//line wasgood/templates/item.qtpl:83
				qw422016.N().S(`/approve" class="approve">аппрувнуть</a>`)
				//line wasgood/templates/item.qtpl:83
			}
			//line wasgood/templates/item.qtpl:83
			qw422016.N().S(`
					</span>
					`)
			//line wasgood/templates/item.qtpl:85
		}
		//line wasgood/templates/item.qtpl:85
		qw422016.N().S(`
				</div>
			</header>
			<div class="uk-comment-body text">`)
		//line wasgood/templates/item.qtpl:88
		qw422016.E().S(review.Text)
		//line wasgood/templates/item.qtpl:88
		qw422016.N().S(`</div>
		</article>
		</li>
	`)
		//line wasgood/templates/item.qtpl:91
	}
	//line wasgood/templates/item.qtpl:91
	qw422016.N().S(`
	</ul>

	`)
	//line wasgood/templates/item.qtpl:94
	if user != nil {
		//line wasgood/templates/item.qtpl:94
		qw422016.N().S(`
		<form class="uk-form" action="`)
		//line wasgood/templates/item.qtpl:95
		streamurl(qw422016, "/items/")
		//line wasgood/templates/item.qtpl:95
		qw422016.N().D(item.ID)
		//line wasgood/templates/item.qtpl:95
		qw422016.N().S(`/review?redirect=`)
		//line wasgood/templates/item.qtpl:95
		qw422016.N().U(c.Request.URL.String())
		//line wasgood/templates/item.qtpl:95
		qw422016.N().S(`" method="post" id="commentform">
			<article class="uk-comment uk-grid uk-grid-small" data-uk-grid-match>
				<header class="uk-comment-header uk-width-1-6">
					<img class="uk-comment-avatar avatar" src="`)
		//line wasgood/templates/item.qtpl:98
		qw422016.E().S(user.Photo)
		//line wasgood/templates/item.qtpl:98
		qw422016.N().S(`" alt="">
					<h4 class="uk-comment-title username">`)
		//line wasgood/templates/item.qtpl:99
		qw422016.E().S(user.Name)
		//line wasgood/templates/item.qtpl:99
		qw422016.N().S(`</h4>
				</header>
				<div class="uk-comment-body uk-width-4-6"><textarea class="uk-width-1-1" style="height: 100%" name="text" id="comment" placeholder="текст вашего отзыва..."></textarea></div>
				<input class="uk-width-1-6 uk-button" name="submit" type="submit" value="Submit"/>
			</article>
		</form>
	`)
		//line wasgood/templates/item.qtpl:105
	} else {
		//line wasgood/templates/item.qtpl:105
		qw422016.N().S(`
		<div class="uk-flex uk-flex-middle">
			<div class="uk-margin-small-right">Чтобы оставить отзыв: </div>
			<div class="login-vk uk-margin-small-right">
				<a class="uk-button uk-button-primary" href="`)
		//line wasgood/templates/item.qtpl:109
		streamurl(qw422016, "/auth/vk/login")
		//line wasgood/templates/item.qtpl:109
		qw422016.N().S(`?redirect=`)
		//line wasgood/templates/item.qtpl:109
		qw422016.N().U(c.Request.URL.String())
		//line wasgood/templates/item.qtpl:109
		qw422016.N().S(`" title="Sign in with VK">
					Войти <span class="uk-hidden-small">через</span> <i class="uk-icon-vk"></i>
				</a>
			</div> 
			<div class="login-fb">
				<a class="uk-button uk-button-primary" href="`)
		//line wasgood/templates/item.qtpl:114
		streamurl(qw422016, "/auth/fb/login")
		//line wasgood/templates/item.qtpl:114
		qw422016.N().S(`?redirect=`)
		//line wasgood/templates/item.qtpl:114
		qw422016.N().U(c.Request.URL.String())
		//line wasgood/templates/item.qtpl:114
		qw422016.N().S(`" title="Sign in with Facebook">
					Войти <span class="uk-hidden-small">через</span> <i class="uk-icon-facebook-square"></i>
				</a>
			</div>
		</div>
	`)
		//line wasgood/templates/item.qtpl:119
	}
	//line wasgood/templates/item.qtpl:119
	qw422016.N().S(`
`)
//line wasgood/templates/item.qtpl:120
}

//line wasgood/templates/item.qtpl:120
func writereviewsBlock(qq422016 qtio422016.Writer, c *gin.Context, item *models.Item) {
	//line wasgood/templates/item.qtpl:120
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/item.qtpl:120
	streamreviewsBlock(qw422016, c, item)
	//line wasgood/templates/item.qtpl:120
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/item.qtpl:120
}

//line wasgood/templates/item.qtpl:120
func reviewsBlock(c *gin.Context, item *models.Item) string {
	//line wasgood/templates/item.qtpl:120
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/item.qtpl:120
	writereviewsBlock(qb422016, c, item)
	//line wasgood/templates/item.qtpl:120
	qs422016 := string(qb422016.B)
	//line wasgood/templates/item.qtpl:120
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/item.qtpl:120
	return qs422016
//line wasgood/templates/item.qtpl:120
}

//line wasgood/templates/item.qtpl:122
func StreamItem(qw422016 *qt422016.Writer, c *gin.Context, p *ItemPage) {
	//line wasgood/templates/item.qtpl:122
	qw422016.N().S(`
`)
	//line wasgood/templates/item.qtpl:123
	streamheader(qw422016, c, &p.Page)
	//line wasgood/templates/item.qtpl:123
	qw422016.N().S(`
`)
	//line wasgood/templates/item.qtpl:124
	item := p.Item

	//line wasgood/templates/item.qtpl:124
	qw422016.N().S(`
<div class="content item" data-item-id="`)
	//line wasgood/templates/item.qtpl:125
	qw422016.N().D(item.ID)
	//line wasgood/templates/item.qtpl:125
	qw422016.N().S(`">
	<div class="uk-panel uk-panel-box">
		`)
	//line wasgood/templates/item.qtpl:127
	streamitemPart(qw422016, item, c)
	//line wasgood/templates/item.qtpl:127
	qw422016.N().S(`
	</div>
	<div class="uk-panel uk-panel-box uk-panel-header">
		<div class="uk-panel-title">Отзывы</div>
		`)
	//line wasgood/templates/item.qtpl:131
	streamreviewsBlock(qw422016, c, item)
	//line wasgood/templates/item.qtpl:131
	qw422016.N().S(`
	</div>
</div>
`)
	//line wasgood/templates/item.qtpl:134
	streamfooter(qw422016, c, &p.Page)
	//line wasgood/templates/item.qtpl:134
	qw422016.N().S(`
`)
//line wasgood/templates/item.qtpl:135
}

//line wasgood/templates/item.qtpl:135
func WriteItem(qq422016 qtio422016.Writer, c *gin.Context, p *ItemPage) {
	//line wasgood/templates/item.qtpl:135
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line wasgood/templates/item.qtpl:135
	StreamItem(qw422016, c, p)
	//line wasgood/templates/item.qtpl:135
	qt422016.ReleaseWriter(qw422016)
//line wasgood/templates/item.qtpl:135
}

//line wasgood/templates/item.qtpl:135
func Item(c *gin.Context, p *ItemPage) string {
	//line wasgood/templates/item.qtpl:135
	qb422016 := qt422016.AcquireByteBuffer()
	//line wasgood/templates/item.qtpl:135
	WriteItem(qb422016, c, p)
	//line wasgood/templates/item.qtpl:135
	qs422016 := string(qb422016.B)
	//line wasgood/templates/item.qtpl:135
	qt422016.ReleaseByteBuffer(qb422016)
	//line wasgood/templates/item.qtpl:135
	return qs422016
//line wasgood/templates/item.qtpl:135
}