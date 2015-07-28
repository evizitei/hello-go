Hello, Go

Use Dinghy! (https://github.com/codekitchen/dinghy)

Assuming from this point forward that dinghy is up and running.

Base Image is here: https://registry.hub.docker.com/_/golang/

Right now, this is just a hello world webapp in a docker container.

To build the image:

```bash
$> docker build -t hello-go
```

Once the image exists, startup the server with:

```bash
$> docker run --publish 6060:8080 --name test-go --rm hello-go
```

Find your address with

```bash
$> dinghy ip
```

You should be able to see json output containing "hello world!" returned in the browser by
visiting "http://dinghy-ip:6060".
