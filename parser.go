package parser

import (
	"errors"
	"regexp"
	"strings"
)

var videoRegexpList = []*regexp.Regexp{
	//regexp.MustCompile(`.*(?:youtu.be\/|v\/|u\/|\w\/|embed\/|shorts\/|watch\?v=)([^#\&\?]*).*`),
	regexp.MustCompile(`(?:v|embed|shorts|watch\?v)i?(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`([^"&?/=%]{11})`),
}

func ExtractID(url string) (string, error) {
	videoID := "UNKNOWNID"
	if strings.Contains(url, "youtu") || strings.ContainsAny(url, "\"?&/<%=") {
		videoID = url
		for _, re := range videoRegexpList {
			if isMatch := re.MatchString(videoID); isMatch {
				subs := re.FindStringSubmatch(videoID)
				videoID = subs[1]
			}
		}
	}
	if strings.ContainsAny(videoID, "?&/<%=") {
		return "", errors.New("Invalid characters in video id")
	}
	if len(videoID) < 10 {
		return "", errors.New("The video id must be at least 10 characters long")
	}
	return videoID, nil
}
