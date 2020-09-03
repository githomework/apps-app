package app

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type appType struct {
	Folder         string
	FolderAndSlash string
	Name           string
	NameWithExe    string
	FilePath       string
	ConfigPath     string
}

var (
	App appType
)

func init() {
	exe, _ := os.Executable()
	// strange problem: if FolderAndSlash is used instead of f, the change in FolderAndSlash is not visible outside of this function.
	f, n := filepath.Split(exe)
	App.FolderAndSlash = f
	App.NameWithExe = n
	App.FilePath = App.FolderAndSlash + App.NameWithExe
	App.Folder = App.FolderAndSlash[:len(App.FolderAndSlash)-1] // get rid of trailing backslash of folder
	App.Name = App.NameWithExe
	if len(App.Name) >= 4 && App.Name[len(App.Name)-4:] == ".exe" {
		App.Name = App.Name[:len(App.Name)-4] // get rid of .exe of name
	}

	ff, err := os.OpenFile(App.FolderAndSlash+App.Name+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	log.SetOutput(ff)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Config entry for the app, may reside in config.toml in the current folder, or another .toml file in another folder.
	// The path of the .toml file is stated in config.toml in the current folder.
	var configFile struct {
		ConfigPath string `toml:"config_file"`
	}

	if _, err := toml.DecodeFile(App.FolderAndSlash+"config.toml", &configFile); err != nil {
		log.Printf("%s", err)
	}

	App.ConfigPath = configFile.ConfigPath

}
