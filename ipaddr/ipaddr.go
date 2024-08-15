package ipaddr

import (
	"errors"
	"github.com/tangthinker/aliyun-ddns/curl"
	"log"
	"regexp"
)

var IPv6NotSupportError = errors.New("IPv6 address not supported")

const IPv6Regexp = `/^([\da-fA-F]{1,4}:){6}((25[0-5]|2[0-4]\d|1?\d\d?)\.){3}(25[0-5]|2[0-4]\d|1?\d\d?)$|^([\da-fA-F]{0,4}:):([\da-fA-F]{1,4}:){0,4}((25[0-5]|2[0-4]\d|1?\d\d?)\.){3}(25[0-5]|2[0-4]\d|1?\d\d?)$|^([\da-fA-F]{1,4}:){2}:([\da-fA-F]{1,4}:){0,3}((25[0-5]|2[0-4]\d|1?\d\d?)\.){3}(25[0-5]|2[0-4]\d|1?\d\d?)$|^([\da-fA-F]{1,4}:){3}:([\da-fA-F]{1,4}:){0,2}((25[0-5]|2[0-4]\d|1?\d\d?)\.){3}(25[0-5]|2[0-4]\d|1?\d\d?)$|^([\da-fA-F]{1,4}:){4}:([\da-fA-F]{1,4}:){0,1}((25[0-5]|2[0-4]\d|1?\d\d?)\.){3}(25[0-5]|2[0-4]\d|1?\d\d?)$|^([\da-fA-F]{1,4}:){5}:((25[0-5]|2[0-4]\d|1?\d\d?)\.){3}(25[0-5]|2[0-4]\d|1?\d\d?)$|^([\da-fA-F]{1,4}:){7}[\da-fA-F]{1,4}$|^[\da-fA-F]{0,4}:((:[\da-fA-F]{1,4}){1,6}|:)$|^([\da-fA-F]{1,4}:){2}((:[\da-fA-F]{1,4}){1,5}|:)$|^([\da-fA-F]{1,4}:){3}((:[\da-fA-F]{1,4}){1,4}|:)$|^([\da-fA-F]{1,4}:){4}((:[\da-fA-F]{1,4}){1,3}|:)$|^([\da-fA-F]{1,4}:){5}((:[\da-fA-F]{1,4}){1,2}|:)$|^([\da-fA-F]{1,4}:){6}:([\da-fA-F]{1,4})?$|^([\da-fA-F]{1,4}:){6}:$/`

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

	match, err := regexp.Match(IPv6Regexp, resp.Body())

	if err != nil {
		return "", err
	}

	if !match {
		log.Println("IPv6 address not supported, return ip addr is ", string(resp.Body()))
		return "", IPv6NotSupportError
	}

	return string(resp.Body()), nil

}
