package VersionsComparator

import "regexp"

var alphs = regexp.MustCompile("([a-zA-Z]+)|([0-9]+)|(~)")

type Version struct {
	epoch   int
	version string
	release string
}
