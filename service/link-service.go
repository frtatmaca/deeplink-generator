package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/deeplink/constant"
	"github.com/deeplink/core/logger"
	"github.com/deeplink/helper"
	"github.com/deeplink/models"
	"github.com/deeplink/repository"
)

type LinkService struct {
	repository  repository.Repository
	linkContext helper.LinkContext
	log         logger.ILogger
}

func NewLinkService(repository repository.Repository, linkContext helper.LinkContext, log logger.ILogger) LinkService {
	return LinkService{repository, linkContext, log}
}

func (p *LinkService) DeepLinkToUrl(uri *models.UriRequest) (response models.UriResponse, err error) {
	var link models.Link
	url, err := p.linkContext.DeepLinkToUrl(uri.Link)
	if err != nil {
		p.log.Error("LinkService context DeepLinkToUrl: ", err)
		return
	}

	link.Link = url.String()
	link.DeepLink = uri.Link
	link.Type = constant.DeepLinkToUrl

	err = p.repository.Add(constant.INSERT_LINK, time.Now().String(), time.Now().String(), "", link.DeepLink, link.Link, link.Type)
	if err != nil {
		p.log.Error("LinkService repository DeepLinkToUrl add: ", err)

		return
	}

	response.Link = link.Link
	response.DeepLink = link.DeepLink

	return
}

func (p *LinkService) UrlToDeepLink(uri *models.UriRequest) (response models.UriResponse, err error) {
	var link models.Link

	url, err := p.linkContext.UrlToDeepLink(uri.Link)
	if err != nil {
		p.log.Error("LinkService context UrlToDeepLink: ", err)
		return
	}
	link.DeepLink = strings.Replace(url.String(), "ty:", "ty://", -1)
	link.Link = uri.Link
	link.Type = constant.UrlToDeepLink

	err = p.repository.Add(constant.INSERT_LINK, time.Now().String(), time.Now().String(), "", link.DeepLink, link.Link, link.Type)
	if err != nil {
		p.log.Error("LinkService repository UrlToDeepLink add: ", err)
		return
	}

	response.Link = link.Link
	response.DeepLink = link.DeepLink

	return
}

func (p *LinkService) Test() (err error) {
	fmt.Println("frt frt 2")
	test, err := p.repository.GetAll()
	fmt.Println(test)
	return
}
