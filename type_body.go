package v1

type TTMLBody struct {
	Duration int
	Parts    []TTMLSongPart
}

type TTMLSongPart struct {
	Begin    int
	End      int
	SongPart string
	// Lines holds the ordered collection of lines in this song's text.
	Lines []TTMLLine
}

type TTMLLine struct {
	Begin int
	End   int
	// LineParts holds the ordered collection of words/syllables of this line.
	// Take note that it's always going to be length of 1 if the timing type is "Line".
	LineParts     []TTMLLinePart
	Agent         string
	Key           string
	IsBackground  bool
	Translations  map[string]string
	Romanizations map[string]string
}

type TTMLLinePart struct {
	Begin         int
	End           int
	Value         string
	HasSpaceAfter bool
}
