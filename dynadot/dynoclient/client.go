package dynoclient

import (
	"github.com/vitalick/dynadotTerraformProvider/dynadot/dns"
	"github.com/vitalick/dynadotTerraformProvider/dynadot/dynorecord"
)

type DynadotClient struct {
	Domain, Key string
	client      *dns.DNS
}

func NewDynadotClient(domain, key string) *DynadotClient {
	return &DynadotClient{
		Domain: domain,
		Key:    key,
		client: dns.NewDNS(),
	}
}

func (d *DynadotClient) CreateDNS(record *dynorecord.DynadotRecord) (int, error) {
	return d.client.CreateDNS(record)
}

func (d *DynadotClient) GetDNS(name, zoneID string) (*dynorecord.DynadotRecord, error) {
	return d.client.GetDNS(name, zoneID)
}

func (d *DynadotClient) DeleteDNS(name, zoneID string) error {
	return d.client.DeleteDNS(name, zoneID)
}

// Additional methods for other Dynadot API endpoints
