package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/satoshun/subpro"
)

func main() {
	app := cli.NewApp()
	app.Name = "subpro"
	app.Version = "2.1.0-SNAPSHOT"
	app.Author = "satoshun"
	app.Email = ""
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
			Action: func(c *cli.Context) error {
				v := subpro.Config{c}
				if v.IsInValidArgs() {
					log.Fatal("please input group and project path")
				}
				if v.IsExist() {
					log.Fatal("Already file exists")
				}
				checkErr(v.CreateDir(0755))
				checkErr(subpro.CopyFile(v.ProjectSettingPath(), v.SrcConfigPath()))

				// overwrite project file
				file, err := ioutil.ReadFile(v.ProjectSettingPath())
				checkErr(err)
				sublSetting := subpro.UnMarshalSetting(file)
				sublSetting.Folders[0].Path = v.ProjectDir()

				checkErr(ioutil.WriteFile(v.ProjectSettingPath(), subpro.MarshalSetting(sublSetting), 0644))

				log.Println("Create", v.ProjectName())
				cmd := subpro.OpenCommand(v.ProjectSettingPath())
				checkErr(cmd.Run())

				return nil
			},
		},
		{
			Name:      "delete",
			ShortName: "d",
			Usage:     "delete project",
			Action: func(c *cli.Context) error {
				projectName := c.Args().Get(0)
				if projectName == "" {
					log.Fatal("please input want to delete a project name")
				}

				filepath.Walk(subpro.BasePath(c), func(p string, info os.FileInfo, err error) error {
					if info.IsDir() {
						return nil
					}

					name := strings.Split(path.Base(p), ".")[0]
					if name == projectName {
						log.Println("Delete", p)
						checkErr(os.Remove(p))
					}

					return nil
				})
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		projectName := c.Args().Get(0)
		if projectName == "" {
			return errors.New("nothing params")
		}

		filepath.Walk(subpro.BasePath(c), func(p string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if !subpro.IsSublimeFile(p) {
				return nil
			}

			name := strings.Split(path.Base(p), ".")[0]
			if name == projectName {
				log.Println("Open", name)
				cmd := subpro.OpenCommand(p)
				checkErr(cmd.Run())
			}

			return nil
		})

		return nil
	}

	app.Run(os.Args)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
