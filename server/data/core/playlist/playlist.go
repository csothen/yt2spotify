package playlist

type List struct {
	Id      string
	Url     string
	Sources map[string]*Source
}

type Source struct {
	Url string
}
