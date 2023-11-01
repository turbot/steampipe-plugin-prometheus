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
	Header []requestHeaderPair
}

type requestHeaderPair struct {
	headerName  string
	headerValue string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	for _, h := range t.Header {
		if h.headerName != "" {
			req.Header.Add(h.headerName, h.headerValue)
		}
	}
	return http.DefaultTransport.RoundTrip(req)
}

func connectRaw(ctx context.Context, cc *connection.ConnectionCache, c *plugin.Connection) (v1.API, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "prometheus"
	if cachedData, ok := cc.Get(ctx, cacheKey); ok {
		return cachedData.(v1.API), nil
	}

	var address string
	var requestHeader map[string]string

	// Prefer config settings
	prometheusConfig := GetConfig(c)

	address = os.Getenv("PROMETHEUS_URL")
	if prometheusConfig.Address != nil {
		address = *prometheusConfig.Address
	}

	if prometheusConfig.RequestHeader != nil {
		requestHeader = prometheusConfig.RequestHeader
	}

	// Error if the minimum config is not set
	if address == "" {
		// Panic since we cannot create a valid empty API to return
		panic("address must be configured")
	}

	// Request header
	var headerNameAndValue []requestHeaderPair
	for k, v := range requestHeader {
		if requestHeader[k] != "" {
			headerNameAndValue = append(headerNameAndValue, requestHeaderPair{
				headerName:  k,
				headerValue: v,
			})
		}
	}

	client, err := api.NewClient(api.Config{
		Address: address,
		Client: &http.Client{
			Transport: &transport{
				Header: headerNameAndValue,
			},
		},
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
