package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/sjhitchner/soapcalc/domain"
)

var (
	dataPath  string
	newRecipe bool
	findLipid string
	showLipid string
)

func init() {
	flag.StringVar(&dataPath, "d", "data", "Path to soap data")
	flag.BoolVar(&newRecipe, "n", false, "Create a new recipe")
	flag.StringVar(&findLipid, "f", "", "Find Lipid")
	flag.StringVar(&showLipid, "l", "", "Show Lipid")
}

func main() {
	flag.Parse()

	lipids, err := LoadLipids(dataPath + "/sap.json")
	CheckError(err)

	if newRecipe {
		fmt.Println(domain.RecipeInput{
			Units:              domain.Ounces,
			LyeType:            domain.NaOH,
			LipidWeight:        30,
			WaterToLipidRatio:  .35,
			SuperFatPercentage: 0.05,
			FragranceRatio:     0.5,
			Lipids: []domain.LipidInput{
				domain.LipidInput{
					Lipid:      "Olive Oil",
					Percentage: 1,
				},
			},
		}.String())
		return
	}

	if findLipid != "" {
		findLipid = strings.ToLower(findLipid)
		for lipid, _ := range lipids {
			if strings.Contains(strings.ToLower(lipid), findLipid) {
				fmt.Println(lipid)
			}
		}
		return
	}

	if showLipid != "" {
		CheckError(ShowLipid(lipids, showLipid))
		return
	}

	recipeFilename := flag.Arg(0)

	recipeIn, err := domain.LoadRecipeInput(recipeFilename)
	CheckError(err)

	fmt.Println(recipeIn)

	recipe, err := Calculate(lipids, recipeIn)
	CheckError(err)

	fmt.Println(recipe)
}

func Calculate(lipids map[string]domain.Lipid, recipeIn *domain.RecipeInput) (*domain.Recipe, error) {

	recipe := domain.Recipe{
		Name:               recipeIn.Name,
		Units:              recipeIn.Units,
		LyeType:            recipeIn.LyeType,
		WaterToLipidRatio:  recipeIn.WaterToLipidRatio,
		SuperFatPercentage: recipeIn.SuperFatPercentage,
		FragranceRatio:     recipeIn.FragranceRatio,
	}

	var totalPercentage float64

	for _, lipidInput := range recipeIn.Lipids {
		lipid, ok := lipids[lipidInput.Lipid]
		if !ok {
			return nil, fmt.Errorf("Invalid Value Lipid '%s'", lipidInput.Lipid)
		}

		totalPercentage += lipidInput.Percentage

		weight := recipeIn.LipidWeight * lipidInput.Percentage

		recipe.Lipids = append(recipe.Lipids, domain.LipidComponent{
			Name:       lipidInput.Lipid,
			Percentage: lipidInput.Percentage,
			Weight:     weight,
		})

		//recipe.LyeWeight += lipid.NaOH * weight * (1 - recipeIn.SuperFatPercentage)
		recipe.LyeWeight += lipid.NaOH * weight
		recipe.LipidWeight += weight
		recipe.Iodine += lipidInput.Percentage * float64(lipid.Iodine)
		recipe.INS += lipidInput.Percentage * float64(lipid.INS)

		recipe.Lauric += lipidInput.Percentage * lipid.Lauric * 100
		recipe.Myristic += lipidInput.Percentage * lipid.Myristic * 100
		recipe.Palmitic += lipidInput.Percentage * lipid.Palmitic * 100
		recipe.Stearic += lipidInput.Percentage * lipid.Stearic * 100
		recipe.Ricinoleic += lipidInput.Percentage * lipid.Ricinoleic * 100
		recipe.Oleic += lipidInput.Percentage * lipid.Oleic * 100
		recipe.Linoleic += lipidInput.Percentage * lipid.Linoleic * 100
		recipe.Linolenic += lipidInput.Percentage * lipid.Linolenic * 100

		recipe.Hardness += lipidInput.Percentage * float64(lipid.Hardness)
		recipe.Cleansing += lipidInput.Percentage * float64(lipid.Cleansing)
		recipe.Condition += lipidInput.Percentage * float64(lipid.Condition)
		recipe.Bubbly += lipidInput.Percentage * float64(lipid.Bubbly)
		recipe.Creamy += lipidInput.Percentage * float64(lipid.Creamy)
	}

	if math.Abs(totalPercentage-1) > 0.001 {
		return nil, fmt.Errorf("Lipid percentages (%.3f) do not add up to 1", totalPercentage)
	}

	recipe.LyeWeight *= (1 - recipeIn.SuperFatPercentage)
	recipe.FragranceWeight = recipeIn.FragranceRatio * recipe.LipidWeight / 16
	recipe.WaterWeight = recipeIn.WaterToLipidRatio * recipeIn.LipidWeight

	recipe.TotalWeight = recipe.WaterWeight +
		recipe.FragranceWeight +
		recipe.LyeWeight +
		recipe.LipidWeight

	return &recipe, nil
}

func ShowLipid(lipids map[string]domain.Lipid, lipidName string) error {

	lipid, ok := lipids[lipidName]
	if !ok {
		return fmt.Errorf("Invalid Value Lipid '%s'", lipidName)
	}

	fmt.Println(lipid)

	return nil
}

func LoadLipids(filename string) (map[string]domain.Lipid, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lipids, err := domain.ReadAllLipids(f)
	if err != nil {
		return nil, err
	}

	lipidMap := make(map[string]domain.Lipid)
	for _, lipid := range lipids {
		lipidMap[lipid.Name] = lipid
	}

	return lipidMap, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
