package media

import (
	"strconv"

	"github.com/Nerzal/tinydom"
)

const videoAlt = "Your browser does not support the video tag."

const (
	MP4  = "video/mp4"
	OGG  = "video/ogg"
	WEBM = "video/webm"
)

type Video struct {
	*tinydom.Element
}

type VideoSource struct {
	Source string
	Type   string
}

func NewVideo() *Video {
	videoElem := tinydom.GetDocument().CreateElement("video")
	return &Video{videoElem}
}

func NewVideoParams(
	width, height int,
	autoplay, muted, controls bool,
	sources ...*VideoSource,
) *Video {
	result := NewVideo()
	result.SetHeight(height)
	result.SetWidth(width)

	if autoplay {
		result.SetAutoplay()
	}

	if muted {
		result.SetMuted()
	}

	if controls {
		result.SetControl()
	}

	for i := range sources {
		source := sources[i]
		result.SetSource(source)
	}

	return result
}

func (v *Video) SetAutoplay() *Video {
	v.Set("autoplay", true)
	return v
}

func (v *Video) SetMuted() *Video {
	v.Set("muted", true)
	return v
}

func (v *Video) SetControl() *Video {
	v.Set("controls", "controls")
	return v
}

func (v *Video) SetAltText() *Video {
	v.SetInnerHTML(videoAlt)
	return v
}

func (v *Video) Reload() *Video {
	v.Call("load")
	return v
}

func (v *Video) SetSource(source *VideoSource) *Video {
	src := tinydom.GetDocument().CreateElement("source")
	src.Set("src", source.Source)
	src.Set("type", source.Type)

	v.AppendChild(src)
	return v
}

func (v *Video) SetWidth(width int) *Video {
	v.Set("width", strconv.Itoa(width))
	return v
}

func (v *Video) Width() (int, error) {
	width := v.Get("width").String()
	result, err := strconv.Atoi(width)
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (v *Video) SetHeight(height int) *Video {
	v.Set("height", strconv.Itoa(height))
	return v
}

func (v *Video) Height() (int, error) {
	height := v.Get("height").String()
	result, err := strconv.Atoi(height)
	if err != nil {
		return -1, err
	}

	return result, nil
}
