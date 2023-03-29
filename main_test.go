package gohanlp

import (
	"fmt"
	"testing"

	"github.com/hankcs/gohanlp/hanlp"
)

// TestMain_test .
func TestMain_test(t *testing.T) {
	client := hanlp.HanLPClient(hanlp.WithAuth("")) // auth

	// s, _ := client.Parse([]string{"2021年HanLPv2.1为生产环境带来次世代最先进的多语种NLP技术。阿婆主来到北京立方庭参观自然语义科技公司。",
	// 	"尊敬的匿名用户，你的调用次数超过了每分钟2次"},
	// 	hanlp.WithLanguage("zh"))
	// fmt.Println(s)

	// resp, _ := client.ParseObj("2021年HanLPv2.1为生产环境带来次世代最先进的多语种NLP技术。阿婆主来到北京立方庭参观自然语义科技公司。",
	// 	hanlp.WithLanguage("zh"))
	// fmt.Println(resp)

	// gecRes, _ := client.KeyphraseExtraction("2021年HanLPv2.1为生产环境待来次世代最先进的多语种NLP技术。阿婆主来到北京立方庭参观自然语义科举公司")
	// fmt.Println(gecRes)

	// stsRes, _ := client.SemanticTextualSimilarity([][]string{
	// 	{"看图猜一电影名", "看图猜电影"},
	// 	{"无线路由器怎么无线上网", "无线上网卡和无线路由器怎么用"},
	// 	{"北京到上海的动车票", "上海到北京的动车票"},
	// })
	// fmt.Println(stsRes)

	// saRes, _ := client.SentimentAnalysis([]string{`“这是一部男人必看的电影。”人人都这么说。但单纯从性别区分，就会让这电影变狭隘。
	//     《肖申克的救赎》突破了男人电影的局限，通篇几乎充满令人难以置信的温馨基调，而电影里最伟大的主题是“希望”。
	//     当我们无奈地遇到了如同肖申克一般囚禁了心灵自由的那种囹圄，我们是无奈的老布鲁克，灰心的瑞德，还是智慧的安迪？
	//     运用智慧，信任希望，并且勇敢面对恐惧心理，去打败它？
	//     经典的电影之所以经典，因为他们都在做同一件事——让你从不同的角度来欣赏希望的美好。`})
	// fmt.Println(saRes)

	tcRes, _ := client.TextClassification([]string{`“这是一部男人必看的电影。”人人都这么说。但单纯从性别区分，就会让这电影变狭隘。
	    《肖申克的救赎》突破了男人电影的局限，通篇几乎充满令人难以置信的温馨基调，而电影里最伟大的主题是“希望”。
	    当我们无奈地遇到了如同肖申克一般囚禁了心灵自由的那种囹圄，我们是无奈的老布鲁克，灰心的瑞德，还是智慧的安迪？
	    运用智慧，信任希望，并且勇敢面对恐惧心理，去打败它？
	    经典的电影之所以经典，因为他们都在做同一件事——让你从不同的角度来欣赏希望的美好。`}, "")
	fmt.Println(tcRes)

	// ab, _ := client.About()
	// fmt.Println(ab)
}
