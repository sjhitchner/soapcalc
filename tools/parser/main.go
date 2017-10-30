package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Records [][]string

func (t Records) Len() int {
	return len(t)
}

func (t Records) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}

func (t Records) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

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

func main() {

	lipids := make(map[string]Lipid)

	ins, err := ParseRecords("data/ins.txt")
	CheckError(err)
	CheckError(AddValues(lipids, ins, "INS", 1))

	iodine, err := ParseRecords("data/iodine.txt")
	CheckError(err)
	CheckError(AddValues(lipids, iodine, "Iodine", 1))

	linoleic, err := ParseRecords("data/linoleic.txt")
	CheckError(err)
	CheckError(AddValues(lipids, linoleic, "Linoleic", 1))

	myristic, err := ParseRecords("data/myristic.txt")
	CheckError(err)
	CheckError(AddValues(lipids, myristic, "Myristic", 1))

	oleic, err := ParseRecords("data/oleic.txt")
	CheckError(err)
	CheckError(AddValues(lipids, oleic, "Oleic", 1))

	ricinoleic, err := ParseRecords("data/ricinoleic.txt")
	CheckError(err)
	CheckError(AddValues(lipids, ricinoleic, "Ricinoleic", 1))

	lauric, err := ParseRecords("data/lauric.txt")
	CheckError(err)
	CheckError(AddValues(lipids, lauric, "Lauric", 1))

	linolenic, err := ParseRecords("data/linolenic.txt")
	CheckError(err)
	CheckError(AddValues(lipids, linolenic, "Linolenic", 1))

	palmitic, err := ParseRecords("data/palmitic.txt")
	CheckError(err)
	CheckError(AddValues(lipids, palmitic, "PalmiticS", 1))

	stearic, err := ParseRecords("data/stearic.txt")
	CheckError(err)
	CheckError(AddValues(lipids, stearic, "Stearic", 1))

	qualities, err := ParseRecords("data/qualities.txt")
	CheckError(err)
	CheckError(AddValues(lipids, qualities, "Hardness", 1))
	CheckError(AddValues(lipids, qualities, "Cleansing", 2))
	CheckError(AddValues(lipids, qualities, "Condition", 3))
	CheckError(AddValues(lipids, qualities, "Bubbly", 4))
	CheckError(AddValues(lipids, qualities, "Creamy", 5))

	sap, err := ParseRecords("data/sap.txt")
	CheckError(err)
	CheckError(AddValues(lipids, sap, "NaOH", 1))
	CheckError(AddValues(lipids, sap, "KOH", 2))
	CheckError(AddValues(lipids, sap, "Description", 3))

	for _, lipid := range lipids {
		fmt.Println(lipid)
	}
}

func AddValues(lipids map[string]Lipid, records Records, fieldName string, index int) error {
	for _, record := range records {

		name := record[0]

		lipid, ok := lipids[name]
		if !ok {
			lipid = Lipid{
				Name: name,
			}
		}

		if err := AddValue(&lipid, record, fieldName, index); err != nil {
			return err
		}

		lipids[name] = lipid
	}

	return nil
}

func AddValue(lipid *Lipid, record []string, fieldName string, index int) error {
	ps := reflect.ValueOf(lipid)
	s := ps.Elem()
	if s.Kind() != reflect.Struct {
		return fmt.Errorf("Not at struct")
	}

	f := s.FieldByName(fieldName)
	if f.IsValid() {
		// A Value can be changed only if it is
		// addressable and was not obtained by
		// the use of unexported struct fields.
		if f.CanSet() {
			// change value of N
			switch f.Kind() {
			case reflect.Float64:
				if strings.Contains(record[index], "%") {
					str := strings.Replace(record[index], "%", "", 1)
					x, err := strconv.ParseFloat(str, 64)
					if err != nil {
						return fmt.Errorf("Can't convert %s to int", record[index])
					}
					x = x / 100
					f.SetFloat(x)
				} else {
					x, err := strconv.ParseFloat(record[index], 64)
					if err != nil {
						return fmt.Errorf("Can't convert %s to int", record[index])
					}
					f.SetFloat(x)
				}

			case reflect.String:
				f.SetString(record[index])

			case reflect.Int:
				i, err := strconv.ParseInt(record[index], 10, 64)
				if err != nil {
					return fmt.Errorf("Can't convert %s to int", record[index])
				}

				if !f.OverflowInt(i) {
					f.SetInt(i)
				}
			}
		}
	}

	return nil
}

func ParseRecords(filename string) (Records, error) {
	records, err := FileData(filename)
	if err != nil {
		return nil, err
	}
	sort.Sort(records)
	return records, nil
}

func FileData(filename string) (Records, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return Records(records), nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
