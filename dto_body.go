package v1

type spanRoleType string

const (
	spanRoleTypeBg spanRoleType = "x-bg"
	spanRoleTypeTl spanRoleType = "x-translation"
	spanRoleTypeRo spanRoleType = "x-roman"
)

type bodySpan struct {
	BeginTime string       `xml:"begin,attr"`
	EndTime   string       `xml:"end,attr"`
	Role      spanRoleType `xml:"role,attr"`
	Lang      string       `xml:"lang,attr"`
	Spans     []bodySpan   `xml:"span"`
	Value     string       `xml:",chardata"`
	InnerXML  string       `xml:",innerxml"`
}

type bodyP struct {
	BeginTime string     `xml:"begin,attr"`
	EndTime   string     `xml:"end,attr"`
	TtmAgent  string     `xml:"agent,attr"`
	ITunesKey string     `xml:"key,attr"`
	Spans     []bodySpan `xml:"span"`
	Value     string     `xml:",chardata"`
	InnerXML  string     `xml:",innerxml"`
}

type bodyDiv struct {
	BeginTime      string  `xml:"begin,attr"`
	EndTime        string  `xml:"end,attr"`
	ITunesSongPart string  `xml:"songPart,attr"`
	P              []bodyP `xml:"p"`
}

type body struct {
	Dur     string    `xml:"dur,attr"`
	BodyDiv []bodyDiv `xml:"div"`
}
