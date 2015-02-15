package root

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

const (
	aHrefToken                 = `href="`
	formActionToken            = `action="`
	blogBaseURL                = "http://blog.startuplive.in/"
	wordpressBeginContentToken = `<div id="content">`
	wordpressEndContentToken   = `</div><!-- #content -->`
	wordpressBeginSidebarToken = `<div id="primary" class="aside main-aside">`
	wordpressEndSidebarToken   = `</div><!-- #primary .aside -->`
	titleOpeningTag            = "<title>"
	titleClosingTag            = "</title>"
	metaDescriptionBeginToken  = `<meta name="description" content="`
	metaDescriptionEndToken    = `" />`
	publishedDateToken         = `<abbr class="published" title="`
)

func init() {
	Blog = newBlogPage(
		func(ctx *Context) (wordpressURL string) {
			wordpressURL = blogBaseURL + "?scrape=1"
			if ctx.Request.URL.RawQuery != "" {
				wordpressURL += "&" + ctx.Request.URL.RawQuery
			}
			return wordpressURL
		},
	)

	Blog_0 = newBlogPage(
		func(ctx *Context) (wordpressURL string) {
			wordpressURL = fmt.Sprintf("%s/%s/?scrape=1", blogBaseURL, ctx.URLArgs[0])
			if ctx.Request.URL.RawQuery != "" {
				wordpressURL += "&" + ctx.Request.URL.RawQuery
			}
			return wordpressURL
		},
	)

	Blog_0_1 = newBlogPage(
		func(ctx *Context) (wordpressURL string) {
			wordpressURL = fmt.Sprintf("%s/%s/%s/?scrape=1", blogBaseURL, ctx.URLArgs[0], ctx.URLArgs[1])
			if ctx.Request.URL.RawQuery != "" {
				wordpressURL += "&" + ctx.Request.URL.RawQuery
			}
			return wordpressURL
		},
	)

	Blog_0_1_2 = newBlogPage(
		func(ctx *Context) (wordpressURL string) {
			wordpressURL = fmt.Sprintf("%s/%s/%s/%s/?scrape=1", blogBaseURL, ctx.URLArgs[0], ctx.URLArgs[1], ctx.URLArgs[2])
			if ctx.Request.URL.RawQuery != "" {
				wordpressURL += "&" + ctx.Request.URL.RawQuery
			}
			return wordpressURL
		},
	)
}

func newBlogPage(getWordpressURL func(ctx *Context) (wordpressURL string)) *Page {
	return &Page{
		OnPreRender: func(page *Page, ctx *Context) (err error) {
			wordpressURL := getWordpressURL(ctx)
			r, err := http.Get(wordpressURL)
			if err != nil {
				return err
			}
			defer r.Body.Close()
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return err
			}
			if len(b) == 0 {
				return errors.New("Empty response from URL: " + wordpressURL)
			}

			ctx.Data = string(b)
			return nil
		},
		Title: Render(
			func(ctx *Context) (err error) {
				s := ctx.Data.(string)
				begin := strings.Index(s, titleOpeningTag)
				end := strings.Index(s, titleClosingTag)
				if begin == -1 || end == -1 || end <= begin {
					return errors.New("Invalid Wordpress head HTML")
				}
				begin += len(titleOpeningTag)
				ctx.Response.WriteString(s[begin:end])
				return nil
			},
		),
		MetaDescription: Render(
			func(ctx *Context) (err error) {
				s :=
					ctx.Data.(string)
				begin := strings.Index(s, metaDescriptionBeginToken)
				if begin == -1 {
					return nil // no meta description
				}
				end := strings.Index(s[begin:], metaDescriptionEndToken)
				if end == -1 {
					return errors.New("Invalid Wordpress head HTML")
				}
				end += begin
				begin += len(metaDescriptionBeginToken)
				if end <= begin {
					return errors.New("Invalid Wordpress head HTML")
				}
				ctx.Response.WriteString(s[begin:end])
				return nil
			},
		),
		Scripts: Renderers{
			SCRIPT(HackernewsCode),
			JQuery,
			StylesheetLink("/lightbox/lightbox.min.css"),
			ScriptLink("/lightbox/jquery.lightbox.min.js"),
			SCRIPT(`
				$(document).ready(function() {
					$('.blog-content img').each(function(){
						if($(this).width() > 600){
							$(this).attr('width', '600px');
							$(this).removeAttr('height');
						}
					});
				});
	   		`),
			IndirectRenderer(&Config.Page.DefaultScripts),
		},
		Content: PublicPageStructure(
			"menu-area",
			PublicPageLogo(),
			HeaderMenu(),
			DynamicView(extractWordpressContent),
			nil,
		),
	}
}

func replaceWordpressLinks(s string) string {
	s = strings.Replace(s, aHrefToken+blogBaseURL, aHrefToken+"/blog/", -1)
	s = strings.Replace(s, aHrefToken+"/blog/wp-content/", aHrefToken+blogBaseURL+"wp-content/", -1)
	s = strings.Replace(s, formActionToken+blogBaseURL, formActionToken+"/blog/", -1)
	s = strings.Replace(s, "blog.startuplive.in%2F", "startuplive.in%2Fblog%2F", -1)
	s = strings.Replace(s, `data-url="http://blog.startuplive.in`, `data-url="http://startuplive.in/blog`, -1)
	return s
}

func extractWordpressContent(ctx *Context) (View, error) {
	s :=
		ctx.Data.(string)
	begin := strings.Index(s, wordpressBeginContentToken)
	end := strings.LastIndex(s, wordpressEndContentToken)
	if begin == -1 || end == -1 || end <= begin {
		return nil, errors.New("Invalid Wordpress content HTML")
	}
	begin += len(wordpressBeginContentToken)
	content := s[begin:end]
	content = replaceWordpressLinks(content)

	begin = strings.LastIndex(s, wordpressBeginSidebarToken)
	end = strings.LastIndex(s, wordpressEndSidebarToken)
	if begin == -1 || end == -1 || end <= begin {
		return nil, errors.New("Invalid Wordpress sidebar HTML")
	}
	begin += len(wordpressBeginSidebarToken)
	sidebar := s[begin:end]
	sidebar = replaceWordpressLinks(sidebar)

	//publishedDateToken

	/*
			var smButtons Views
			if len(ctx.URLArgs) == 1 {
				//		likeButton := fmt.Sprintf(
				//			`<iframe src="//www.facebook.com/plugins/like.php?href=%s&amp;send=false&amp;layout=button_count&amp;width=50&amp;show_faces=false&amp;action=like&amp;colorscheme=light&amp;font&amp;height=21&amp;appId=248334038577169" scrolling="no" frameborder="0" style="border:none; overflow:hidden; width:50px; height:21px;" allowTransparency="true"></iframe>`,
				//			url.QueryEscape(response.URLString()),
				//		)
				//		smButtons = DIV("blog-sm-buttons", HTML(likeButton))
				fbSDK := HTML(`<div id="fb-root"></div>
		<script>(function(d, s, id) {
			var js, fjs = d.getElementsByTagName(s)[0];
			if (d.getElementById(id)) return;
			js = d.createElement(s); js.id = id;
			js.src = "//connect.facebook.net/en_US/all.js#xfbml=1&appId=248334038577169";
			fjs.parentNode.insertBefore(js, fjs);
		}(document, 'script', 'facebook-jssdk'));</script>`)
				fbLike := Printf(
					`<div class="fb-like" data-href="%s" data-send="false" data-layout="box_count" data-width="55" data-show-faces="true"></div>`,
					response.URLString(),
				)
				smButtons = append(smButtons, fbSDK, fbLike)
			}
	*/

	return Views{
		//DIV("blog-sm-buttons", smButtons),
		DIV("blog-content", HTML(content)),
		DIV("blog-sidebar", HTML(sidebar)),
		DivClearBoth(),
	}, nil
}
