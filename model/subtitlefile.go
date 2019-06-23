package model

//SubtitleFile contains the informantion of an external subtitle file
type SubtitleFile struct {
	Title    string
	Filename string
	Dir      string
	Language []string
	Format   string
}
