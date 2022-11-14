package API

import (
	"encoding/json"
	"os"
)

func InSet(set map[Package]bool, pack Package) bool {
	if _, ok := set[pack]; ok {
		return true
	}
	return false
}

func (difference *Difference) Add(n int, pack Package) {
	if n == 1 {
		difference.FirstUniqueArray = append(difference.FirstUniqueArray, pack)
	} else if n == 2 {
		difference.SecondUniqueArray = append(difference.SecondUniqueArray, pack)
	}
}

func writeJSON(path string, difference interface{}) error {
	f, _ := os.Create(path)
	defer f.Close()
	diffJSON, err := json.MarshalIndent(difference, "", "\t")
	if err != nil {
		return err
	}
	f.Write(diffJSON)
	return nil
}

func Compare(first, second map[Package]bool, path string) error {
	var difference = &Difference{}
	for pack, _ := range first {
		if !InSet(second, pack) {
			difference.Add(1, pack)
		} else {
			delete(second, pack)
		}
		delete(first, pack)
	}
	for pack, _ := range second {
		difference.Add(2, pack)
	}
	return writeJSON(path, difference)
}
