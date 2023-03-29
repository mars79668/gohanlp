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

	return h.post("/parse", req, getHeader(options))
}

/*
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

	return h.post("/grammatical_error_correction", req, getHeader(options))
}

/*
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

	return h.post("/keyphrase_extraction", req, getHeader(options))
}

/*
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

	return h.post("/semantic_textual_similarity", req, getHeader(options))
}

func (h *hanlp) post(uri string, hreq *HanReq, header http.Header) (string, error) {
	resp, err := req.Post(h.opts.URL+uri, req.BodyJSON(hreq), header)
	if err != nil {
		return "", err
	}

	if resp.Response().StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("HttpCode:%d\n%s", resp.Response().StatusCode, resp.String())
	}

	return resp.ToString()
}

func (h *hanlp) postObj(uri string, hreq *HanReq, header http.Header) (*HanResp, error) {
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

func (h *hanlp) get(uri string, header http.Header) (string, error) {
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

	b, err := h.get("/about", getHeader(options))
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

	return h.postObj("/parse", req, getHeader(options))
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
	b, err := h.post("/parse", req, getHeader(options))
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
