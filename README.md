## subpro

it's management project of Sublime Text with terminal.

### install

use homebrew for macOS

```shell
$ brew tap satoshun/commands
$ brew install subpro
$ brew install subpro-completion
```

add under code to .bashrc

```shell
if [ -f `brew --prefix`/etc/bash_completion.d/subpro_completion ]; then
    source `brew --prefix`/etc/bash_completion.d/subpro_completion
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
$ subpro
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
