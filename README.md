## subpro

it's management project of Sublime Text with terminal.

### install

use homebrew for macOS.

```shell
brew tap satoshun/commands
brew install subpro

or

go get -u -v github.com/satoshun/subpro/cmd
```

add completion for bash.

```shell
$ brew install subpro-completion

# add under code to .bashrc
if [ -f `brew --prefix`/etc/bash_completion.d/subpro_completion ]; then
    source `brew --prefix`/etc/bash_completion.d/subpro_completion
fi
```

add completion for zsh.

```shell
$ brew install subpro-zcompletion

# add under code to .zshrc
if [ -f `brew --prefix`/etc/zsh_completion.d/subpro_zcompletion ]; then
    source `brew --prefix`/etc/zsh_completion.d/subpro_zcompletion
fi
```

### usage

```shell
subpro -h
NAME:
   subpro - management sublime text project

USAGE:
   subpro [global options] command [command options] [arguments...]

VERSION:
   2.0.0

COMMANDS:
   create, c  create project
   delete, d  delete project
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --base, -b ''  define base path
   --version, -v  print the version
   --help, -h   show help
```

### Example

project list && open project

```shell
subpro [tab]
Display all 131 possibilities? (y or n)
Amalgam                                 beaker                                  ghost                                   roboguice
Android-Universal-Image-Loader          bitcoin                                 gist_css                                routes
...
subpro ghost
# open ghost project with sublime text]
```

create project

```shell
ls
martini

subpro create golang martini

# create martini project into golang directory.

...
$ subpro martini
# open martini project with sublime text
```

delete project

```shell
subpro delete martini
```

### License

MIT(http://opensource.org/licenses/mit-license.php)
