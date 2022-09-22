/**
Author: Isaac Irani <isri@boozt.com>
Modified: 22-09-2022
**/

package main

import (
	"log"
	"os"
	"time"
)

const dockerfileTemplate = "Dockerfile"

var projects = [3]string{"project1", "project2", "project3"}

func main() {

	for _, element := range projects {
		dockerfilename := generateDockerfile(element)

		println("Succesfully generated dockerfile: " + dockerfilename)
	}

}

func generateDockerfile(name string) string {
	currentTime := time.Now()

	//
	baseDockerfile, err1 := os.ReadFile(dockerfileTemplate)

	if err1 != nil {
		log.Fatalf("failed opening file: %s", err1)
	}

	dockerfileAddition := []byte("#added this comment on " + currentTime.Format("2006-01-02 15:04:05") + " for project " + name + "\n")

	dockerfileContent := append(dockerfileAddition, baseDockerfile...)

	dockerfile, err := os.Create("generated/dockerfile_" + name)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer dockerfile.Close()

	_, err2 := dockerfile.Write(dockerfileContent)

	if err2 != nil {
		log.Fatalf("failed to write to file: %s", err1)
	}

	return dockerfile.Name()
}
