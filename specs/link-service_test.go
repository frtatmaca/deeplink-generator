package testing

import (
	"testing"

	"github.com/deeplink/models"
)

func TestUrlToDeepLinkService(t *testing.T) {
	model := models.UriRequest{Link: "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892&merchantId=105064"}

	actual, _ := linkService.UrlToDeepLink(&model)
	var expect models.UriResponse
	expect.DeepLink = "?CampaignId=439892&ContentId=1925865&MerchantId=105064&Page=Product"
	expect.Link = "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892&merchantId=105064"

	if actual != expect {
		t.Errorf("actual was incorrect, got: %+v, want: %+v.", actual, expect)
	}
}

func TestDeepLinkToUrlService(t *testing.T) {
	model := models.UriRequest{Link: "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064"}

	actual, _ := linkService.DeepLinkToUrl(&model)
	var expect models.UriResponse
	expect.DeepLink = "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064"
	expect.Link = "/brand/name-p-1925865?boutiqueId=439892&merchantId=105064"

	if actual != expect {
		t.Errorf("actual was incorrect, got: %+v, want: %+v.", actual, expect)
	}
}
