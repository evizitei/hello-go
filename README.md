Hello, Go

Base Image is here: https://registry.hub.docker.com/_/golang/

Right now, this is just a hello world file in a docker container.

To build the image:

```bash
$> docker build -t my-golang-tag
```

Once the image exists, run the "Hello, world" executable with:

```bash
$> docker run -it --rm --name hello-go my-golang-tag
```

You should see "hello world" printed out before the container exits
