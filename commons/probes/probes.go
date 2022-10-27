package probes

type Config struct {
	ReadinessPath        string `mapstructure:"readiness_path"`
	LivenessPath         string `mapstructure:"liveness_path"`
	Port                 string `mapstructure:"port"`
	Pprof                string `mapstructure:"pprof"`
	PrometheusPath       string `mapstructure:"prometheus_path"`
	PrometheusPort       string `mapstructure:"prometheus_port"`
	CheckIntervalSeconds int    `mapstructure:"check_intervalS=_seconds"`
}
