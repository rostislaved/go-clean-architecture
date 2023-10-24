package providerhelpers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-resty/resty/v2"
	"go.uber.org/multierr"
	"moul.io/http2curl"
)

var RetryCondition = func(r *resty.Response, err error) bool {
	// retry if return is true

	if err != nil {
		return true
	}

	switch r.StatusCode() {
	case
		// 400:
		http.StatusRequestTimeout,
		http.StatusConflict,
		http.StatusTooManyRequests,

		// 500:
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:

		return true
	}

	return false
}

type Endpoint struct {
	Method  string
	Path    string
	Headers map[string]string
}

func CreateRequest(ctx context.Context, client *resty.Client, endpoint Endpoint) *resty.Request {
	req := client.R()

	req.Method = endpoint.Method
	req.URL = endpoint.Path

	req.SetContext(ctx)

	return req
}

func ValidateEndpoints(endpoints interface{}) (err error) {
	refValue := reflect.ValueOf(endpoints)

	n := refValue.NumField()

	var combinedErr error

	for i := 0; i < n; i++ {
		fieldName := refValue.Type().Field(i).Name

		fieldInterface := refValue.FieldByName(fieldName).Interface()

		err := ValidateEndpoint(fieldInterface)
		if err != nil {
			errString := fmt.Sprintf("\nэндпоинт: [%v] заполнен некорректно: %s", refValue.Type().Field(i).Name, err.Error())
			err := errors.New(errString)

			combinedErr = multierr.Append(combinedErr, err)
		}
	}

	if combinedErr != nil {
		return combinedErr
	}

	return nil
}

func ValidateEndpoint(endpoint interface{}) error {
	//method, ok1 := reflect.TypeOf(endpoint).FieldByName("Method")
	//path, ok2 := reflect.TypeOf(endpoint).FieldByName("Path")
	//if !(ok1 && ok2) {
	//	return nil
	//}
	e, ok := endpoint.(Endpoint)
	if !ok {
		return nil
	}

	methodValid := allHttpMethods[e.Method]
	methodNotValid := !methodValid

	pathIsEmpty := e.Path == ""

	if methodNotValid || pathIsEmpty {
		errString := ""

		if methodNotValid {
			errString = errString + fmt.Sprintf("\nМетод: [%s] не распознан", e.Method)
		}

		if pathIsEmpty {
			errString = errString + fmt.Sprintf("\nПуть не может быть пустым")
		}

		err := errors.New(errString)

		return err
	}

	return nil
}

var allHttpMethods = map[string]bool{
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodPost:    true,
	http.MethodPut:     true,
	http.MethodPatch:   true,
	http.MethodDelete:  true,
	http.MethodConnect: true,
	http.MethodOptions: true,
	http.MethodTrace:   true,
}

func PrintRequestHook(client *resty.Client, request *http.Request) error {
	curl, err := http2curl.GetCurlCommand(request)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("---http2curl------")
	fmt.Println(curl.String())
	fmt.Println("---http2curl------")
	fmt.Println()

	return nil
}

func And(hooks ...func(client *resty.Client, request *http.Request) error) func(client *resty.Client, request *http.Request) error {
	return func(client *resty.Client, request *http.Request) error {
		for _, hook := range hooks {
			err := hook(client, request)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
