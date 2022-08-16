# go_chat

## Client
for getting server
```
go get gitlab.com/wushyrussia/go_chat/client
```
Starting using default url 'localhost:8085'
```
go run gitlab.com/wushyrussia/go_chat/client
```

Starting with custom url 'some_url_or_ip:8085'
```
go run gitlab.com/wushyrussia/go_chat/client -server="some_url_or_ip:server_port"
```

## Server
- Default port '8085'
- Default chat rooms "General chat",  "Next gen chat",  "Programmers chat"

for getting server
```
go get gitlab.com/wushyrussia/go_chat/server
```

Starting using port default 8085
```
go run gitlab.com/wushyrussia/go_chat/server
```
also, you can start using custom port:
```
go run gitlab.com/wushyrussia/go_chat/server -port=8080
```

also you could add chat rooms
```
go run gitlab.com/wushyrussia/go_chat/server -rooms="room 1,room 2,room 3"
```

### using with docker:
build server docker img:
```
cd $project_folder/server/Deployment
docker build -t go_server_image .
```
run server img
```
docker run go_server_image
```
stop server img
```
docker container ls
docker stop <containerId>
```

### using with Monk:
```
cd $project_folder/server/Deployment/monk
monk load app.yaml
monk run wushyrussia/app
```