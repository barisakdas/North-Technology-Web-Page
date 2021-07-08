package helpers

import (
	. "NorthTechWebPage/Log"
	"path/filepath"
)

func Include(path string) []string {
	files, err := filepath.Glob("Admin/views/templates/*.html")
	if err != nil {
		LogJson("Site","Error","Helpers","Include","Could not pull html files in templates folder!!",err.Error())
	}

	pathFiles, err := filepath.Glob("Admin/views/"+path+"*.html")
	if err != nil {
		LogJson("Site","Error","Helpers","Include","Could not pull html files in templates folder!!",err.Error())
	}

	for _, file := range pathFiles {
		files = append(files, file)
	}

	return files
}
