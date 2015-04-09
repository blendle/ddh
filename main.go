package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsouza/go-dockerclient"
	"github.com/go-martini/martini"
)

func main() {
	endpoint := os.Getenv("ENDPOINT")
	imageName := os.Getenv("IMAGE")
	containerName := os.Getenv("CONTAINER_NAME")

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	if len(username) == 0 || len(password) == 0 {
		fmt.Println("missing USERNAME/PASSWORD env!")
		os.Exit(1)
	}

	dockerURI := os.Getenv("DOCKERSOCKET")
	client, _ := docker.NewClient(dockerURI)

	err := client.Ping()
	if err != nil {
		fmt.Println("unable to connect to docker:", err)
		fmt.Println("(did you use `docker run -v /var/run/docker.sock:/var/run/docker.sock ...`?)")
		os.Exit(1)
	}

	image := docker.PullImageOptions{Repository: imageName, Tag: "latest"}
	auth := docker.AuthConfiguration{Username: username, Password: password}

	passes := strings.Split(os.Getenv("PASS_ENV"), " ")
	envs := make([]string, len(passes))
	for i, env := range passes {
		envs[i] = env + "=" + os.Getenv(env)
	}

	config := docker.Config{Image: imageName, Env: envs}
	if len(os.Getenv("CMD")) != 0 {
		// TODO: Parse this CMD string as bashlike:
		config.Cmd = strings.Split(os.Getenv("CMD"), " ")
	}
	create := docker.CreateContainerOptions{Name: containerName, Config: &config}
	hostConfig := docker.HostConfig{PublishAllPorts: true}

	links := os.Getenv("LINKS")
	if len(links) != 0 {
		hostConfig.Links = strings.Split(links, " ")
	}

	deploy := make(chan int, 100)

	go func() {
		for {
			_ = <-deploy
			fmt.Println("Pulling image:", imageName)
			client.PullImage(image, auth)
			fmt.Println("Removing old container:", containerName)
			client.RemoveContainer(docker.RemoveContainerOptions{ID: containerName, Force: true})
			fmt.Println("Creating new container:", containerName)
			container, _ := client.CreateContainer(create)
			fmt.Println("Starting container:", container.ID)
			client.StartContainer(container.ID, &hostConfig)
		}
	}()
	m := martini.Classic()
	m.Get(endpoint, func() string {
		deploy <- 0
		return "OK\n"
	})
	m.RunOnAddr(":8080")
}
