package server

type ApiResp struct {
	LocA     string  `json:"locA,omitempty"`
	LocB     string  `json:"locB,omitempty"`
	LatA     float64 `json:"latA,omitempty"`
	LonA     float64 `json:"lonA,omitempty"`
	LatB     float64 `json:"latB,omitempty"`
	LonB     float64 `json:"lonB,omitempty"`
	Distance float64 `json:"distance,omitempty"`
	Azimuth  float64 `json:"azimuth,omitempty"`
}
