package utils

import (
	"GoLog/commen"
	"strings"
	"sync"

	"github.com/88250/lute"
	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
)

// 主要对文章和评论的content进行渲染

var (
	engine *lute.Lute
	once   sync.Once
)

//首先获取markdown引擎
func getEngine() *lute.Lute {
	once.Do(func() {
		engine = lute.New(func(lute *lute.Lute) {
			lute.SetToC(true)
			lute.SetSanitize(true)
			lute.SetGFMTaskListItem(true)
		})
	})
	return engine
}

// 将markdown转化为html
func ToHtml(md string) string {
	if len(md) == 0 {
		return ""
	}
	return getEngine().MarkdownStr("", md)
}

// 主要对markdown或string转化后的html做一些后处理
func HandleHtmlContent(htmlContent string) string {
	// 首先利用bluemonday库做清理
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").OnElements("code")
	htmlContent = p.Sanitize(htmlContent)
	// 调用goquery，并对html的内容做进一步的处理
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		commen.GVA_LOG.Error("html处理错误，请检查", zap.Error(err))
	}
	//对所有a标签做处理
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		// 如果没有title，则设置title
		title := selection.AttrOr("title", "")
		if len(title) == 0 {
			selection.SetAttr("title", selection.Text())
		}
	})

	// 后续可添加对html的各种其他操作

	if htmlStr, err := doc.Find("body").Html(); err == nil {
		return htmlStr
	}
	return htmlContent
}
