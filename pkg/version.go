package pkg

var version = "unset"

func SetVersion(v string) {
	version = v
}

func GetVersion() string {
	return version
}
