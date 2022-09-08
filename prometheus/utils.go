package prometheus

import (
	"context"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (v1.API, error) {
	return connectRaw(ctx, d)
}

func connectRaw(_ context.Context, d *plugin.QueryData) (v1.API, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "prometheus"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(v1.API), nil
	}

	var address string

	// Prefer config settings
	prometheusConfig := GetConfig(d.Connection)
	if prometheusConfig.Address != nil {
		address = *prometheusConfig.Address
	}

	// Error if the minimum config is not set
	if address == "" {
		// Panic since we cannot create a valid empty API to return
		panic("address must be configured")
	}

	client, err := api.NewClient(api.Config{
		Address: address,
	})

	conn := v1.NewAPI(client)

	if err != nil {
		return conn, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}
