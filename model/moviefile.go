package model

import (
	"fmt"
	"time"
)

const (
	gibBytes = 1024 * 1024 * 1024
	mibBytes = 1024 * 1024
	kibBytes = 1024
)

//MovieFile contains all information of a single movie file
type MovieFile struct {
	Title           string
	Path            string
	Filename        string
	Dir             string
	Size            uint64
	Duration        time.Duration
	Format          string
	VideoStreams    []VideoStream
	AudioStreams    []AudioStream
	SubtitleStreams []SubtitleStream
}

//MovieFileExternal is the data format for external usage
type MovieFileExternal struct {
	Title           string           `json:"title"`
	Filename        string           `json:"filename"`
	Dir             string           `json:"dir"`
	SizeBytes       uint64           `json:"size_bytes"`
	SizeText        string           `json:"size_text"`
	DurationSec     uint32           `json:"duration_seconds"`
	DurationText    string           `json:"duration_text"`
	Format          string           `json:"format"`
	VideoStreams    []VideoStream    `json:"videostreams"`
	AudioStreams    []AudioStream    `json:"audiostreams"`
	SubtitleStreams []SubtitleStream `json:"subtitlestreams"`
	SubtitleFiles   []SubtitleFile   `json:"subtitlefiles"`
}

//FromMovieFile convert a db model to presentation model
func (me *MovieFileExternal) FromMovieFile(m *MovieFile) {
	me.Title = m.Title
	me.Filename = m.Filename
	me.Dir = m.Dir
	me.SizeBytes = m.Size
	me.SizeText = bytesText(m.Size)
	me.DurationText = m.Duration.String()
	me.DurationSec = uint32(m.Duration.Seconds())
	me.Format = m.Format
	me.VideoStreams = append(me.VideoStreams, m.VideoStreams...)
	me.AudioStreams = append(me.AudioStreams, m.AudioStreams...)
	me.SubtitleStreams = append(me.SubtitleStreams, m.SubtitleStreams...)
}

//VideoStream contains the information of a single video stream
type VideoStream struct {
	Width    int
	Height   int
	Codec    string
	BitRate  int
	Language string
}

//AudioStream contains the information of a single audio stream
type AudioStream struct {
	Codec         string
	BitRate       int
	Channels      uint
	ChannelLayout string
	Language      string
}

//SubtitleStream contains the information of a single audio stream
type SubtitleStream struct {
	Codec     string
	Languages []string
	Title     string
}

func bytesText(size uint64) string {
	if size > gibBytes {
		return fmt.Sprintf("%.2f GiB", float64(size)/gibBytes)
	} else if size > mibBytes {
		return fmt.Sprintf("%.2f MiB", float64(size)/mibBytes)
	} else {
		return fmt.Sprintf("%.2f KiB", float64(size)/kibBytes)
	}
}
