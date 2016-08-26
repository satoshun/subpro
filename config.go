package subpro

import (
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

type Config struct {
	C *cli.Context
}

func (v *Config) Group() string {
	return v.C.Args().Get(0)
}

func (v *Config) GroupPath() string {
	return path.Join(BasePath(v.C), v.C.Args().Get(0))
}

func (v *Config) ProjectName() string {
	return v.C.Args().Get(1)
}

func (v *Config) ProjectDir() string {
	d, _ := os.Getwd()
	return path.Join(d, v.ProjectName())
}

func (v *Config) ProjectSettingPath() string {
	name := v.C.Args().Get(2)
	if name == "" {
		name = path.Base(v.ProjectName())
	}
	return path.Join(v.GroupPath(), name) + "." + projectSuffix
}

func (v *Config) SrcConfigPath() string {
	configPath := path.Join(BasePath(v.C), v.Group()+"."+projectSuffix)
	if _, err := os.Stat(configPath); err == nil {
		return configPath
	}

	return BaseConfigPath(v.C)
}

func (v *Config) IsInValidArgs() bool {
	return len(v.C.Args()) < 2
}

func (v *Config) CreateDir(perm os.FileMode) error {
	return os.MkdirAll(v.GroupPath(), perm)
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

func IsSublimeFile(path string) bool {
	return strings.HasSuffix(path, projectSuffix)
}

func BaseConfigPath(c *cli.Context) string {
	return path.Join(BasePath(c), baseSublimeConfigPath)
}

func (v *Config) IsExist() bool {
	flag := false
	name := v.ProjectName() + "." + projectSuffix

	filepath.Walk(BasePath(v.C), func(p string, info os.FileInfo, err error) error {
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
