version := "v0.4.0"

build:
    go build -ldflags "-X github.com/avadhanij/cluide/cmd.versionString={{version}}" -o=./dist/cluide

localdist:
    cp ./dist/cluide $HOME/.local/bin/cluide