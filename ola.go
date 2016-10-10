package main

import (
	"log"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

var (
	dockerEndpoint = "unix:///var/run/docker.sock"
)

func main() {
	networkName := os.Getenv("ADD_TO_NETWORK")
	if networkName == "" {
		log.Fatal("No network name?")
	}

	client, err := docker.NewClient(dockerEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	addContainerNetwork := func(cid string) {
		log.Printf("Checking if container %s should be added to the %s network\n", cid, networkName)
		container, err := client.InspectContainer(cid)
		if err != nil {
			log.Println(err)
			return
		}

		if container.Config.Labels["com.docker.swarm.service.id"] == "" {
			log.Printf("IGNORING: container %s is not a service\n", cid)
			return
		}

		for name, _ := range container.NetworkSettings.Networks {
			if name == networkName {
				log.Printf("IGNORING: container %s is already a member of %s\n", cid, networkName)
				return
			}
		}

		err = client.ConnectNetwork(networkName, docker.NetworkConnectionOptions{Container: cid})
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Added container %s to network %s\n", cid, networkName)
	}

	events := make(chan *docker.APIEvents)
	err = client.AddEventListener(events)
	if err != nil {
		log.Fatal(err)
	}

	containers, _ := client.ListContainers(docker.ListContainersOptions{})
	for _, c := range containers {
		go addContainerNetwork(c.ID)
	}

	for {
		select {
		case event, ok := <-events:
			if !ok {
				log.Println("Events closed")
				os.Exit(2)
			}

			log.Printf("Got event %s for container %s", event.Action, event.Actor.ID)
			if event.Action == "start" {
				go addContainerNetwork(event.Actor.ID)
			}
		}
	}
}
