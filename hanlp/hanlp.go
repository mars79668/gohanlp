package hanlp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/imroc/req"
	"github.com/xxjwxc/public/mylog"
)

type hanlp struct {
	opts Options
}

// HanLPClient build client
func HanLPClient(opts ...Option) *hanlp {
	options := Options{ // default
		URL:      "https://www.hanlp.com/api",
		Language: "zh",
	}

	for _, f := range opts { // deal option
		f(&options)
	}

	return &hanlp{
		opts: options,
	}
}

/*
分词，词性标注 。。。
Parse a piece of text.

	Args:
	    text: A document (str), or a list of sentences (List[str]).
	    tokens: A list of sentences where each sentence is a list of tokens.
	    tasks: The tasks to predict. Use ``tasks=[...]`` to run selected tasks only. Dependent tasks will be
	        automatically selected.
	    skip_tasks: The tasks to skip. Use ``skip_tasks='tok/fine'`` to enable coarse tokenization for all tasks.
	        Use ``tasks=['tok/coarse', ...]`` and ``skip_tasks='tok/fine'`` to enable coarse tokenization for
	        selected tasks.
	    language: The language of input text or tokens. ``None`` to use the default language on server.

	Returns:
	    A :class:`~hanlp_common.document.Document`.

	Examples::

	    # Use tasks=[...] to run selected tasks only
	    HanLP('晓美焰来到自然语义科技公司', tasks=['pos', 'ner'])

	    # Use skip_tasks='tok/fine' to enable coarse tokenization for all tasks
	    HanLP('晓美焰来到自然语义科技公司', skip_tasks='tok/fine')

	    # Use tasks=['tok/coarse', ...] and skip_tasks='tok/fine' to enable
	    # coarse tokenization for selected tasks
	    HanLP('晓美焰来到自然语义科技公司', tasks=['tok/coarse','pos'],skip_tasks='tok/fine')


	Raises:
	    HTTPError: Any errors happening on the Internet side or the server side. Refer to the ``code`` and ``msg``
	        of the exception for more details. A list of common errors :

	- ``400 Bad Request`` indicates that the server cannot process the request due to a client
	  fault (e.g., text too long, language unsupported).
	- ``401 Unauthorized`` indicates that the request lacks **valid** ``auth`` credentials for the API.
	- ``422 Unprocessable Entity`` indicates that the content type of the request entity is not in
	  proper json format.
	- ``429 Too Many Requests`` indicates the user has sent too many requests in a given
	  amount of time ("rate limiting").
*/
func (h *hanlp) Parse(text []string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	req := &HanReq{
		Text:      text,
		Language:  options.Language, // (zh,mnt)
		Tasks:     options.Tasks,
		SkipTasks: options.SkipTasks,
	}

	return h.Post("/parse", req, getHeader(options))
}

/*
文本纠错
Grammatical Error Correction (GEC) is the task of correcting different kinds of errors in text such as

	spelling, punctuation, grammatical, and word choice errors.

	Args:
	    text: Text potentially containing different kinds of errors such as spelling, punctuation,
	        grammatical, and word choice errors.
	    language: The language of input text. ``None`` to use the default language.

	Returns:
	    Corrected text.

	Examples::

	    HanLP.grammatical_error_correction(['每个青年都应当有远大的报复。',
	                                        '有的同学对语言很兴趣。'])
	    # Output:
	    [
	        '每个青年都应当有远大的抱负。',
	        '有的同学对语言很有兴趣。'
	    ]
*/
func (h *hanlp) GrammaticalErrorCorrection(text []string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	req := &HanReq{
		Text:     text,
		Language: options.Language, // (zh,mnt)
	}

	return h.Post("/grammatical_error_correction", req, getHeader(options))
}

/*
关键词提取
Keyphrase extraction aims to identify keywords or phrases reflecting the main topics of a document.

	Args:
	    text: The text content of the document. Preferably the concatenation of the title and the content.
	    topk: The number of top-K ranked keywords or keyphrases.
	    language: The language of input text or tokens. ``None`` to use the default language on server.

	Returns:
	    A dictionary containing each keyword or keyphrase and its ranking score :math:`s`, :math:`s \in [0, 1]`.

	Examples::

	    HanLP.keyphrase_extraction(
	        '自然语言处理是一门博大精深的学科，掌握理论才能发挥出HanLP的全部性能。 '
	        '《自然语言处理入门》是一本配套HanLP的NLP入门书，助你零起点上手自然语言处理。', topk=3)
	    # Output:
	    {'自然语言处理': 0.800000011920929,
	     'HanLP的全部性能': 0.5258446335792542,
	     '一门博大精深的学科': 0.421421080827713}
*/
func (h *hanlp) KeyphraseExtraction(text string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	if options.Topk == 0 {
		options.Topk = 10
	}

	req := &HanReq{
		Text:     text,
		Language: options.Language, // (zh,mnt)
		Topk:     options.Topk,
	}

	return h.Post("/keyphrase_extraction", req, getHeader(options))
}

/*
相似度分析
Semantic textual similarity deals with determining how similar two pieces of texts are.

	Args:
	    text: A pair or pairs of text.
	    language: The language of input text. ``None`` to use the default language.

	Returns:
	    Similarities.

	Examples::

	    HanLP.semantic_textual_similarity([
	        ('看图猜一电影名', '看图猜电影'),
	        ('无线路由器怎么无线上网', '无线上网卡和无线路由器怎么用'),
	        ('北京到上海的动车票', '上海到北京的动车票'),
	    ])
	    # Output:
	    [
	        0.9764469, # Similarity of ('看图猜一电影名', '看图猜电影')
	        0.0,       # Similarity of ('无线路由器怎么无线上网', '无线上网卡和无线路由器怎么用')
	        0.0034587  # Similarity of ('北京到上海的动车票', '上海到北京的动车票')
	    ]
*/
func (h *hanlp) SemanticTextualSimilarity(text [][]string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	if options.Topk == 0 {
		options.Topk = 10
	}

	req := &HanReq{
		Text:     text,
		Language: options.Language, // (zh,mnt)
		Topk:     options.Topk,
	}

	return h.Post("/semantic_textual_similarity", req, getHeader(options))
}

/*
文本分类
Text classification is the task of assigning a sentence or document an appropriate category.

	The categories depend on the chosen dataset and can range from topics.

	Args:
	    text: A document or a list of documents.
	    model: The model to use for prediction.
	    topk: ``True`` or ``int`` to return the top-k labels.
	    prob: Return also probabilities.

	Returns:

	    Classification results.
*/
func (h *hanlp) TextClassification(text []string, model string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}
	if model == "" {
		model = "news_zh"
	}

	req := &HanReq{
		Text:     text,
		Language: options.Language, // (zh,mnt)
		Topk:     false,
		Model:    model,
	}

	return h.Post("/text_classification", req, getHeader(options))
}

/*
情感分析
Sentiment analysis is the task of classifying the polarity of a given text. For instance,

	a text-based tweet can be categorized into either "positive", "negative", or "neutral".

	Args:
	    text: A document or a list of documents.
	    language (str): The default language for each :func:`~hanlp_restful.HanLPClient.parse` call.
	        Contact the service provider for the list of languages supported.
	        Conventionally, ``zh`` is used for Chinese and ``mul`` for multilingual.
	        Leave ``None`` to use the default language on server.

	Returns:

	    Sentiment polarity as a numerical value which measures how positive the sentiment is.

	Examples::

	    HanLP.language_identification('''“这是一部男人必看的电影。”人人都这么说。但单纯从性别区分，就会让这电影变狭隘。
	    《肖申克的救赎》突破了男人电影的局限，通篇几乎充满令人难以置信的温馨基调，而电影里最伟大的主题是“希望”。
	    当我们无奈地遇到了如同肖申克一般囚禁了心灵自由的那种囹圄，我们是无奈的老布鲁克，灰心的瑞德，还是智慧的安迪？
	    运用智慧，信任希望，并且勇敢面对恐惧心理，去打败它？
	    经典的电影之所以经典，因为他们都在做同一件事——让你从不同的角度来欣赏希望的美好。''')
	    0.9505730271339417
*/
func (h *hanlp) SentimentAnalysis(text []string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	req := &HanReq{
		Text:     text,
		Language: options.Language, // (zh,mnt)
		Topk:     false,
	}

	return h.Post("/sentiment_analysis", req, getHeader(options))
}

/*
生成式自动摘要
Abstractive Summarization is the task of generating a short and concise summary that captures the

	salient ideas of the source text. The generated summaries potentially contain new phrases and sentences that
	may not appear in the source text.

	Args:
	    text: The text content of the document.
	    language: The language of input text or tokens. ``None`` to use the default language on server.

	Returns:
	    Summarization.

	Examples::

	    HanLP.abstractive_summarization('''
	    每经AI快讯，2月4日，长江证券研究所金属行业首席分析师王鹤涛表示，2023年海外经济衰退，美债现处于历史高位，
	    黄金的趋势是值得关注的；在国内需求修复的过程中，看好大金属品种中的铜铝钢。
	    此外，在细分的小品种里，建议关注两条主线，一是新能源，比如锂、钴、镍、稀土，二是专精特新主线。（央视财经）
	    ''')
	    # Output:
	    '长江证券：看好大金属品种中的铜铝钢'
*/
func (h *hanlp) AbstractiveSummarization(text string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	req := &HanReq{
		Text:     text,
		Language: options.Language, // (zh,mnt)

	}

	return h.Post("/abstractive_summarization", req, getHeader(options))
}

/*
抽取式自动摘要
Single document summarization is the task of selecting a subset of the sentences which best

	represents a summary of the document, with a balance of salience and redundancy.

	Args:
	    text: The text content of the document.
	    topk: The maximum number of top-K ranked sentences. Note that due to Trigram Blocking tricks, the actual
	        number of returned sentences could be less than ``topk``.
	    language: The language of input text or tokens. ``None`` to use the default language on server.

	Returns:
	    A dictionary containing each sentence and its ranking score :math:`s \in [0, 1]`.

	Examples::

	    HanLP.extractive_summarization('''
	    据DigiTimes报道，在上海疫情趋缓，防疫管控开始放松后，苹果供应商广达正在逐步恢复其中国工厂的MacBook产品生产。
	    据供应链消息人士称，生产厂的订单拉动情况正在慢慢转强，这会提高MacBook Pro机型的供应量，并缩短苹果客户在过去几周所经历的延长交货时间。
	    仍有许多苹果笔记本用户在等待3月和4月订购的MacBook Pro机型到货，由于苹果的供应问题，他们的发货时间被大大推迟了。
	    据分析师郭明錤表示，广达是高端MacBook Pro的唯一供应商，自防疫封控依赖，MacBook Pro大部分型号交货时间增加了三到五周，
	    一些高端定制型号的MacBook Pro配置要到6月底到7月初才能交货。
	    尽管MacBook Pro的生产逐渐恢复，但供应问题预计依然影响2022年第三季度的产品销售。
	    苹果上周表示，防疫措施和元部件短缺将继续使其难以生产足够的产品来满足消费者的强劲需求，这最终将影响苹果6月份的收入。
	    ''')
	    # Output:
	    {'据DigiTimes报道，在上海疫情趋缓，防疫管控开始放松后，苹果供应商广达正在逐步恢复其中国工厂的MacBook产品生产。': 0.9999,
	     '仍有许多苹果笔记本用户在等待3月和4月订购的MacBook Pro机型到货，由于苹果的供应问题，他们的发货时间被大大推迟了。': 0.5800,
	     '尽管MacBook Pro的生产逐渐恢复，但供应问题预计依然影响2022年第三季度的产品销售。': 0.5422}
*/
func (h *hanlp) ExtractiveSummarization(text string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	if options.Topk == 0 {
		options.Topk = 3
	}
	req := &HanReq{
		Text:     text,
		Language: options.Language, // (zh,mnt)
		Topk:     options.Topk,
	}

	return h.Post("/extractive_summarization", req, getHeader(options))
}

/*
Text style transfer aims to change the style of the input text to the target style while preserving its

	content.

	Args:
	    text: Source text.
	    target_style: Target style.
	    language: The language of input text. ``None`` to use the default language.

	Returns:
	    Text or a list of text of the target style.

	Examples::

	    HanLP.text_style_transfer(['国家对中石油抱有很大的期望.', '要用创新去推动高质量的发展。'],
	                              target_style='gov_doc')
	    # Output:
	    [
	        '国家对中石油寄予厚望。',
	        '要以创新驱动高质量发展。'
	    ]

	    HanLP.text_style_transfer('我看到了窗户外面有白色的云和绿色的森林',
	                              target_style='modern_poetry')
	    # Output:
	    '我看见窗外的白云绿林'
*/
func (h *hanlp) TextStyleTransfer(text []string, style string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	if len(style) == 0 {
		style = "modern_poetry"
	}
	req := &HanReq{
		Text:        text,
		Language:    options.Language, // (zh,mnt)
		TargetStyle: style,
	}

	return h.Post("/text_style_transfer", req, getHeader(options))
}

func (h *hanlp) Post(uri string, hreq *HanReq, header http.Header) (string, error) {
	resp, err := req.Post(h.opts.URL+uri, req.BodyJSON(hreq), header)
	if err != nil {
		return "", err
	}

	if resp.Response().StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("HttpCode:%d\n%s", resp.Response().StatusCode, resp.String())
	}

	return resp.ToString()
}

func (h *hanlp) PostObj(uri string, hreq *HanReq, header http.Header) (*HanResp, error) {
	resp, err := req.Post(h.opts.URL+uri, req.BodyJSON(hreq), header)
	if err != nil {
		return nil, err
	}

	b, err := resp.ToBytes()
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("HttpCode:%d\n%s", resp.Response().StatusCode, resp.String())
	}

	return UnmarshalHanResp(b)
}

func (h *hanlp) Get(uri string, header http.Header) (string, error) {
	resp, err := req.Get(h.opts.URL+uri, header)
	if err != nil {
		return "", err
	}

	if resp.Response().StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("HttpCode:%d\n%s", resp.Response().StatusCode, resp.String())
	}

	return resp.ToString()
}

func (h *hanlp) About(opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	b, err := h.Get("/about", getHeader(options))
	if err != nil {
		mylog.Error(err)
		return "", err
	}

	return string(b), nil
}

// Parse parse object
func (h *hanlp) ParseObj(text []string, opts ...Option) (*HanResp, error) {
	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	req := &HanReq{
		Text:      text,
		Language:  options.Language, // (zh,mnt)
		Tasks:     options.Tasks,
		SkipTasks: options.SkipTasks,
	}

	return h.PostObj("/parse", req, getHeader(options))
}

// ParseAny parse any request parms
func (h *hanlp) ParseAny(text []string, resp interface{}, opts ...Option) error {
	reqType := reflect.TypeOf(resp)
	if reqType.Kind() != reflect.Ptr {
		return fmt.Errorf("req type not a pointer:%v", reqType)
	}

	options := h.opts
	for _, f := range opts { // option
		f(&options)
	}

	req := &HanReq{
		Text:      text,
		Language:  options.Language, // (zh,mnt)
		Tasks:     options.Tasks,
		SkipTasks: options.SkipTasks,
	}
	b, err := h.Post("/parse", req, getHeader(options))
	if err != nil {
		return err
	}

	switch v := resp.(type) {
	case *string:
		*v = b
	case *[]byte:
		*v = []byte(b)
	case *HanResp:
		tmp, e := UnmarshalHanResp([]byte(b))
		*v, err = *tmp, e
	default:
		err = json.Unmarshal([]byte(b), v)
	}

	if err != nil {
		return err
	}

	return nil
}

// marshal obj
func UnmarshalHanResp(b []byte) (*HanResp, error) {
	var hr hanResp
	err := json.Unmarshal(b, &hr)
	if err != nil {
		mylog.Error(err)
		return nil, err
	}
	resp := &HanResp{
		TokFine:   hr.TokFine,
		TokCoarse: hr.TokCoarse,
		PosCtb:    hr.PosCtb,
		PosPku:    hr.PosPku,
		Pos863:    hr.Pos863,
	}

	// ner/pku
	for _, v := range hr.NerPku {
		var tmp []NerTuple
		for _, v1 := range v {
			switch t := v1.(type) {
			case []interface{}:
				{
					tmp = append(tmp, NerTuple{
						Entity: t[0].(string),
						Type:   t[1].(string),
						Begin:  int(t[2].(float64)),
						End:    int(t[3].(float64)),
					})
				}
			default:
				mylog.Error("%v : not unmarshal", t)
			}
		}
		resp.NerPku = append(resp.NerPku, tmp)
	}
	// ----------end

	// ner/msra
	for _, v := range hr.NerMsra {
		var tmp []NerTuple
		for _, v1 := range v {
			switch t := v1.(type) {
			case []interface{}:
				{
					tmp = append(tmp, NerTuple{
						Entity: t[0].(string),
						Type:   t[1].(string),
						Begin:  int(t[2].(float64)),
						End:    int(t[3].(float64)),
					})
				}
			default:
				mylog.Error("%v : not unmarshal", t)
			}
		}
		resp.NerMsra = append(resp.NerMsra, tmp)
	}
	// ----------end

	// ner/ontonotes
	for _, v := range hr.NerOntonotes {
		var tmp []NerTuple
		for _, v1 := range v {
			switch t := v1.(type) {
			case []interface{}:
				{
					tmp = append(tmp, NerTuple{
						Entity: t[0].(string),
						Type:   t[1].(string),
						Begin:  int(t[2].(float64)),
						End:    int(t[3].(float64)),
					})
				}
			default:
				mylog.Error("%v : not unmarshal", t)
			}
		}
		resp.NerOntonotes = append(resp.NerOntonotes, tmp)
	}
	// ----------end

	// srl
	for _, v := range hr.Srl {
		var tmp [][]SrlTuple
		for _, v1 := range v {
			var tmp1 []SrlTuple
			for _, v2 := range v1 {
				switch t := v2.(type) {
				case []interface{}:
					{
						tmp1 = append(tmp1, SrlTuple{
							ArgPred: t[0].(string),
							Label:   t[1].(string),
							Begin:   int(t[2].(float64)),
							End:     int(t[3].(float64)),
						})
					}
				default:
					mylog.Error("%v : not unmarshal", t)
				}
			}
			tmp = append(tmp, tmp1)
		}
		resp.Srl = append(resp.Srl, tmp)
	}
	// -------------end

	// dep
	for _, v := range hr.Dep {
		var tmp []DepTuple
		for _, v1 := range v {
			switch t := v1.(type) {
			case []interface{}:
				{
					tmp = append(tmp, DepTuple{
						Head:     int(t[0].(float64)),
						Relation: t[1].(string),
					})
				}
			default:
				mylog.Error("%v : not unmarshal", t)
			}
		}
		resp.Dep = append(resp.Dep, tmp)
	}
	// ------------end
	// sdp
	for _, v := range hr.Sdp {
		var tmp [][]DepTuple
		for _, v1 := range v {
			var tmp1 []DepTuple
			for _, v2 := range v1 {
				switch t := v2.(type) {
				case []interface{}:
					{
						tmp1 = append(tmp1, DepTuple{
							Head:     int(t[0].(float64)),
							Relation: t[1].(string),
						})
					}
				default:
					mylog.Error("%v : not unmarshal", t)
				}
			}
			tmp = append(tmp, tmp1)
		}
		resp.Sdp = append(resp.Sdp, tmp)
	}
	// ------------end
	// Con
	resp.Con = dealCon(hr.Con)
	// ------------end

	// Con          []interface{}
	return resp, nil
}

func getHeader(opts Options) http.Header {
	header := make(http.Header)
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	if len(opts.Auth) > 0 {
		header.Add("Authorization", "Basic "+opts.Auth)
	}
	return header
}

func dealCon(info []interface{}) (re []ConTuple) {
	if len(info) == 0 {
		return nil
	}

	switch t := info[0].(type) {
	case string:
		{
			tmp1 := ConTuple{
				Key: t,
			}
			if len(info) == 2 {
				tmp1.Value = dealCon(info[1].([]interface{}))
			}
			// else { // It doesn't exist in theory
			// 	fmt.Println(info)
			// }
			re = append(re, tmp1)
		}
	case []interface{}:
		{
			for _, t1 := range info {
				tmp1 := ConTuple{}
				tmp1.Value = dealCon(t1.([]interface{}))
				re = append(re, tmp1)
			}
		}
	}

	return re
}
