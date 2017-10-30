package domain

import (
	"encoding/json"
	"io"
)

type Lipid struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	ScientificName string  `json:"scientific_name"`
	NaOH           float64 `json:"naoh"`
	KOH            float64 `json:"koh"`
	Iodine         int     `json:"iodine"`
	INS            int     `json:"ins"`
	Lauric         float64 `json:"lauric"`     // Sat 12:0
	Myristic       float64 `json:"myristic"`   // Sat 14:0
	Palmitic       float64 `json:"palmitic"`   // Sat 16:0
	Stearic        float64 `json:"stearic"`    // Sat 18:0
	Ricinoleic     float64 `json:"ricinoleic"` // MonoUnsat 18:1
	Oleic          float64 `json:"oleic"`      // MonoUnsat 18:1
	Linoleic       float64 `json:"linoleic"`   // PolyUnsat 18:2
	Linolenic      float64 `json:"linolenic"`  // PolyUnsat 18:3
	Hardness       int     `json:"hardness"`
	Cleansing      int     `json:"cleansing"`
	Condition      int     `json:"condition"`
	Bubbly         int     `json:"bubbly"`
	Creamy         int     `json:"creamy"`
}

func (t Lipid) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func ReadAllLipids(reader io.Reader) ([]Lipid, error) {

	lipids := make([]Lipid, 0, 200)

	dec := json.NewDecoder(reader)
	for {
		var lipid Lipid
		if err := dec.Decode(&lipid); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		lipids = append(lipids, lipid)
	}

	return lipids, nil
}
