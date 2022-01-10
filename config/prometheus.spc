connection "prometheus" {
  plugin = "prometheus"

  # HTTP address of your prometheus server
  # address = "http://localhost:9090"

  # List of metrics that will be considered for dynamic table creation
  # Please refer to https://prometheus.io/docs/prometheus/latest/querying/basics
  # for information about supported expressions
  # For example:
  #   - ".+" matches all metrics
  #   - "prometheus_http_request.*" matches metrics starting with "prometheus_http_request"
  #   - ".*error.*" matches metrics containing the word "error"
  # metrics = [".+"]
}
