ola
===

# Ola?

Found [this](http://www.behindthename.com/random/random.php?number=1&gender=both&surname=&all=no&usage_nor=1).

But this project should be dead [fixed] pretty [soon](https://github.com/docker/docker/issues/27082).

## Okay ...

This thing listens to Docker events for new container starts and adds those containers to a local network.  It must run locally on each Docker node.  This is probably only relevant for you if you are using SwarmKit and services.

### Ever got this error?

```
Error response from daemon: network bridge is not eligible for docker services
```

### But why?

Stuff running in containers started as Docker services (which aren't publically exposed) are a bit hard to get at from anything that is not also a service.  At least I didn't figure it out yet.

But one easy-way solution seems to be just adding the container to the local bridge network and using that address.


# Super

Just start this thing on each node in your Docker cluster and hope for the best.  Specify the local network name via `ADD_TO_NETWORK` environment variable.

```
$ docker run -d --name ola_ola --env ADD_TO_NETWORK=bridge --volume /var/run/docker.sock:/var/run/docker.sock symfoni/ola:latest
```

or if you are feeling brave;

```
docker service create --name ola_ola --mode global --env ADD_TO_NETWORK=bridge --mount type=bind,source=/var/run/docker.sock,target=/var/run/docker.sock symfoni/ola:latest
```
