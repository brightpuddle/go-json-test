package main

import (
	_ "embed"

	// "encoding/json"

	"github.com/goccy/go-json"
)

//go:embed faults.json
var faultData []byte

// FaultRecommendation is fault recommendation data from CCO.
type FaultRecommendation struct {
	Explanation    string   `json:"explanation"`
	Recommendation string   `json:"recommendation"`
	Steps          []string `json:"steps"`
}

// Faults decodes fault data from CCO
func Faults() (map[string]*FaultRecommendation, error) {

	res := make(map[string]*FaultRecommendation)
	if err := json.Unmarshal(faultData, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func main() {
	Faults()
}
