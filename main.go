package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"theztd/gomi/git"
)

func main() {
	var path, prom_path string
	var check bool
	flag.StringVar(&path, "path", "", "Absolute path to the git repo root, default is curent dir.")
	flag.BoolVar(&check, "check", false, "Only show changes but don't do anything.")
	flag.StringVar(&prom_path, "prom-path", "", "Absolute path to node_exporter tex-file path, when is not defined, nothing will be generated.")

	flag.Parse()

	log.Printf("Git watch has been starter (with flags path: %s, check: %t)", path, check)

	// ensure that given path is absolute
	absPath, _ := filepath.Abs(path)

	repo := git.New(absPath)
	if prom_path != "" {
		defer prometheusReporter(prom_path, &repo)
	}

	repo.Config = git.GitConfig{
		Author: "Automat",
		Email:  "automat@gin05",
	}
	repo.GetChanges()

	/*
		Run only if there are any uncommited changes in the repo
	*/
	if repo.ChangedCount >= 1 {
		// Prepare commit message
		hostname, _ := os.Hostname()
		msg := fmt.Sprintf("Automat commits %d files on server %s\n----- Detail -----\n%s",
			repo.ChangedCount,
			hostname,
			repo.ChangedFilesRaw)

		// If check mode, print message and exit 99
		if check {
			log.Println("Run in CHECK mode. Commit change will be:")
			log.Println(msg)
			repo.Status = 99
			os.Exit(99)
		}

		// Add files to commit
		if repo.AddAll() != nil {
			log.Panicln("Unable to add files before commit!")
		}

		// commit changes with defined message
		if repo.Commit(msg) != nil {
			log.Panicln("Unable to commit changes!")
		}

		repo.Status = 0

	} else {
		log.Println("Nothing to commit, repository is clear!")
		repo.Status = 90
	}

	//log.Println(CommitIfChanged(path, check))
}
