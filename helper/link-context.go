package helper

import (
	"net/url"
	"strings"

	"github.com/deeplink/core/logger"
)

type LinkContext struct {
	log logger.ILogger
}

func NewLinkContext(log logger.ILogger) LinkContext {
	return LinkContext{log}
}

func (p *LinkContext) DeepLinkToUrl(deepLink string) (uri *url.URL, err error) {
	var context IConvertor
	u, errUrl := url.Parse(deepLink)
	if err != nil {
		return uri, errUrl
	}
	q := u.Query()

	if v, ok := q["Page"]; ok && v[0] == "Product" {
		context = NewProductLinkConvertor()
	} else if v, ok := q["Page"]; ok && v[0] == "Search" {
		context = NewSearchLinkConvertor()
	} else {
		context = NewOtherLinkConvertor()
	}

	uri, err = context.DeepLinkToUrl(q)

	return
}

func (p *LinkContext) UrlToDeepLink(uri string) (deepLink *url.URL, err error) {
	var context IConvertor
	u, errUrl := url.Parse(uri)
	if err != nil {
		return deepLink, errUrl
	}

	q := u.Query()

	if strings.Contains(uri, "-p-") {
		context = NewProductLinkConvertor()
	} else if strings.Contains(uri, "q=") {
		context = NewSearchLinkConvertor()
	} else {
		context = NewOtherLinkConvertor()
	}

	deepLink, err = context.UrlToDeepLink(u.Path, q)

	return
}
