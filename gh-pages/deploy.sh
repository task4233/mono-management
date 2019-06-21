#!/bin/bash

GITROOTPATH="$(git rev-parse --show-toplevel)"
docker run -t --rm -v "$GITROOTPATH/gh-pages":/gh-pages -v "$GITROOTPATH/backend":/go/src/app golang:1.12.5 /bin/bash -c "
    cd /go/src/app &&
    rm -rf ./vendor &&
    rm -rf /gh-pages/dist/backend &&
    mkdir -p /gh-pages/dist/backend/internal &&
    mkdir -p /gh-pages/dist/backend/router &&
    go get golang.org/x/tools/cmd/godoc &&
    cp /go/src/golang.org/x/tools/godoc/static/jquery.js /gh-pages/dist/backend/ &&
    cp /go/src/golang.org/x/tools/godoc/static/godocs.js /gh-pages/dist/backend/ &&
    cp /go/src/golang.org/x/tools/godoc/static/style.css /gh-pages/dist/backend/ &&
    godoc -url 'http://localhost:6060/pkg/app?m=all' | sed 's/\/lib\/godoc/\/mono-management\/backend/g' > /gh-pages/dist/backend/index.html &&
    godoc -url 'http://localhost:6060/pkg/app/internal?m=all' | sed 's/\/lib\/godoc/\/mono-management\/backend/g' > /gh-pages/dist/backend/internal/index.html &&
    godoc -url 'http://localhost:6060/pkg/app/router?m=all' | sed 's/\/lib\/godoc/\/mono-management\/backend/g' > /gh-pages/dist/backend/router/index.html
    "

cd "$GITROOTPATH/gh-pages"
npm run deploy
