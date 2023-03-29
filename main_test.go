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

	stsRes, _ := client.SemanticTextualSimilarity([][]string{
		{"看图猜一电影名", "看图猜电影"},
		{"无线路由器怎么无线上网", "无线上网卡和无线路由器怎么用"},
		{"北京到上海的动车票", "上海到北京的动车票"},
	})
	fmt.Println(stsRes)

	// ab, _ := client.About()
	// fmt.Println(ab)
}
