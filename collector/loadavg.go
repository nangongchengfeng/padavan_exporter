package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/ssh"
	"strings"
)

type loadAverageCollector struct {
	metrics []*prometheus.Desc
	sc      *ssh.Client
}

func (l *loadAverageCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range l.metrics {
		ch <- metric
	}
}

func (l *loadAverageCollector) Collect(ch chan<- prometheus.Metric) {
	splict := strings.Split(mustGetContent(l.sc, "/proc/loadavg"), " ")
	for i, metric := range l.metrics {
		ch <- prometheus.MustNewConstMetric(metric, prometheus.GaugeValue, mustParseFloat(splict[i]))
	}
}

func NewLoadAverageCollector(sc *ssh.Client) *loadAverageCollector {
	return &loadAverageCollector{
		metrics: []*prometheus.Desc{
			prometheus.NewDesc("node_load1", "1m load average.", nil, nil),
			prometheus.NewDesc("node_load5", "5m load average.", nil, nil),
			prometheus.NewDesc("node_load15", "15m load average.", nil, nil),
		},
		sc: sc,
	}
}
