package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
)

const (
	projectSuffix         = "sublime-project"
	baseSublimeConfigPath = "base/base.sublime-project"
)

type config struct {
	c *cli.Context
}

func (v config) group() string {
	return v.c.Args().Get(0)
}

func (v config) groupPath() string {
	return path.Join(basePath(v.c), v.c.Args().Get(0))
}

func (v config) projectName() string {
	return v.c.Args().Get(1)
}

func (v config) projectDir() string {
	d, _ := os.Getwd()
	return path.Join(d, v.projectName())
}

func (v config) projectSettingPath() string {
	name := v.c.Args().Get(2)
	if name == "" {
		name = path.Base(v.projectName())
	}
	return path.Join(v.groupPath(), name) + "." + projectSuffix
}

func (v config) srcConfigPath() string {
	configPath := path.Join(basePath(v.c), v.group()+"."+projectSuffix)
	if _, err := os.Stat(configPath); err == nil {
		return configPath
	}

	return baseConfigPath(v.c)
}

func (v config) isValidCreate() bool {
	return len(v.c.Args()) >= 2
}

func basePath(c *cli.Context) string {
	for _, ca := range [...]string{c.String("base")} {
		if ca != "" {
			return ca
		}
	}
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, ".subpro") + "/"
}

func baseConfigPath(c *cli.Context) string {
	return path.Join(basePath(c), baseSublimeConfigPath)
}

func (v config) isExistFile() bool {
	flag := false
	name := v.projectName() + "." + projectSuffix

	filepath.Walk(basePath(v.c), func(p string, info os.FileInfo, err error) error {
		if flag {
			return filepath.SkipDir
		}

		if info.IsDir() {
			return nil
		}

		if name == info.Name() {
			flag = true
		}

		return nil
	})

	return flag
}

func main() {
	app := cli.NewApp()
	app.Name = "subpro"
	app.Version = "2.0.2"
	app.Usage = "management sublime text project"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "base, b",
			Usage: "define base path",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "create project",
			Action: func(c *cli.Context) {
				v := config{c}
				if !v.isValidCreate() {
					log.Fatal("please input group and project path")
				}
				if v.isExistFile() {
					log.Fatal("Already file exists")
				}
				os.MkdirAll(v.groupPath(), 0755)
				cmd := copyFile(v.srcConfigPath(), v.projectSettingPath())
				cmd.Run()

				// overwrite project path
				file, err := ioutil.ReadFile(v.projectSettingPath())
				if err != nil {
					log.Fatal(err)
				}
				sublSetting := UnMarshal(file)
				sublSetting.Folders[0].Path = v.projectDir()
				err = ioutil.WriteFile(v.projectSettingPath(), Marshal(sublSetting), 0644)
				if err != nil {
					log.Fatal(err)
				}

				log.Println("Create", v.projectName())
				cmd = openSublText(v.projectSettingPath())
				cmd.Run()
			},
		},
		{
			Name:      "delete",
			ShortName: "d",
			Usage:     "delete project",
			Action: func(c *cli.Context) {
				projectName := c.Args().Get(0)
				if projectName == "" {
					log.Fatal("please input want to delete a project name")
				}

				filepath.Walk(basePath(c), func(p string, info os.FileInfo, err error) error {
					if info.IsDir() {
						return nil
					}

					name := strings.Split(path.Base(p), ".")[0]
					if name == projectName {
						log.Println("Delete", p)
						cmd := deleteFile(p)
						cmd.Run()
					}

					return nil
				})
			},
		},
	}

	app.Action = func(c *cli.Context) {
		projectName := c.Args().Get(0)

		filepath.Walk(basePath(c), func(p string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if !strings.HasSuffix(p, projectSuffix) {
				return nil
			}

			name := strings.Split(path.Base(p), ".")[0]
			if name == projectName {
				log.Println("Open", name)
				cmd := openSublText(p)
				cmd.Run()
			}

			return nil
		})
	}

	app.Run(os.Args)
}
