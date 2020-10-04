package app

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type AppType struct {
	Folder         string
	FolderAndSlash string
	Name           string
	NameWithExe    string
	FilePath       string
	ConfigPath     string
}

func (app *AppType) Setup() {
	exe, _ := os.Executable()
	// strange problem: if FolderAndSlash is used instead of f, the change in FolderAndSlash is not visible outside of this function.
	f, n := filepath.Split(exe)
	app.FolderAndSlash = f
	app.NameWithExe = n
	app.FilePath = app.FolderAndSlash + app.NameWithExe
	app.Folder = app.FolderAndSlash[:len(app.FolderAndSlash)-1] // get rid of trailing backslash of folder
	app.Name = app.NameWithExe
	if len(app.Name) >= 4 && app.Name[len(app.Name)-4:] == ".exe" {
		app.Name = app.Name[:len(app.Name)-4] // get rid of .exe of name
	}

	ff, err := os.OpenFile(app.FolderAndSlash+"log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	if _, err := toml.DecodeFile(app.FolderAndSlash+"config.toml", &configFile); err != nil {
		log.Printf("%s", err)
	}

	app.ConfigPath = configFile.ConfigPath

}

func DecodeFileTOML(path string, configPtr *interface{}) {
	if _, err := toml.DecodeFile(path, configPtr); err != nil {
		log.Printf("%s, %s", path, err)
	}
}
