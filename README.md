# echo-server

An HTTP echo server designed for testing applications and proxies.

![workflow dockerhub](https://github.com/erikcc02/echo-server/actions/workflows/built-to-dockerhub.yaml/badge.svg)



## Running

### Docker

simple command:

```bash
docker run erikcc02/echo-server
```

mapping port:

```bash
docker run -p 8555:8000 erikcc02/echo-server
```

## Configure echo server

to set the port that the application will run, set the value of the variable `ECHOSERVER_PORT`

```bash
docker run -it -p 8555:5000 -e ECHOSERVER_PORT=5000 echo-server
```
