set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

dkbd:
    docker build . -t ras334/multipacman:dev

dkp:
    just dkbd
    docker login
    docker push  ras334/multipacman:dev

prune:
    docker image prune -f

[working-directory: 'frontend']
ui:
    flutter build web
    rm -r ../core/web/
    cp -r build/web ../core/web
