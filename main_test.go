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

	// tcRes, _ := client.TextClassification([]string{`“这是一部男人必看的电影。”人人都这么说。但单纯从性别区分，就会让这电影变狭隘。
	//     《肖申克的救赎》突破了男人电影的局限，通篇几乎充满令人难以置信的温馨基调，而电影里最伟大的主题是“希望”。
	//     当我们无奈地遇到了如同肖申克一般囚禁了心灵自由的那种囹圄，我们是无奈的老布鲁克，灰心的瑞德，还是智慧的安迪？
	//     运用智慧，信任希望，并且勇敢面对恐惧心理，去打败它？
	//     经典的电影之所以经典，因为他们都在做同一件事——让你从不同的角度来欣赏希望的美好。`}, "")
	// fmt.Println(tcRes)

	// esRes, _ := client.AbstractiveSummarization(`据DigiTimes报道，在上海疫情趋缓，防疫管控开始放松后，苹果供应商广达正在逐步恢复其中国工厂的MacBook产品生产。
	//     据供应链消息人士称，生产厂的订单拉动情况正在慢慢转强，这会提高MacBook Pro机型的供应量，并缩短苹果客户在过去几周所经历的延长交货时间。
	//     仍有许多苹果笔记本用户在等待3月和4月订购的MacBook Pro机型到货，由于苹果的供应问题，他们的发货时间被大大推迟了。
	//     据分析师郭明錤表示，广达是高端MacBook Pro的唯一供应商，自防疫封控依赖，MacBook Pro大部分型号交货时间增加了三到五周，
	//     一些高端定制型号的MacBook Pro配置要到6月底到7月初才能交货。
	//     尽管MacBook Pro的生产逐渐恢复，但供应问题预计依然影响2022年第三季度的产品销售。
	//     苹果上周表示，防疫措施和元部件短缺将继续使其难以生产足够的产品来满足消费者的强劲需求，这最终将影响苹果6月份的收入。`)
	// fmt.Println(esRes)

	// asRes, _ := client.ExtractiveSummarization(`“每经AI快讯，2月4日，长江证券研究所金属行业首席分析师王鹤涛表示，2023年海外经济衰退，美债现处于历史高位，
	//     黄金的趋势是值得关注的；在国内需求修复的过程中，看好大金属品种中的铜铝钢。
	//     此外，在细分的小品种里，建议关注两条主线，一是新能源，比如锂、钴、镍、稀土，二是专精特新主线。（央视财经）`)
	// fmt.Println(asRes)

	// tstRes, _ := client.TextStyleTransfer([]string{"要以创新驱动高质量发展", "我看到了窗户外面有白色的云和绿色的森林", "国家对中石油寄予厚望"}, "modern_poetry")
	// fmt.Println(tstRes)
	govRes, _ := client.TextStyleTransfer([]string{"要以创新驱动高质量发展", "我看到了窗户外面有白色的云和绿色的森林", "国家对中石油寄予厚望"}, "gov_doc")
	fmt.Println(govRes)

	// ab, _ := client.About()
	// fmt.Println(ab)
}
