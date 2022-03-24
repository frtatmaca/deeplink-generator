package testing

import (
	"bytes"
	"encoding/json"

	"github.com/deeplink/config"
	logInstance "github.com/deeplink/core/logger"
	"github.com/deeplink/handler"
	"github.com/deeplink/helper"
	"github.com/deeplink/models"
	"github.com/deeplink/repository"
	"github.com/deeplink/service"

	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	l              logInstance.ILogger   = logInstance.NewLogger()
	linkContext    helper.LinkContext    = helper.NewLinkContext(l)
	linkRepository repository.Repository = repository.NewRepository()
	linkService    service.LinkService   = service.NewLinkService(linkRepository, linkContext, l)
	linkHandler    handler.LinkHandler   = handler.NewLinkHandler(linkService)
)

func checkStatus(t *testing.T, res *httptest.ResponseRecorder, status int) {
	if status != res.Code {
		t.Error("expected", status, "got", res.Code)
	}
}

func checkEquality(t *testing.T, res *httptest.ResponseRecorder, expect models.UriResponse) {
	actual := models.UriResponse{}
	_ = json.Unmarshal(res.Body.Bytes(), &actual)

	if actual != expect {
		t.Error("expected", "500", "got", res.Code)
	}
}

func TestUrlToDeepLink(t *testing.T) {
	model := models.UriRequest{Link: "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892&merchantId=105064"}

	modelJson, _ := json.Marshal(model)
	req, err := http.NewRequest(http.MethodPost, "/urltodeeplink", bytes.NewBuffer(modelJson))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	h := http.HandlerFunc(linkHandler.UrlToDeepLink)
	h.ServeHTTP(res, req)
	checkStatus(t, res, http.StatusCreated)

	var expect models.UriResponse
	expect.DeepLink = "?CampaignId=439892&ContentId=1925865&MerchantId=105064&Page=Product"
	expect.Link = "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892&merchantId=105064"
	checkEquality(t, res, expect)
}

func TestDeepLinkToUrl(t *testing.T) {
	_ = config.SetConfig("DB_TYPE", "TEST")

	model := models.UriRequest{Link: "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064"}

	modelJson, _ := json.Marshal(model)
	req, err := http.NewRequest(http.MethodPost, "/deeplinktourl", bytes.NewBuffer(modelJson))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	h := http.HandlerFunc(linkHandler.DeepLinkToUrl)
	h.ServeHTTP(res, req)

	checkStatus(t, res, http.StatusCreated)

	var expect models.UriResponse
	expect.DeepLink = "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064"
	expect.Link = "/brand/name-p-1925865?boutiqueId=439892&merchantId=105064"
	checkEquality(t, res, expect)
}
