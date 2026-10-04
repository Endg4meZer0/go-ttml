package ttml

type TTMLMetadata struct {
	// Agents is basically a list of "performers" on the track,
	// mapped to their respective IDs that are used in the lyrics body.
	// The slice of strings represents a set of 3 values:
	// - type (e.g. "person", "group", "other", etc.)
	// - name type (e.g. "full", I think it's mostly unused but I can be wrong)
	// - value (e.g. "Jacob Charlton", "C")
	Agents      map[string][]string
	TrackName   []string
	Artists     []string
	Album       []string
	ISRC        string
	Songwriters []string
	// Spotify ID
	SpotifyID string
	// Apple Music ID
	AppleMusicID string
	// QQMusic ID
	QQMusicID string
	// NetEase Cloud Music ID
	NCMMusicID                string
	LyricAuthorGitHubID       string
	LyricAuthorGitHubUsername string
	// AdditionalMetadata holds any other found metadata fields (stuff like "<ttm:copyright>" and etc.)
	AdditionalMetadata map[string]string
}
