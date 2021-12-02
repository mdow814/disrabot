package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Build struct {
	Items []string `yaml:"items"`
}

type BuildData struct {
	Builds []Build `yaml:"builds"`
}

func formatBuild(build Build) string {
	var formattedBuild string
	for _, j := range build.Items {
		formattedBuild += j + "\n"
	}
	return formattedBuild
}

func LoadBuilds() BuildData {
	builds := BuildData{}
	buildsFile, err := ioutil.ReadFile("utils/builds.yaml")
	if err != nil {
		fmt.Println("Error reading file")
	}
	err = yaml.Unmarshal(buildsFile, &builds)
	if err != nil {
		fmt.Println("Couldnt unmarshal builds file")
	}
	return builds
}

func (b *BuildData) GetBuild(randomNum int) string {
	build := b.Builds[randomNum]
	return fmt.Sprintf("Here's your build!\n\n%s", formatBuild(build))
}
