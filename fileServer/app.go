package fileserver

import (
	"html/template"
)

type FileServerData struct {
	Data template.HTML
	Time string
}
