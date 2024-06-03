package types

import (
	"encoding/json"
)

type Value struct {
	Memo string `json:"memo"`
}

type Gentxs []Gentx
type Gentx struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}

type Genutil struct {
	Gentxs Gentxs `json:"gentxs"`
}



