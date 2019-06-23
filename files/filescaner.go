package files

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/moviegeek/moviesub/model"
	ffmpegmodel "github.com/xfrr/goffmpeg/models"
	"github.com/xfrr/goffmpeg/transcoder"
)

//ScanCallback callback for filescaner to call when find a media file
type ScanCallback func(*model.MovieFile)

//ScanMediaFiles scans the root directory and find all media files
func (fs *FileScaner) ScanMediaFiles(cb ScanCallback) {
	err := filepath.Walk(fs.root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if isMediaFile(info.Name()) {

			log.Printf("probing file: %s", path)
			mediaFile, err := probeMediaFile(path)
			if err != nil {
				log.Printf("something wrong with media file [%s]: %v, skip it", path, err)
				return nil
			}

			m := &model.MovieFile{}

			absPath, err := filepath.Abs(path)
			if err != nil {
				log.Printf("failed to get abs path from %s, use non-abs path. %v", path, err)
				absPath = ""
			}

			m.Path = absPath
			m.Dir, m.Filename = filepath.Split(absPath)

			m.Title = m.Filename[:strings.LastIndex(m.Filename, ".")]

			convertMediaFile(mediaFile, m)

			cb(m)
		}

		return nil
	})

	if err != nil {
		log.Printf("error when scaning folder %s: %v", fs.root, err)
	}
}

func isMediaFile(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".mkv" || ext == ".mp4"
}

func addStream(movieFile *model.MovieFile, stream ffmpegmodel.Streams) {
	if stream.CodecType == "video" {
		addVideoStream(movieFile, stream)
	} else if stream.CodecType == "audio" {
		addAudioStream(movieFile, stream)
	} else if stream.CodecType == "subtitle" {
		addSubtitleStream(movieFile, stream)
	} else {
		log.Printf("unknown stream type: '%s', skip", stream.CodecType)
	}
}

func addVideoStream(movieFile *model.MovieFile, stream ffmpegmodel.Streams) {
	//TODO: add bitrate later
	vs := model.VideoStream{}
	vs.Width = stream.Width
	vs.Height = stream.Height
	vs.Codec = stream.CodecName

	movieFile.VideoStreams = append(movieFile.VideoStreams, vs)
}

func addAudioStream(movieFile *model.MovieFile, stream ffmpegmodel.Streams) {
	as := model.AudioStream{}
	as.Codec = stream.CodecName

	movieFile.AudioStreams = append(movieFile.AudioStreams, as)
}

func addSubtitleStream(movieFile *model.MovieFile, stream ffmpegmodel.Streams) {
	ss := model.SubtitleStream{}
	ss.Codec = stream.CodecName
	ss.Languages = strings.Split(stream.Tags.Language, ",")

	movieFile.SubtitleStreams = append(movieFile.SubtitleStreams, ss)
}

func probeMediaFile(filepath string) (*ffmpegmodel.Mediafile, error) {
	inputFile := filepath
	outputFile := "test.mkv"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputFile, outputFile)
	if err != nil {
		log.Printf("failed to probe the media file [%s]: %v", filepath, err)
		return nil, err
	}

	mediaFile := trans.MediaFile()

	return mediaFile, nil
}

func convertMediaFile(mediaFile *ffmpegmodel.Mediafile, movieFile *model.MovieFile) {
	format := mediaFile.Metadata().Format
	if s, err := strconv.Atoi(format.Size); err == nil {
		movieFile.Size = uint64(s)
	}

	if f, err := strconv.ParseFloat(format.Duration, 32); err == nil {
		d := int(f)
		movieFile.Duration = time.Duration(d) * time.Second
	}

	streams := mediaFile.Metadata().Streams
	for _, s := range streams {
		addStream(movieFile, s)
	}
}
