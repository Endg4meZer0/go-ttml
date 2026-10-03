package ttml

type ttmName struct {
	Type  string `xml:"type,attr"`
	Value string
}

type ttmAgent struct {
	Type  string  `xml:"type,attr"`
	XmlId string  `xml:"id,attr"`
	Name  ttmName `xml:"name"`
}

type iTunesMetadata struct {
	Songwriters []string `xml:"songwriters>songwriter"`
}

type amllMetadata struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type headMetadata struct {
	TtmAgents      []ttmAgent     `xml:"agent"`
	TtmTitle       string         `xml:"title"`
	TtmDesc        string         `xml:"desc"`
	TtmCopyright   string         `xml:"copyright"`
	TtmActor       string         `xml:"actor"`
	ITunesMetadata iTunesMetadata `xml:"iTunesMetadata"`
	AMLLMetadata   []amllMetadata `xml:"meta"`
}

type head struct {
	HeadMetadata headMetadata `xml:"metadata"`
}
