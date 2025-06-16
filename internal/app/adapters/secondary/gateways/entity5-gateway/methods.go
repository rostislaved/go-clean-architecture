package entity5_gateway

import (
	"context"
	"net/url"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity5"
	providerhelpers "github.com/rostislaved/go-clean-architecture/internal/pkg/provider-helpers"
)

func (prv *Entity5Gateway) Get(ctx context.Context, input struct{}) (entities []entity5.Entity5, err error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	values := url.Values{
		"param1": []string{"value1", "value2"},
		"param2": []string{"value1"},
	}

	req := providerhelpers.CreateRequest(ctx, prv.client, prv.config.Endpoints.Get)

	request := toRequest(input)

	var response ResponseGet

	req.
		SetQueryParamsFromValues(values).
		SetBody(request).
		ForceContentType("application/json").
		SetResult(&response)

	resp, err := req.Send()
	if err != nil {
		return
	}

	err = providerhelpers.ValidateStatusCode(resp.StatusCode(), resp.Body())
	if err != nil {
		return
	}

	entities = response.ToEntity()

	return entities, nil
}

type (
	RequestGet  struct{}
	ResponseGet struct{}
)

func (r ResponseGet) ToEntity() (entities []entity5.Entity5) {
	// mapping

	return entities
}

func toRequest(input struct{}) (request RequestGet) {
	// mapping

	return request
}
