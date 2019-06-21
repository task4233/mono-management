#!/bin/bash

GITROOTPATH="$(git rev-parse --show-toplevel)"
docker run -t --rm -v "$GITROOTPATH/gh-pages":/gh-pages -v "$GITROOTPATH/backend":/go/src/app golang:1.12.5 /bin/bash -c "/gh-pages/buildGodoc.sh"

cd "$GITROOTPATH/gh-pages"
npm install
npm run deploy
