package config

type GeoIPConfig struct {
	GeoIP2File string `yaml:"maxmind-geoip2-file" json:"geoip2-file" default:"/var/lib/GeoIP/GeoIP2-Country.mmdb"`
}
