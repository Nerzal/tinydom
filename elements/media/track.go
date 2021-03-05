package media

import "github.com/Nerzal/tinydom"

// <track src="fgsubtitles_en.vtt" kind="subtitles" srclang="en" label="English">

type Track struct {
	*tinydom.Element
}

func NewTrack(src, kind, srcLang, label string) *Track {
	trackElem := tinydom.GetDocument().CreateElement("track")
	trackElem.Set("src", src)
	trackElem.Set("kind", kind)
	trackElem.Set("srcLang", srcLang)
	trackElem.Set("label", label)

	return &Track{trackElem}
}
