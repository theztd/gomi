package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (g *Git) AddAll() (err error) {
	if os.Chdir(g.Path) != nil {
		log.Panic("Unable to open directory %s", g.Path)
	}

	cmd := exec.Command("git", "add", ".")
	return cmd.Run()
}

func (g *Git) Commit(msg string) (err error) {
	if os.Chdir(g.Path) != nil {
		log.Panic("Unable to open directory %s", g.Path)
	}

	author := fmt.Sprintf("--author=\"%s\"", g.AuthorString())
	cmd := exec.Command("git", "commit", "-am", msg, author)

	return cmd.Run()
}
