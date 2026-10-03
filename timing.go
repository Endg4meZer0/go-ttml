package v1

import (
	"math"
	"strconv"
	"strings"
)

func parseTiming(in string) (out int) {
	if len(in) == 0 {
		return -1
	}

	parts := strings.Split(in, ":")
	switch len(parts) {
	case 3: // HH:MM:SS.mmm
		h, err := strconv.Atoi(parts[0])
		if err != nil {
			return -1
		}
		m, err := strconv.Atoi(parts[1])
		if err != nil {
			return -1
		}
		secs, _ := strings.CutSuffix(parts[2], "s")
		s, err := strconv.ParseFloat(secs, 64)
		if err != nil {
			return -1
		}

		return h*60*60*1000 + m*60*1000 + int(math.Round(s*1000))
	case 2: // MM:SS.mmm
		m, err := strconv.Atoi(parts[0])
		if err != nil {
			return -1
		}
		secs, _ := strings.CutSuffix(parts[1], "s")
		s, err := strconv.ParseFloat(secs, 64)
		if err != nil {
			return -1
		}

		return m*60*1000 + int(math.Round(s*1000))
	case 1:
		secs, _ := strings.CutSuffix(in, "s")
		s, err := strconv.ParseFloat(secs, 64)
		if err != nil {
			return -1
		}

		return int(math.Round(s * 1000))
	}

	return -1
}
