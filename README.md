## 💻 RESTFull *Echo** API experiments

Simple go API suing [echo](https://echo.labstack.com/guide/) web framework written in golang

## 🚀 Techs

- [Linux maybe](https://ubuntu.com/)
- [echo](https://echo.labstack.com/guide/)
- [Docker](https://docs.docker.com/get-docker)
- [Golang](https://go.dev/)
- [Task](https://taskfile.dev/)

## 🧘🏿‍ Features

- [x] Get header content
- [x] dockerize app

## Enjoy it by using Taskfile

```bash
    # build app
    task build 

    # build docker image
    task build-image

    # start a container with previous build image
    task deploy
```

> Please refer to [**Taskfile**](Taskfile.yaml) for more dev commands
