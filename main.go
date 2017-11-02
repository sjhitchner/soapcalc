package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sjhitchner/soapcalc/domain"
)

var (
	dataPath string
)

func init() {
	flag.StringVar(&dataPath, "d", "data", "Path to soap data")
}

func main() {
	flag.Parse()

	recipeFilename := flag.Arg(0)

	recipeIn, err := domain.LoadRecipeInput(recipeFilename)
	CheckError(err)

	fmt.Println(recipeIn)

	lipids, err := LoadLipids(dataPath + "/sap.json")
	CheckError(err)

	recipe, err := Calculate(lipids, recipeIn)
	CheckError(err)

	fmt.Println(recipe)
}

func Calculate(lipids map[string]domain.Lipid, recipeIn *domain.RecipeInput) (*domain.Recipe, error) {

	recipe := domain.Recipe{
		Units:              recipeIn.Units,
		LyeType:            recipeIn.LyeType,
		WaterToLipidRatio:  recipeIn.WaterToLipidRatio,
		SuperFatPercentage: recipeIn.SuperFatPercentage,
	}

	for _, lipidInput := range recipeIn.Lipids {
		lipid, ok := lipids[lipidInput.Lipid]
		if !ok {
			return nil, fmt.Errorf("Invalid Value Lipid '%s'", lipidInput.Lipid)
		}

		weight := recipeIn.LipidWeight * lipidInput.Percentage

		recipe.Lipids = append(recipe.Lipids, domain.LipidComponent{
			Name:       lipidInput.Lipid,
			Percentage: lipidInput.Percentage,
			Weight:     weight,
		})

		recipe.LyeWeight += lipid.NaOH * weight * (1 - recipeIn.SuperFatPercentage)
		//recipe.Iodine += lipidInput.Percentage * lipid.Iodine
		//recipe.INS += lipidInput.Percentage * lipid.INS
		recipe.Lauric += lipidInput.Percentage * lipid.Lauric
		recipe.Myristic += lipidInput.Percentage * lipid.Myristic
		recipe.Palmitic += lipidInput.Percentage * lipid.Palmitic
		recipe.Stearic += lipidInput.Percentage * lipid.Stearic
		recipe.Ricinoleic += lipidInput.Percentage * lipid.Ricinoleic
		recipe.Oleic += lipidInput.Percentage * lipid.Oleic
		recipe.Linoleic += lipidInput.Percentage * lipid.Linoleic
		recipe.Linolenic += lipidInput.Percentage * lipid.Linolenic
		//recipe.Hardness += lipidInput.Percentage * lipid.Hardness
		//recipe.Cleansing += lipidInput.Percentage * lipid.Cleansing
		//recipe.Condition += lipidInput.Percentage * lipid.Condition
		//recipe.Bubbly += lipidInput.Percentage * lipid.Bubbly
		//recipe.Creamy += lipidInput.Percentage * lipid.Creamy
	}

	recipe.WaterWeight = recipeIn.WaterToLipidRatio * recipeIn.LipidWeight

	/*
		for _, lipid := range lipids {
			fmt.Println(lipid)
		}
	*/

	return &recipe, nil
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
