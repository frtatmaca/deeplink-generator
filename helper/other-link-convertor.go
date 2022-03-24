package helper

import (
	"net/url"

	"github.com/deeplink/config"
)

type OtherLinkConvertor struct{}

func NewOtherLinkConvertor() *OtherLinkConvertor {
	return &OtherLinkConvertor{}
}

func (s *OtherLinkConvertor) DeepLinkToUrl(q map[string][]string) (uri *url.URL, err error) {
	uri, _ = url.Parse(config.Config("BASE_URL"))

	return
}

func (s *OtherLinkConvertor) UrlToDeepLink(path string, q map[string][]string) (deepLink *url.URL, err error) {
	deepLink, _ = url.Parse(config.Config("BASE_DEEPURL"))
	uriQue := deepLink.Query()

	uriQue.Set("Page", "Home")

	deepLink.RawQuery = uriQue.Encode()

	return
}
