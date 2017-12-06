package zapier

import (
	"context"
	"net/url"
	"strconv"
)

type (
	TemplatesService service

	Template struct {
		CreateURL        string `json:"create_url"`
		Description      string `json:"description"`
		DescriptionPlain string `json:"description_plain"`
		DescriptionRaw   string `json:"description_raw"`
		Id               int    `json:"id"`
		Slug             string `json:"slug"`
		Status           string `json:"status"`
		Steps            []Step `json:"steps"`
		Title            string `json:"title"`
		Type             string `json:"type"`
		URL              string `json:"url"`
	}

	Step struct {
		API         string `json:"api"`
		Description string `json:"description"`
		HexColor    string `json:"hex_color"`
		Id          int    `json:"id"`
		Image       string `json:"image"`
		Images      Images `json:"images"`
		Slug        string `json:"slug"`
		Title       string `json:"title"`
		URL         string `json:"url"`
	}

	Images struct {
		URL128x128 string `json:"url_128x128"`
		URL16x16   string `json:"url_16x16"`
		URL32x32   string `json:"url_32x32"`
		URL64x64   string `json:"url_64x64"`
	}
)

// Get all Zap Templates for my app
//
// Zapier Partner API docs: https://zapier.com/developer/documentation/v2/partner-api/#partner-api-resources
func (svc *TemplatesService) List(ctx context.Context, limit int) ([]*Template, *Response, error) {
	if limit == 0 {
		limit = 25
	}
	params := &url.Values{}
	params.Set("limit", strconv.Itoa(limit))
	u := "zap-templates"
	req, err := svc.client.NewRequest("GET", u, params, nil)
	if err != nil {
		return nil, nil, err
	}

	var templates []*Template
	resp, err := svc.client.Do(ctx, req, &templates)
	return templates, resp, nil
}
