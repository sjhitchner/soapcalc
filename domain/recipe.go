package domain

import (
	"encoding/json"
	"os"
)

const (
	Ounces Unit = "Ounces"
	Grams  Unit = "grams"

	NaOH LyeType = "NaOH"
	KOH  LyeType = "KOH"
)

type Unit string
type LyeType string

type Recipe struct {
	Name               string           `json:"name"`
	Description        string           `json:"description"`
	Units              Unit             `json:"units"`
	LyeType            LyeType          `json:"lye_type"`
	LipidWeight        float64          `json:"lipid_weight"`
	WaterToLipidRatio  float64          `json:"water_to_lipid_ratio"`
	SuperFatPercentage float64          `json:"super_fat_percentage"`
	FraganceRatio      float64          `json:"fragrance_ratio"`
	Lipids             []LipidComponent `json:"lipids"`
	WaterWeight        float64          `json:"water_weight"`
	LyeWeight          float64          `json:"lye_weight"`
	FragranceiWeight   float64          `json:"fragrance_weight"`
	Iodine             float64          `json:"iodine"`
	INS                float64          `json:"ins"`
	Lauric             float64          `json:"lauric"`     // Sat 12:0
	Myristic           float64          `json:"myristic"`   // Sat 14:0
	Palmitic           float64          `json:"palmitic"`   // Sat 16:0
	Stearic            float64          `json:"stearic"`    // Sat 18:0
	Ricinoleic         float64          `json:"ricinoleic"` // MonoUnsat 18:1
	Oleic              float64          `json:"oleic"`      // MonoUnsat 18:1
	Linoleic           float64          `json:"linoleic"`   // PolyUnsat 18:2
	Linolenic          float64          `json:"linolenic"`  // PolyUnsat 18:3
	Hardness           float64          `json:"hardness"`
	Cleansing          float64          `json:"cleansing"`
	Condition          float64          `json:"condition"`
	Bubbly             float64          `json:"bubbly"`
	Creamy             float64          `json:"creamy"`
}

type LipidComponent struct {
	Name       string  `json:"name"`
	Percentage float64 `json:"percentage"`
	Weight     float64 `json:"weight"`
}

func (t Recipe) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type RecipeInput struct {
	Units              Unit         `json:"units"`
	LyeType            LyeType      `json:"lye_type"`
	LipidWeight        float64      `json:"lipid_weight"`
	WaterToLipidRatio  float64      `json:"water_to_lipid_ratio"`
	SuperFatPercentage float64      `json:"super_fat_percentage"`
	FraganceRatio      float64      `json:"fragrance_ratio"`
	Lipids             []LipidInput `json:"lipids"`
}

type LipidInput struct {
	Lipid      string  `json:"lipid"`
	Percentage float64 `json:"percentage"`
}

func (t RecipeInput) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func LoadRecipeInput(filename string) (*RecipeInput, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var recipeInput RecipeInput

	if err := json.NewDecoder(f).Decode(&recipeInput); err != nil {
		return nil, err
	}

	return &recipeInput, nil
}
