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

	lipids, err := LoadLipids(dataPath + "/sap.json")
	CheckError(err)

	for _, lipid := range lipids {
		fmt.Println(lipid)
	}
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
