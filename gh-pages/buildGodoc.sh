#!/bin/bash
cd /go/src/app
mv ./vendor /tmp/ 2>/dev/null
rm -rf /gh-pages/dist/backend
mkdir -p /gh-pages/dist/backend/internal
mkdir -p /gh-pages/dist/backend/router
go get golang.org/x/tools/cmd/godoc
cp /go/src/golang.org/x/tools/godoc/static/jquery.js /gh-pages/dist/backend/
cp /go/src/golang.org/x/tools/godoc/static/godocs.js /gh-pages/dist/backend/
cp /go/src/golang.org/x/tools/godoc/static/style.css /gh-pages/dist/backend/
godoc -url 'http://localhost:6060/pkg/app?m=all' | sed 's/\/lib\/godoc/\/mono-management\/backend/g' > /gh-pages/dist/backend/index.html
godoc -url 'http://localhost:6060/pkg/app/internal?m=all' | sed 's/\/lib\/godoc/\/mono-management\/backend/g' > /gh-pages/dist/backend/internal/index.html
godoc -url 'http://localhost:6060/pkg/app/router?m=all' | sed 's/\/lib\/godoc/\/mono-management\/backend/g' > /gh-pages/dist/backend/router/index.html
mv /tmp/vendor ./ 2>/dev/null
