build:
    go build -o=dist/cluide

localdist:
    cp ./dist/cluide $HOME/.local/bin/cluide

clean:
    rm -rf dist/
