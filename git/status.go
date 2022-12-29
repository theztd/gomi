package git

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func (g *Git) GetChanges() []GitFile {
	if os.Chdir(g.Path) != nil {
		log.Panic("Unable to open directory %s", g.Path)
	}

	log.Println(os.Getwd())

	cmd := exec.Command("git", "status", "--porcelain")

	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.Run()

	g.ChangedFilesRaw = out.String()
	changes := strings.Split(g.ChangedFilesRaw, "\n")

	g.ChangedCount = len(changes) - 1
	if g.ChangedCount < 1 {
		return []GitFile{}
	}

	var ret []GitFile
	for _, l := range changes {
		if len(l) > 1 {
			ret = append(ret, GitFile{
				Code: strings.TrimSpace(l[0:2]),
				Path: strings.TrimSpace(l[3:]),
			})
		}
	}
	g.ChangedFiles = ret
	return ret
}
