package helper

import "net/url"

type IConvertor interface {
	DeepLinkToUrl(map[string][]string) (*url.URL, error)
	UrlToDeepLink(string, map[string][]string) (*url.URL, error)
}
