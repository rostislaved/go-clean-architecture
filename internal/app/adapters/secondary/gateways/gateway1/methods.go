package gateway1

import (
	"context"
	"net/url"
	"time"

	providerhelpers "github.com/rostislaved/go-clean-architecture/internal/pkg/provider-helpers"
)

func (prv *Gateway1) Method1(ctx context.Context, input struct{}) (output struct{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	values := url.Values{
		"param1": []string{"value1", "value2"},
		"param2": []string{"value1"},
	}

	req := providerhelpers.CreateRequest(ctx, prv.client, prv.config.Endpoints.SignFile)

	req.
		SetQueryParamsFromValues(values).
		SetBody(input).
		ForceContentType("application/json").
		SetResult(&output)

	resp, err := req.Send()
	if err != nil {
		return
	}

	err = providerhelpers.ValidateStatusCode(resp.StatusCode(), resp.Body())
	if err != nil {
		return
	}

	return
}
