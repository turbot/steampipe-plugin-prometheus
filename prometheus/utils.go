package prometheus

import (
	"context"
	"os"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"

	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (v1.API, error) {
	api, err := getConnectionMemoize(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return api.(v1.API), nil
}

var getConnectionMemoize = plugin.HydrateFunc(getConnectionUncached).Memoize(memoize.WithCacheKeyFunction(getConnectionCackeKey))

func getConnectionUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
return connectRaw(ctx, d.Connection)
}

func getConnectionCackeKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "connectPrometheus"
	return key, nil
}

func connectRaw(ctx context.Context, c *plugin.Connection) (v1.API, error) {

	var address string
	// Prefer config settings
	prometheusConfig := GetConfig(c)

	address = os.Getenv("PROMETHEUS_URL")
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

	if err != nil {
		plugin.Logger(ctx).Error("connectRaw", "client connection error", err)
		return nil, err
	}

	conn := v1.NewAPI(client)

	if err != nil {
		plugin.Logger(ctx).Error("connectRaw", "cache-set", err)
		return conn, err
	}

	return conn, nil
}
