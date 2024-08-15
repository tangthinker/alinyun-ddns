package ddns

import (
	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	DDNSConfig "github.com/tangthinker/aliyun-ddns/config"
	"github.com/tangthinker/aliyun-ddns/ipaddr"
	"log"
	"time"
)

type Operator interface {
	UpdateDNS(recordId string, RR string, recordType string) error
	GetRecord(domainName string, RR string, recordType string) (*Record, error)
	UpdateInterval(domainName string, RR string, recordType string)
}

type Record struct {
	RecordId   string
	RR         string
	Value      string
	RecordType string
}

type Client struct {
	*dns.Client
	cachedHostIP string
}

func NewDNSClient() (*Client, error) {
	config := &openapi.Config{}

	accessKeyId := DDNSConfig.Config.GetString("access_key_id")
	accessKeySecret := DDNSConfig.Config.GetString("access_key_secret")
	regionId := DDNSConfig.Config.GetString("region_id")

	config.AccessKeyId = &accessKeyId
	config.AccessKeySecret = &accessKeySecret
	config.RegionId = &regionId

	result := &dns.Client{}

	result, err := dns.NewClient(config)

	if err != nil {
		return nil, err
	}

	return &Client{Client: result}, nil
}

func (c *Client) UpdateInterval(domainName string, RR string, recordType string) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover error: ", err)
			return
		}
	}()

	interval := DDNSConfig.Config.GetDuration("ddns_interval")

	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			var currentHostIP string
			var err error

			// ipv4
			if recordType == "A" {
				currentHostIP, err = ipaddr.GetIPv4Addr()
			}

			// ipv6
			if recordType == "AAAA" {
				currentHostIP, err = ipaddr.GetIPv6Addr()
			}

			if err != nil {
				log.Println("get ip address error: ", err)
				continue
			}

			if currentHostIP == "" {
				log.Println("current ip address is empty")
				continue
			}

			if currentHostIP != c.cachedHostIP {
				c.cachedHostIP = currentHostIP
				record, err := c.GetRecord(domainName, RR, recordType)
				if err != nil {
					log.Println("get record error: ", err)
					continue
				}
				err = c.UpdateDNS(record.RecordId, RR, recordType)
				if err != nil {
					log.Println("update dns error: ", err)
					continue
				}
				log.Println("update dns success")
				log.Println("current ip: ", currentHostIP)
			}
		}
	}

}

func (c *Client) UpdateDNS(recordId string, RR string, recordType string) error {
	currentHostIP, err := ipaddr.GetIPv6Addr()

	if err != nil {
		return err
	}

	req := &dns.UpdateDomainRecordRequest{}
	req.RR = &RR
	req.RecordId = &recordId
	req.Value = &currentHostIP
	req.Type = &recordType

	_, err = c.UpdateDomainRecord(req)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetRecord(domainName string, RR string, recordType string) (*Record, error) {
	req := &dns.DescribeDomainRecordsRequest{}
	req.DomainName = &domainName
	req.RRKeyWord = &RR
	req.Type = &recordType

	resp, err := c.DescribeDomainRecords(req)

	if err != nil {
		return nil, err
	}

	if len(resp.Body.DomainRecords.Record) == 0 {
		return nil, nil
	}

	record := resp.Body.DomainRecords.Record[0]

	return &Record{
		RecordId:   *record.RecordId,
		RR:         *record.RR,
		Value:      *record.Value,
		RecordType: *record.Type,
	}, nil
}
