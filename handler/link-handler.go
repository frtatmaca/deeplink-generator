package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/deeplink/models"
	"github.com/deeplink/service"
)

type LinkHandler struct {
	services service.LinkService
}

func NewLinkHandler(linkService service.LinkService) LinkHandler {
	return LinkHandler{linkService}
}

// DeepLinkToUrl godoc
// @Summary Deep Link To Url
// @Description Deep Link To Url
// @Tags deeplinktourl
// @Accept  json
// @Produce  json
// @Param DeepLinkToUrl body models.UriResponse true "Link Converter Input"
// @Success 200 {object} models.UriResponse
// @Router /deeplinktourl [post]
func (p *LinkHandler) DeepLinkToUrl(w http.ResponseWriter, r *http.Request) {
	request, err := getRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := p.services.DeepLinkToUrl(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	getResponse(w, response)
}

func (p *LinkHandler) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("frt frt 1")
	err := p.services.Test()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UrlToDeepLink godoc
// @Summary Url To Deep Link
// @Description Deep Link To Url
// @Tags urltodeeplink
// @Accept  json
// @Produce  json
// @Param UrlToDeepLink body models.UriResponse true "Link Converter Input"
// @Success 200 {object} models.UriResponse
// @Router /urltodeeplink [post]
func (p *LinkHandler) UrlToDeepLink(w http.ResponseWriter, r *http.Request) {
	request, err := getRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := p.services.UrlToDeepLink(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	getResponse(w, response)
}

func getRequest(r *http.Request) (models.UriRequest, error) {
	request := models.UriRequest{}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return request, err
	}

	err = json.Unmarshal(bodyBytes, &request)
	return request, err
}

func getResponse(w http.ResponseWriter, response models.UriResponse) {
	keysJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(keysJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
