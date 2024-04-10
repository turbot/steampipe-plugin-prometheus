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

// roundTripperWithBasicAuth is a custom RoundTripper that includes basic authentication headers
type roundTripperWithBasicAuth struct {
	http.RoundTripper
	username string
	password string
}

func (rt *roundTripperWithBasicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(rt.username, rt.password)
	return rt.RoundTripper.RoundTrip(req)
}

func connect(ctx context.Context, d *plugin.QueryData) (v1.API, error) {
	return connectRaw(ctx, d.ConnectionCache, d.Connection)
}

func connectRaw(ctx context.Context, cc *connection.ConnectionCache, c *plugin.Connection) (v1.API, error) {

	var address string
	username := ""
	password := ""

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "prometheus"
	if cachedData, ok := cc.Get(ctx, cacheKey); ok {
		return cachedData.(v1.API), nil
	}

	// Prefer config settings
	prometheusConfig := GetConfig(c)

	address = os.Getenv("PROMETHEUS_URL")
	if prometheusConfig.Address != nil {
		address = *prometheusConfig.Address
	}

	if prometheusConfig.Username != nil {
		username = *prometheusConfig.Username
	}
	if prometheusConfig.Password != nil {
		password = *prometheusConfig.Password
	}

	// Error if the minimum config is not set
	if address == "" {
		// Panic since we cannot create a valid empty API to return
		panic("address must be configured")
	}

	//handle authentication
	rt := api.DefaultRoundTripper
	rt = &roundTripperWithBasicAuth{
		RoundTripper: rt,
		username:     username,
		password:     password,
	}

	client, err := api.NewClient(api.Config{
		Address:      address,
		RoundTripper: rt,
	})

	if err != nil {
		plugin.Logger(ctx).Error("connectRaw", "client connection error", err)
		return nil, err
	}

	conn := v1.NewAPI(client)

	// Save to cache
	err = cc.Set(ctx, cacheKey, conn)

	if err != nil {
		plugin.Logger(ctx).Error("connectRaw", "cache-set", err)
		return conn, err
	}

	return conn, nil
}
