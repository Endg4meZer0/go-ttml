package v1

type TTML struct {
	TimingType
	Metadata TTMLMetadata
	Body     TTMLBody
}

type TimingType string

const (
	TimingTypeLine TimingType = "Line"
	TimingTypeWord TimingType = "Word"
)
