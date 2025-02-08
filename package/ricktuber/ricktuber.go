package ricktuber

import "regexp"

type ricktuber struct {
	youtubeUrlsMatcher *regexp.Regexp
	ricktubeUrlTemplate string
}

func (r ricktuber) FindYoutubeVideos(text string) []string {
	foundLinks := r.youtubeUrlsMatcher.FindAllStringSubmatch(text, -1)
	videoIds := make([]string, 0, len(foundLinks))
	for _, match := range foundLinks {
		videoIds = append(videoIds, match[1])
	}

	return videoIds
}


func (r ricktuber) GetLink(youtubeVideoId string) string {
	return r.ricktubeUrlTemplate + youtubeVideoId
}

var Ricktuber = ricktuber{
	youtubeUrlsMatcher:  regexp.MustCompile(`(?:https?://)?(?:www\.)?(?:youtube\.com/(?:shorts/|watch\?v=)|youtu\.be/)([a-zA-Z0-9_-]{11})`),
	ricktubeUrlTemplate: "https://ricktube.ru/watch/",
}