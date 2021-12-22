connection "prometheus" {
  plugin = "prometheus"

  # HTTP address of your prometheus server
  # address = "http://localhost:9090"
  
  # List of metrics that will be considered for dynamic table creation
  # Refer to https://prometheus.io/docs/prometheus/latest/querying/basics/ for 
  # information about expressions that are supported
  # e.g. - "http_requests_.*" returns all metrics starting with http_requests_
  #  - ".*error.*" returns all metrics containing the word 'error'
  # Will return all metrics if not configured
  # metrics = ["http_requests_.*", ".*error.*"]
}
