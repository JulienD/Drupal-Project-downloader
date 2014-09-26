package projectManager

import "fmt"
import "github.com/codeskyblue/go-sh"


type Project struct {
    Name    string
}

// Generates the project's git url
func (p *Project) GetGitUrl() string {
    return  "http://git.drupal.org/project/" + p.Name + ".git"
}

// Download a project in a given dir
func (p *Project) DownloadProject(path string) bool {
    // Creates a shell session to allow to chaine commands in order to keep
    // variables.
    session := sh.NewSession()
    session.SetEnv("BUILD_ID", "123")
    session.SetDir(path)
    fmt.Printf("Downloading: %s \n",  p.Name)

    // Deletes the project dir and download via git the project.
    err := session.Command("rm", "-rf", p.Name).Command("git", "clone", p.GetGitUrl(), p.Name).Run()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func NewProject(name string) *Project {
    return &Project{
        Name: name,
    }
}
