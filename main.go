package main

import (
	DDNSConfig "github.com/tangthinker/aliyun-ddns/config"
	"github.com/tangthinker/aliyun-ddns/ddns"
)

func main() {

	domainName := DDNSConfig.Config.GetString("domain_name")
	RR := DDNSConfig.Config.GetString("RR")
	recordType := DDNSConfig.Config.GetString("record_type")

	ddnsClient, err := ddns.NewDNSClient()

	if err != nil {
		panic(err)
	}

	ddnsClient.UpdateInterval(domainName, RR, recordType)
}
