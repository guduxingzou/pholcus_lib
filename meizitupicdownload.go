package pholcus_lib

// 基础包
import (
	"github.com/henrylee2cn/pholcus/common/goquery"                          //DOM解析
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需
	// . "github.com/henrylee2cn/pholcus/app/spider/common" //选用
	// "github.com/henrylee2cn/pholcus/logs"
	// net包
	 "net/http" //设置http.Header
	// "net/url"
	// 编码包
	// "encoding/xml"
	//"encoding/json"
	// 字符串处理包
	//"regexp"
	// "strconv"
		"strings"
	// 其他包
	 "fmt"
	// "math"
	// "time"
)

func init() {
	FileTest1.Register()
}

var FileTest1 = &Spider{
	Name:        "meizitu图片下载",//url:http://www.meizitu.com/
	Description: "meizitu图片下载",
	// Pausetime: 300,
	// Keyin:   KEYIN,
	// Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:          "http://www.meizitu.com/",
				Rule:         "meizitu",
				ConnTimeout:  -1,
				DownloaderID: 0, //图片等多媒体文件必须使用0（surfer surf go原生下载器）
			})
			
		},

		Trunk: map[string]*Rule{
			"meizitu":{
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find("#picture p a").Each(func(i int, s *goquery.Selection) {
						//s.Find('p>a').
						fmt.Println("打印一下")
						fmt.Println(s.Html())
						fmt.Println("-----------")
						url1,_:=s.Attr("href")
						fmt.Println(url1)
						fmt.Println("??????????????")
						//t:=s.Find('a').Eq(0)
						//fmt.Println("A的html",t.Html())
						if href, ok := s.Attr("href"); ok {
								ctx.AddQueue(&request.Request{
									Url:    href,
									Header: http.Header{"Content-Type": []string{"text/html; charset=gbk"}},
									Rule:   "图片URL",
								})
							}
						})
				},
			},

			"图片URL": {
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find("#picture p img").Each(func(i int, s *goquery.Selection) {
						//s.Find('p>a').
						fmt.Println("二级页面开始！打印一下????????????????????????????????????")
						fmt.Println(s.Html())
						fmt.Println("-----------")
						url1,_:=s.Attr("src")
						fmt.Println(url1)
						fmt.Println("??????????????")
						fmt.Println("二级页面结束！")
						//t:=s.Find('a').Eq(0)
						//fmt.Println("A的html",t.Html())
						if href, ok := s.Attr("src"); ok {
								ctx.AddQueue(&request.Request{
									Url:    href,
									Header: http.Header{"Content-Type": []string{"text/html; charset=gbk"}},
									Rule:   "图片下载",
								})
							}
						
						})
				},
			},
			"图片下载": {
				ParseFunc: func(ctx *Context) {
					fmt.Println("图片链接URL：",ctx.GetUrl())
					picurl:=ctx.GetUrl()
					picname:=strings.Replace(picurl,"http://mm.howkuai.com/wp-content/uploads/","",-1)
					picname=strings.Replace(picname,"/","-",-1)
					ctx.FileOutput(picname) // 等价于ctx.AddFile("baidu")
				},
			},
			
		},
	},
}
