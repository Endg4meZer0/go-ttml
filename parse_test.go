package ttml

import "testing"

func TestParseString(t *testing.T) {
	example := `<tt
	xmlns:ttm="http://www.w3.org/ns/ttml#metadata"
	xmlns:tts="http://www.w3.org/ns/ttml#styling"
	xmlns:amll="http://www.example.com/ns/amll"
	xmlns:itunes="http://music.apple.com/lyric-ttml-internal" itunes:timing="Word">
	<head>
		<metadata>
			<ttm:agent type="person" xml:id="v1"/>
			<ttm:agent type="other" xml:id="v2"/>
			<iTunesMetadata>
				<songwriters>
					<songwriter>Violet Snowdrop</songwriter>
				</songwriters>
			</iTunesMetadata>
			<amll:meta key="musicName" value="blissful"/>
			<amll:meta key="artists" value="fallingwithscissors"/>
			<amll:meta key="album" value="the death and birth of an angel"/>
			<amll:meta key="spotifyId" value="5njTnzkgqKtAvMEuZFAFcb"/>
			<amll:meta key="ttmlAuthorGithubLogin" value="Endg4meZer0"/>
		</metadata>
	</head>
	<body dur="01:33.592">
		<div begin="00:00.326" end="01:33.592">
			<p begin="00:00.326" end="00:02.719" ttm:agent="v2" itunes:key="L1"><span begin="00:00.326" end="00:00.635">Eve</span><span begin="00:00.635" end="00:00.841">ry</span><span begin="00:00.841" end="00:01.136">thing's</span> <span begin="00:01.136" end="00:01.436">fine</span> <span begin="00:01.436" end="00:01.794">on</span> <span begin="00:01.794" end="00:02.236">this</span> <span begin="00:02.236" end="00:02.719">side</span></p>
			<p begin="00:11.734" end="00:15.631" ttm:agent="v1" itunes:key="L5">Could you blame me for wanting to leave?<span ttm:role="x-translation" xml:lang="ru">Можешь ли ты винить меня за то, что я хотел уйти?</span><span ttm:role="x-bg" begin="00:11.537" end="00:15.388"><span begin="00:11.537" end="00:12.037">(The</span> <span begin="00:12.037" end="00:13.037">mo</span><span begin="00:13.037" end="00:15.388">narch)</span></span></p>
		</div>
	</body>
</tt>`

	expectedTTML := TTML{
		TimingType: TimingTypeWord,
		Metadata: TTMLMetadata{
			Agents: map[string][]string{
				"v1": {"person", "", ""},
				"v2": {"other", "", ""},
			},
			Songwriters:               []string{"Violet Snowdrop"},
			TrackName:                 []string{"blissful"},
			Artists:                   []string{"fallingwithscissors"},
			Album:                     []string{"the death and birth of an angel"},
			SpotifyID:                 "5njTnzkgqKtAvMEuZFAFcb",
			LyricAuthorGitHubUsername: "Endg4meZer0",
		},
		Body: TTMLBody{
			Duration: 93592,
			Parts: []TTMLSongPart{
				{
					Begin:    326,
					End:      93592,
					SongPart: "",
					Lines: []TTMLLine{
						{
							Begin:         326,
							End:           2719,
							Agent:         "v2",
							IsBackground:  false,
							Translations:  nil,
							Romanizations: nil,
							LineParts: []TTMLLinePart{
								{
									Begin:         326,
									End:           635,
									Value:         "Eve",
									HasSpaceAfter: false,
								},
								{
									Begin:         635,
									End:           841,
									Value:         "ry",
									HasSpaceAfter: false,
								},
								{
									Begin:         841,
									End:           1136,
									Value:         "thing's",
									HasSpaceAfter: true,
								},
								{
									Begin:         1136,
									End:           1436,
									Value:         "fine",
									HasSpaceAfter: true,
								},
								{
									Begin:         1436,
									End:           1794,
									Value:         "on",
									HasSpaceAfter: true,
								},
								{
									Begin:         1794,
									End:           2236,
									Value:         "this",
									HasSpaceAfter: true,
								},
								{
									Begin:         2236,
									End:           2719,
									Value:         "side",
									HasSpaceAfter: false,
								},
							},
						},
						{
							Begin:        11734,
							End:          15631,
							Agent:        "v1",
							IsBackground: false,
							Translations: map[string]string{
								"ru": "Можешь ли ты винить меня за то, что я хотел уйти?",
							},
							Romanizations: nil,
							LineParts: []TTMLLinePart{
								{
									Begin:         11734,
									End:           15631,
									Value:         "Could you blame me for wanting to leave?",
									HasSpaceAfter: false,
								},
							},
						},
						{
							Begin:         11537,
							End:           15388,
							Agent:         "v1",
							IsBackground:  true,
							Translations:  nil,
							Romanizations: nil,
							LineParts: []TTMLLinePart{
								{
									Begin:         11537,
									End:           12037,
									HasSpaceAfter: true,
									Value:         "(The",
								},
								{
									Begin:         12037,
									End:           13037,
									HasSpaceAfter: false,
									Value:         "mo",
								},
								{
									Begin:         13037,
									End:           15388,
									HasSpaceAfter: false,
									Value:         "narch)",
								},
							},
						},
					},
				},
			},
		},
	}

	resTTML, err := ParseString(example)

	if err != nil {
		t.Fatal(err)
	}

	if resTTML.TimingType != expectedTTML.TimingType {
		t.Fatal("Didn't match TimingType")
	}

}
