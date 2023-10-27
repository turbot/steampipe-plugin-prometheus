package prometheus

import (
	"context"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"

	"github.com/turbot/steampipe-plugin-sdk/v5/connection"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (v1.API, error) {
	return connectRaw(ctx, d.ConnectionCache, d.Connection)
}

type transport struct {
	headerName  string
	headerValue string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.headerName != "" {
		req.Header.Add(t.headerName, t.headerValue)
	}
	return http.DefaultTransport.RoundTrip(req)
}

func connectRaw(ctx context.Context, cc *connection.ConnectionCache, c *plugin.Connection) (v1.API, error) {

	var address string

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "prometheus"
	if cachedData, ok := cc.Get(ctx, cacheKey); ok {
		return cachedData.(v1.API), nil
	}

	var address string
	var headerName string
	var headerValue string
  
	// Prefer config settings
	prometheusConfig := GetConfig(c)

	address = os.Getenv("PROMETHEUS_URL")
	if prometheusConfig.Address != nil {
		address = *prometheusConfig.Address
	}

	if prometheusConfig.HeaderName != nil {
		headerName = *prometheusConfig.HeaderName
	}

	if prometheusConfig.HeaderValue != nil {
		headerValue = *prometheusConfig.HeaderValue
	}

	// Error if the minimum config is not set
	if address == "" {
		// Panic since we cannot create a valid empty API to return
		panic("address must be configured")
	}

	if (headerName != "" && headerValue == "") || (headerName == "" && headerValue != "") {
		panic("must provide either both headerName and headerValue or neither")
	}

	client, err := api.NewClient(api.Config{
		Address: address,
		Client:  &http.Client{Transport: &transport{headerName: headerName, headerValue: headerValue}},
	})

	conn := v1.NewAPI(client)

	if err != nil {
		return conn, err
	}

	// Save to cache
	err = cc.Set(ctx, cacheKey, conn)

	if err != nil {
		plugin.Logger(ctx).Error("connectRaw", "cache-set", err)
	}

	return conn, nil
}
