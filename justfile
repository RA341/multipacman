set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

dkbd:
    docker build . -t ras334/gouda:dev

prune:
    docker image prune -f

devp:
    just dkbd
    docker login
    docker push ras334/gouda:dev

# no cache build
dkc:
    docker build . -t ras334/gouda:local --no-cache

[working-directory: 'frontend']
ui:
    flutter build web
    rm -r ../core/web/
    cp -r build/web ../core/web
