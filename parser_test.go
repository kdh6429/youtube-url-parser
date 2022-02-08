package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVailedURL(t *testing.T) {
	assert := assert.New(t)
	vaildURLs := []string{
		"http://www.youtube.com/watch?v=wDfqXR_5yyQ&feature=feedrec_grec_index",
		"http://www.youtube.com/v/wDfqXR_5yyQ?fs=1&amp;hl=en_US&amp;rel=0",
		"http://www.youtube.com/watch?v=wDfqXR_5yyQ#t=0m10s",
		"http://www.youtube.com/embed/wDfqXR_5yyQ?rel=0",
		"http://www.youtube.com/watch?v=wDfqXR_5yyQ",
		"http://youtu.be/wDfqXR_5yyQ",
		"//www.youtube-nocookie.com/embed/wDfqXR_5yyQ?rel=0",
		"http://www.youtube.com/watch?v=wDfqXR_5yyQ&feature=channel",
		"http://www.youtube.com/watch?v=wDfqXR_5yyQ&playnext_from=TL&videos=osPknwzXEas&feature=sub",
		"http://www.youtube.com/ytscreeningroom?v=wDfqXR_5yyQ",
		"http://youtu.be/wDfqXR_5yyQY",
		"http://www.youtube.com/watch?v=wDfqXR_5yyQY&feature=youtu.be",
		"http://youtu.be/wDfqXR_5yyQ",
		"http://www.youtube.com/watch?v=wDfqXR_5yyQ&feature=channel",
		"http://www.youtube.com/watch?v=wDfqXR_5yyQ&playnext_from=TL&videos=osPknwzXEas&feature=sub",
		"http://www.youtube.com/ytscreeningroom?v=wDfqXR_5yyQ",
		"http://www.youtube.com/embed/wDfqXR_5yyQ?rel=0",
		"http://www.youtube.com/watch?v=wDfqXR_5yyQ",
		"http://youtube.com/v/wDfqXR_5yyQ?feature=youtube_gdata_player",
		"http://youtube.com/vi/wDfqXR_5yyQ?feature=youtube_gdata_player",
		"http://youtube.com/?v=wDfqXR_5yyQ&feature=youtube_gdata_player",
		"http://www.youtube.com/watch?v=wDfqXR_5yyQ&feature=youtube_gdata_player",
		"http://youtube.com/?vi=wDfqXR_5yyQ&feature=youtube_gdata_player",
		"http://youtube.com/watch?v=wDfqXR_5yyQ&feature=youtube_gdata_player",
		"http://youtube.com/watch?vi=wDfqXR_5yyQ&feature=youtube_gdata_player",
		"http://youtu.be/wDfqXR_5yyQ?feature=youtube_gdata_player",
	}

	for _, url := range vaildURLs {
		id, err := ExtractID(url)
		if err != nil {
			t.Errorf("%v, url=%v result=%v expect=%v", err, url, id, "wDfqXR_5yyQ")
		}
		if !assert.Equal("wDfqXR_5yyQ", id) {
			t.Errorf("Wrong ID, url=%v result=%v expect=%v", url, id, "wDfqXR_5yyQ")
		}
	}
}
