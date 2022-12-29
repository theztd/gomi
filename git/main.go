package git

import "fmt"

/*
	Object definition
*/

type GitFile struct {
	Code string
	Path string
}

type GitConfig struct {
	Author string
	Email  string
}

type Git struct {
	Path            string    // path to git root
	ChangedCount    int       // count of changes in repo
	ChangedFiles    []GitFile // list changed files struct
	ChangedFilesRaw string    // list changed files string
	Config          GitConfig // configuration for commit (author, email)
	Status          uint8     // Commit status (0=Ok, 1=Fail, 90=Nothing to do, 99=dry run)
}

/*
	Initialize new object
*/
func New(path string) Git {
	return Git{
		Path:   path,
		Status: 1, // start with commit state Error
	}
}

func (g *Git) AuthorString() string {
	return fmt.Sprintf("%s <%s>", g.Config.Author, g.Config.Email)
}
