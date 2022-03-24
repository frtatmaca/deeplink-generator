package helper

import (
	"net/url"
	"strings"

	"github.com/deeplink/config"
)

type ProductLinkConvertor struct{}

func NewProductLinkConvertor() *ProductLinkConvertor {
	return &ProductLinkConvertor{}
}

func (s *ProductLinkConvertor) DeepLinkToUrl(q map[string][]string) (uri *url.URL, err error) {
	uri, _ = url.Parse(config.Config("BASE_URL"))
	uriQue := uri.Query()

	if val, ok := q["ContentId"]; ok {
		uri.Path = "/brand/name-p-" + val[0]
	}

	if val, ok := q["CampaignId"]; ok {
		uriQue.Set("boutiqueId", val[0])
	}

	if val, ok := q["MerchantId"]; ok {
		uriQue.Set("merchantId", val[0])
	}
	uri.RawQuery = uriQue.Encode()

	return
}

func (s *ProductLinkConvertor) UrlToDeepLink(path string, q map[string][]string) (deepLink *url.URL, err error) {
	deepLink, _ = url.Parse(config.Config("BASE_DEEPURL"))
	uriQue := deepLink.Query()

	productId := path[strings.Index(path, "-p-")+3 : len(path)]

	uriQue.Set("Page", "Product")
	uriQue.Set("ContentId", productId)

	if val, ok := q["boutiqueId"]; ok {
		uriQue.Set("CampaignId", val[0])
	}

	if val, ok := q["merchantId"]; ok {
		uriQue.Set("MerchantId", val[0])
	}

	deepLink.RawQuery = uriQue.Encode()

	return
}
