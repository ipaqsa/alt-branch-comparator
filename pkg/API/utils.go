package API

import (
	"altComparator/pkg/VersionsComparator"
	"encoding/json"
	"os"
)

func InSet(set map[string]Package, packName string) bool {
	if _, ok := set[packName]; ok {
		return true
	}
	return false
}

func (difference *Difference) Add(n int, pack Package) {
	if n == 1 {
		difference.FirstUniqueArray = append(difference.FirstUniqueArray, pack)
	} else if n == 2 {
		difference.SecondUniqueArray = append(difference.SecondUniqueArray, pack)
	} else if n == 3 {
		difference.VersionDifference = append(difference.VersionDifference, pack)
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

func Compare(firstname, secondname string, firstmap, secondmap map[string]Package, path string) error {
	var difference = &Difference{}
	for name, pack := range firstmap {
		if !InSet(secondmap, name) {
			difference.Add(1, pack)
		} else {
			if VersionsComparator.SecondVersionLessFirst(pack.Version, secondmap[name].Version) {
				difference.Add(3, pack)
			}

			delete(secondmap, name)
		}
		delete(firstmap, name)
	}
	for _, pack := range secondmap {
		difference.Add(2, pack)
	}
	println("Unique elements in %s:", firstname, len(difference.FirstUniqueArray))
	println("Unique elements in %s:", secondname, len(difference.SecondUniqueArray))
	println("Elements with a difference version: ", len(difference.VersionDifference))
	return writeJSON(path, difference)
}
