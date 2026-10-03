package v1

type dtoTTML struct {
	ITunesTiming string `xml:"timing,attr"`
	Head         head   `xml:"head"`
	Body         body   `xml:"body"`
}
