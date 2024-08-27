package domain

import "encoding/json"

const USDTRUB = "usdtrub"

type GrntxDepth struct {
	Timestamp int64       `json:"timestamp"`
	Asks      []GrntxRate `json:"asks"`
	Bids      []GrntxRate `json:"bids"`
}

type GrntxRate struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}

func UnmarshalGrntxDepth(data []byte) (GrntxDepth, error) {
	var r GrntxDepth
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GrntxDepth) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
