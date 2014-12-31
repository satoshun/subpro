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
	projectSuffix  = "sublime-project"
	baseConfigPath = "base/base.sublime-project"
)

type Var struct {
	c *cli.Context
}

func (v Var) Group() string {
	return v.c.Args().Get(0)
}

func (v Var) GroupPath() string {
	return path.Join(BasePath(v.c), v.c.Args().Get(0))
}

func (v Var) ProjectName() string {
	return v.c.Args().Get(1)
}

func (v Var) projectDir() string {
	d, _ := os.Getwd()
	return path.Join(d, v.ProjectName())
}

func (v Var) projectSettingPath() string {
	name := v.c.Args().Get(2)
	if name == "" {
		name = path.Base(v.ProjectName())
	}
	return path.Join(v.GroupPath(), name) + "." + projectSuffix
}

func (v Var) srcConfigPath() string {
	configPath := path.Join(BasePath(v.c), v.Group()+"."+projectSuffix)
	if _, err := os.Stat(configPath); err == nil {
		return configPath
	}

	return BaseConfigPath(v.c)
}

func (v Var) isValidCreate() bool {
	return len(v.c.Args()) >= 2
}

func BasePath(c *cli.Context) string {
	for _, ca := range [...]string{c.String("base")} {
		if ca != "" {
			return ca
		}
	}
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, ".subpro") + "/"
}

func BaseConfigPath(c *cli.Context) string {
	return path.Join(BasePath(c), baseConfigPath)
}

func (v Var) isExistFile() bool {
	flag := false
	name := v.ProjectName() + "." + projectSuffix

	filepath.Walk(BasePath(v.c), func(p string, info os.FileInfo, err error) error {
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
	app.Version = "2.0.1"
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
				v := Var{c}
				if !v.isValidCreate() {
					log.Fatal("please input group and project path")
				}
				if v.isExistFile() {
					log.Fatal("Already file exists")
				}
				os.MkdirAll(v.GroupPath(), 0755)
				cmd := CopyFile(v.srcConfigPath(), v.projectSettingPath())
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

				log.Println("Create", v.ProjectName())
				cmd = OpenSublText(v.projectSettingPath())
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

				filepath.Walk(BasePath(c), func(p string, info os.FileInfo, err error) error {
					if info.IsDir() {
						return nil
					}

					name := strings.Split(path.Base(p), ".")[0]
					if name == projectName {
						log.Println("Delete", p)
						cmd := DeleteFile(p)
						cmd.Run()
					}

					return nil
				})
			},
		},
	}

	app.Action = func(c *cli.Context) {
		projectName := c.Args().Get(0)

		filepath.Walk(BasePath(c), func(p string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if !strings.HasSuffix(p, projectSuffix) {
				return nil
			}

			name := strings.Split(path.Base(p), ".")[0]
			if name == projectName {
				log.Println("Open", name)
				cmd := OpenSublText(p)
				cmd.Run()
			}

			return nil
		})
	}

	app.Run(os.Args)
}
