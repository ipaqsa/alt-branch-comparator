package API

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
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

func Compare(first, second map[string]Package, path string) error {
	var difference = &Difference{}
	for name, pack := range first {
		if !InSet(second, name) {
			difference.Add(1, pack)
		} else {
			result := compareVersions(pack.Version, second[name].Version)
			if result == -1 {
				return errors.New("error compare versions")
			}
			if result == 1 {
				difference.Add(3, pack)
			}

			delete(second, name)
		}
		delete(first, name)
	}
	for _, pack := range second {
		difference.Add(2, pack)
	}
	println("Unique elements in 1:", len(difference.FirstUniqueArray))
	println("Unique elements in 2:", len(difference.SecondUniqueArray))
	println("Elements with a difference version: ", len(difference.VersionDifference))
	return writeJSON(path, difference)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func compareVersions(v1, v2 string) int {
	splitedV1 := strings.Split(v1, ".")
	splitedV2 := strings.Split(v2, ".")
	l := min(len(splitedV1), len(splitedV2))
	for i := 0; i < l; i++ {
		n1, err := strconv.Atoi(splitedV1[i])
		if err != nil {
			println(v1)
			errorLoggerAPI.Printf("compareVersions: %s", err.Error())
			return -1
		}
		n2, err := strconv.Atoi(splitedV2[i])
		if err != nil {
			println(v2)
			errorLoggerAPI.Printf("compareVersions: %s", err.Error())
			return -1
		}
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return 0
		}
	}
	if len(splitedV1) > len(splitedV2) {
		return 1
	}
	return 0
}
