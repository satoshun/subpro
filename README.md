## subpro

it's management project of Sublime Text with terminal.

### install

use homebrew for macOS.

```shell
$ brew tap satoshun/commands
$ brew install subpro
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
$ subpro
Example usage:
  subpro [project]
  subpro create|c [package] [project dir]
  subpro delete|del|d [project]
```

### Example

project list && open project

```shell
$ subpro [tab]
Display all 131 possibilities? (y or n)
Amalgam                                 beaker                                  ghost                                   roboguice
Android-Universal-Image-Loader          bitcoin                                 gist_css                                routes
...
$ subpro ghost
# open ghost project with sublime text]
```

create project

```shell
$ ls
martini
$ subpro create golang martini

# create martini project into golang directory.

...
$ subpro martini
# open martini project with sublime text
```

delete project

```shell
$ subpro delete martini
```

### License

MIT(http://opensource.org/licenses/mit-license.php)
