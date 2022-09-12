# Build docker image

## BUILD
```bash
docker build . -t <IMAGE_NAME>
```
## RUN
```bash
docker run --name <NAME> -d \
  --restart=always --net=host \
  -p <HTTP_PORT>:<HTTP_PORT> \
  -p <GRPC_PORT>:<GRPC_PORT> \
   <IMAGE_NAME> \
  /bin/sh -c './binary/ <NAME> run' && \
docker logs -f  <NAME>
```


## EXAMPLE
### Account
BUILD
```bash
docker rmi account && \
docker build . -t account
```
RUN
```bash
docker run --name account -d \
  --restart=always --net=host\
  -p 7010:7010 \
  -p 50010:50010 \
  account \
  /bin/sh -c './binary/account run' && \
docker logs -f account
```


### Public
BUILD
```bash
docker build . -t public
```
RUN
```bash
docker run --name public -d \
  --restart=always --net=host \
  hvx \
  /bin/sh -c './binary/public run' && \
docker logs -f public
```


### ALL APP ONE IMAGE
```bash
docker-compose down && \
docker rmi hvx && \
docker build . -t hvx
```
