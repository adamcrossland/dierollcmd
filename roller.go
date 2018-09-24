// roller
package main

import (
	"encoding/json"
)

type RollSpec struct {
	Sides    int64
	BestOf   int64
	DieCount int64
	Modifier int64
	Times    int64
}

type SetResult struct {
	Total int   `json:"total"`
	Count int   `json:"count"`
	Dies  []int `json:"dice"`
}

type RollResults struct {
	Count int         `json:"count"`
	Rolls []SetResult `json:"rolls"`
}

func (results RollResults) ToJSON() []byte {
	coded, _ := json.Marshal(results)
	return coded
}
