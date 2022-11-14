package API

type Package struct {
	Name      string `json:"name"`
	Epoch     int    `json:"epoch"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	Arch      string `json:"arch"`
	Disttag   string `json:"disttag"`
	Buildtime int    `json:"buildtime"`
	Source    string `json:"source"`
}

type Response struct {
	Request_args map[string]string `json:"request_args"`
	Length       int               `json:"length"`
	Packages     []Package         `json:"packages"`
}

type Difference struct {
	FirstUniqueArray  []Package `json:"first_unique_array"`
	SecondUniqueArray []Package `json:"second_unique_array"`
	VersionDifference []Package `json:"version_difference"`
}
