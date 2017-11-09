**NOTE:** This project has been deprecated. It is outdated, and no longer in
use in any of our projects.

---

Docker Deploy Hook
==================

The _Docker Deploy Hook_ project is a simple docker image that runs a small go
server that listens on a specified, secret, endpoint and pulls and re-deploys
another docker container.

We use this at [Blendle](https://blendle.com) to auto deploy our docker service
when our CI calls this hook.

- - -

### Example

```bash
docker run -P -v /var/run/docker.sock:/var/run/docker.sock -e ENDPOINT=/secret -e USERNAME=hubuser -e PASSWORD=hubpass -e "CMD=sleep 500" blendle/ddh
```

- - -

### Environment (options)

The go script uses these environment variables as config:

| ENV            | Default                     | Description                                                                               |
| -------------- | --------------------------- | ----------------------------------------------------------------------------------------- |
| PASSWORD       |                             | Your dockerhub username when pulling an image                                             |
| USERNAME       |                             | Your dockerhub password when pulling an image                                             |
| DOCKERSOCKET   | unix:///var/run/docker.sock | The URI used to connect to the Docker host                                                |
| ENDPOINT       | /                           | The __secret__ endpoint that triggers the deploy (CHANGE THIS!)                           |
| IMAGE          | ubuntu                      | The image name to pull and re-create                                                      |
| CONTAINER_NAME | ubuntu-test                 | The name of the instance container to kill and start again                                |
| CMD            |                             | The command to run inside the deployed container, can be empty                            |
| PASS_ENV       | PASS_ENV                    | A space seperated list of environment variables to pass on to the newly created container |
| LINKS          |                             | Links for the docker container, format is "container:alias" seperated by space            |
