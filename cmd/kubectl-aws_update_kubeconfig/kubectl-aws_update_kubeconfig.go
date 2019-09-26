package main

import (
	"fmt"
	"github.com/FrankFuenmayor/kubectl-kubeconfig-plugins/pkg/aws"
	"github.com/FrankFuenmayor/kubectl-kubeconfig-plugins/pkg/emoji"
	"github.com/gookit/color"
	"os"
	"strings"
	"time"
)

func main() {

	eks := aws.Eks{Args: os.Args[1:]}

	printf := color.Bold.Printf
	println := color.Bold.Println

	printf("%v Fetching clusters list...\n", emoji.Emoji("1F50E"))
	clusters, err := eks.ListClusters()

	if err != nil {
		panic(err)
	}

	if len(clusters) == 0 {
		println("No cluster(s) found")
		os.Exit(1)
	}

	clusterList := strings.Join(clusters, "\n - ")
	printf("%v Add or update: \n - %v \n", emoji.Emoji("1F4DD"), clusterList)

	stopClock := make(chan struct{}, 1)

	go startClock(" Updating kubeconfig file...", stopClock)

	err = eks.UpdateKubeconfig(clusters)

	close(stopClock)

	if err != nil {
		panic(err)
	}
	println("✅ success.")
}

func startClock(msg string, stop <-chan struct{}) {
	var clocksSequence = [...]emoji.Emoji{"231B", "23F3"}
	for t := range time.Tick(500 * time.Millisecond) {
		select {
		case <-stop:
			break
		default:
			fmt.Print(clocksSequence[t.Second()%len(clocksSequence)])
			color.Bold.Print(msg)
			fmt.Print("\r")
		}
	}
}
