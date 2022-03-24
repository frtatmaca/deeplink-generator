package helper

import (
	"net/url"

	"github.com/deeplink/config"
)

type SearchLinkConvertor struct{}

func NewSearchLinkConvertor() *SearchLinkConvertor {
	return &SearchLinkConvertor{}
}

func (s *SearchLinkConvertor) DeepLinkToUrl(q map[string][]string) (uri *url.URL, err error) {
	uri, _ = url.Parse(config.Config("BASE_URL"))
	uriQue := uri.Query()

	uri.Path = "/sr"

	if val, ok := q["Query"]; ok {
		uriQue.Set("q", val[0])
	}

	uri.RawQuery = uriQue.Encode()

	return
}

func (s *SearchLinkConvertor) UrlToDeepLink(path string, q map[string][]string) (deepLink *url.URL, err error) {
	deepLink, _ = url.Parse(config.Config("BASE_DEEPURL"))
	uriQue := deepLink.Query()

	if val, ok := q["q"]; ok {
		uriQue.Set("Page", "Search")
		uriQue.Set("Query", val[0])
	}

	deepLink.RawQuery = uriQue.Encode()

	return
}
