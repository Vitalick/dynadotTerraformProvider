package dynorecord

type DynadotGeoLocation struct {
	Continent   string
	Country     string
	Subdivision string
}

type DynadotRecord struct {
	Name             string
	Type             string
	ZoneID           string
	Records          []string
	TTL              int
	SetIdent         string
	Weight           int
	Failover         string
	HealthID         string
	MultiValueAnswer bool
	GeoLocation      *DynadotGeoLocation
	Alias            *DynadotRecord
	EvaluateTarget   bool
}
