package v1

import (
	"encoding/xml"
	"errors"
	"strings"
)

func ParseString(data string) (TTML, error) {
	var err error
	dto := dtoTTML{}

	if err := xml.Unmarshal([]byte(data), &dto); err != nil {
		return TTML{}, err
	}

	ttml := TTML{
		TimingType: TimingType(dto.ITunesTiming),
		Metadata:   parseMetadata(&dto.Head.HeadMetadata),
	}
	ttml.Body, err = parseBody(&dto.Body)
	if err != nil {
		return TTML{}, err
	}

	return ttml, nil
}

func parseMetadata(in *headMetadata) (out TTMLMetadata) {
	// TTM agents section
	out.Agents = make(map[string][]string)
	if len(in.TtmAgents) == 0 {
		out.Agents["v1"] = []string{"person", "", ""}
	} else {
		for _, ag := range in.TtmAgents {
			out.Agents[ag.XmlId] = []string{ag.Type, ag.Name.Type, ag.Name.Value}
		}
	}

	// ITunes metadata section
	out.Songwriters = make([]string, len(in.ITunesMetadata.Songwriters))
	copy(out.Songwriters, in.ITunesMetadata.Songwriters)

	// AMLL metadata section
	out.AdditionalMetadata = make(map[string]string)
	for _, meta := range in.AMLLMetadata {
		switch meta.Key {
		case "musicName":
			out.TrackName = append(out.TrackName, meta.Value)
		case "artists":
			out.Artists = append(out.Artists, meta.Value)
		case "album":
			out.Album = append(out.Album, meta.Value)
		case "isrc":
			out.ISRC = meta.Value
		case "spotifyId":
			out.SpotifyID = meta.Value
		case "appleMusicId":
			out.AppleMusicID = meta.Value
		case "ncmMusicId":
			out.NCMMusicID = meta.Value
		case "qqMusicId":
			out.QQMusicID = meta.Value
		case "ttmlAuthorGithub":
			out.LyricAuthorGitHubID = meta.Value
		case "ttmlAuthorGithubLogin":
			out.LyricAuthorGitHubUsername = meta.Value
		default:
			out.AdditionalMetadata[meta.Key] = meta.Value
		}
	}

	// Additional TTM metadata section
	if in.TtmTitle != "" {
		out.TrackName = []string{in.TtmTitle}
	}
	if in.TtmDesc != "" {
		out.AdditionalMetadata["desc"] = in.TtmDesc
	}
	if in.TtmCopyright != "" {
		out.AdditionalMetadata["copyright"] = in.TtmCopyright
	}
	if in.TtmActor != "" {
		out.AdditionalMetadata["actor"] = in.TtmActor
	}

	return
}

func parseBody(in *body) (out TTMLBody, err error) {
	out.Duration = parseTiming(in.Dur)
	if out.Duration == -1 {
		return TTMLBody{}, errors.New("failed to parse body duration")
	}
	for _, div := range in.BodyDiv {
		part := TTMLSongPart{
			Begin:    parseTiming(div.BeginTime),
			End:      parseTiming(div.EndTime),
			SongPart: div.ITunesSongPart,
		}
		if part.Begin == -1 {
			return TTMLBody{}, errors.New("failed to parse begin timing in div")
		}
		if part.End == -1 {
			return TTMLBody{}, errors.New("failed to parse end timing in div")
		}
		if part.Begin > part.End {
			return TTMLBody{}, errors.New("begin timing is later than end timing in div")
		}

		for _, p := range div.P {
			lines, err := parseLineParts(&p)
			if err != nil {
				return TTMLBody{}, err
			}
			part.Lines = append(part.Lines, lines...)
		}

		out.Parts = append(out.Parts, part)
	}

	return
}

func parseLineParts(p *bodyP) (lines []TTMLLine, err error) {
	pBegin := parseTiming(p.BeginTime)
	pEnd := parseTiming(p.EndTime)
	bgLines := make([]TTMLLine, 0)
	line := TTMLLine{
		Begin:        pBegin,
		End:          pEnd,
		Agent:        p.TtmAgent,
		Key:          p.ITunesKey,
		IsBackground: false,
	}
	if line.Begin == -1 {
		return nil, errors.New("failed to parse begin timing in p")
	}
	if line.End == -1 {
		return nil, errors.New("failed to parse end timing in p")
	}
	if line.Begin > line.End {
		return nil, errors.New("begin timing is later than end timing in p")
	}

	spansSplit := strings.Split(p.InnerXML, "<span")

	for i, span := range p.Spans {
		switch span.Role {
		case spanRoleTypeTl:
			if line.Translations == nil {
				line.Translations = make(map[string]string)
			}
			line.Translations[span.Lang] = span.Value
		case spanRoleTypeRo:
			if line.Romanizations == nil {
				line.Romanizations = make(map[string]string)
			}
			line.Romanizations[span.Lang] = span.Value
		case spanRoleTypeBg:
			bgP := bodyP{
				BeginTime: span.BeginTime,
				EndTime:   span.EndTime,
				TtmAgent:  p.TtmAgent,
				ITunesKey: p.ITunesKey,
				Spans:     span.Spans,
				Value:     span.Value,
				InnerXML:  span.InnerXML,
			}
			bgSpanLines, err := parseLineParts(&bgP)
			if err != nil {
				return nil, errors.New("failed to parse bg lines")
			}
			for j := range bgSpanLines {
				bgSpanLines[j].IsBackground = true
			}
			bgLines = append(bgLines, bgSpanLines...)
		default:
			part := TTMLLinePart{
				Begin:         parseTiming(span.BeginTime),
				End:           parseTiming(span.EndTime),
				Value:         span.Value,
				HasSpaceAfter: strings.HasSuffix(spansSplit[i+1], " "),
			}
			line.LineParts = append(line.LineParts, part)
		}
	}

	if strings.TrimSpace(p.Value) != "" {
		line.LineParts = []TTMLLinePart{
			{
				Begin: pBegin,
				End:   pEnd,
				Value: strings.TrimSpace(p.Value),
			},
		}
	}

	if len(line.LineParts) != 0 {
		line.LineParts[len(line.LineParts)-1].HasSpaceAfter = false
	}

	lines = append(lines, line)
	lines = append(lines, bgLines...)

	return
}
