package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"theztd/gomi/git"
)

var text_file string = `
# HELP gomi_uncommited_files Sum of uncommited files
# TYPE gomi_uncommited_files histogram
gomi_uncommited_files{git_path="{{ .Path }}"} {{ .ChangedCount }}

# HELP gomi_commit_status Last status of commit (0=Ok, 1=Fail, 90=Nothing to do, 99=dry run)
# TYPE gomi_commit_status histogram
gomi_commit_status{git_path="{{ .Path }}"} {{ .Status }}
`

func prometheusReporter(path string, repo *git.Git) {
	log.Println("INFO: report to", path)

	// Prepare metrics dir
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		log.Panicln("ERR (mkdir):", err)
	}

	f, err := os.Create(path)
	if err != nil {
		log.Println("ERR:", err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("# Metrics for gomi in dir %s", repo.Path))
	tmpl, err := template.New("metrics").Parse(text_file)
	if err != nil {
		log.Println("ERR:", err)
	}
	tmpl.Execute(f, repo)

	f.Sync()
}
