package prometheus

import (
	"context"
	"os"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"

	"github.com/turbot/steampipe-plugin-sdk/v5/connection"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (v1.API, error) {
	return connectRaw(ctx, d.ConnectionCache, d.Connection)
}

func connectRaw(ctx context.Context, cc *connection.ConnectionCache, c *plugin.Connection) (v1.API, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "prometheus"
	if cachedData, ok := cc.Get(ctx, cacheKey); ok {
		return cachedData.(v1.API), nil
	}

	// Prefer config settings
	prometheusConfig := GetConfig(c)

	address := os.Getenv("PROMETHEUS_ADDRESS")
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
	err = cc.Set(ctx, cacheKey, conn)

	if err != nil {
		plugin.Logger(ctx).Error("connectRaw", "cache-set", err)
	}

	return conn, nil
}
