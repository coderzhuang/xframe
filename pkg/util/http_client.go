package util

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io/ioutil"
	"net/http"
)

func Get(c context.Context) {
	url := "http://127.0.0.1:8080/version"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	otel.GetTextMapPropagator().Inject(c, propagation.HeaderCarrier(req.Header))

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()

	_, _ = ioutil.ReadAll(resp.Body)
}
