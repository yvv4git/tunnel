package speedtest

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yvv4git/tunnel/internal/infrastructure/config"
)

var (
	bytesReceived = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tcp_speedtest_server_bytes_received_total",
		Help: "Total number of bytes received from clients.",
	})

	bytesSent = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tcp_speedtest_server_bytes_sent_total",
		Help: "Total number of bytes sent to clients.",
	})
)

func StartMetricsWebServer(cfg config.MetricsWebServer) {
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), nil); err != nil {
			panic(err)
		}
	}()
}
