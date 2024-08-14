package ipaddr

import (
	"github.com/tangthinker/aliyun-ddns/curl"
)

func GetIPv4Addr() (string, error) {

	resp, err := curl.Client.R().
		EnableTrace().
		Get("https://ipinfo.io/ip")

	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil

}

func GetIPv6Addr() (string, error) {

	resp, err := curl.Client.R().
		EnableTrace().
		Get("https://ifconfig.me/ip")

	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil

}
