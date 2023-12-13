package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	dockerFilesPaths := os.Args[1:]
	if len(dockerFilesPaths) < 1 {
		log.Fatal("You must supply at least one docker compose file and in the format docker-compose_project.yml")
		return
	}

	projectNames := make([]string, 0)
	for _, path := range dockerFilesPaths {
		projectNames = append(projectNames, getProjectName(path))
	}

	resultCompose := DockerCompose{}
	for i, path := range dockerFilesPaths {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		p := DockerCompose{}
		err = yaml.Unmarshal(content, &p)
		if err != nil {
			log.Fatal(err)
		}

		for serviceKey, serviceValue := range p.Services {
			if resultCompose.Services == nil {
				resultCompose.Services = make(map[string]Service)
			}

			for networkKey, networkValue := range serviceValue.Networks {
				if !isProjectPrefixPresent(projectNames, networkKey) {
					serviceValue.Networks[projectNames[i]+"_"+networkKey] = networkValue
					delete(serviceValue.Networks, networkKey)
				}
			}
			resultCompose.Services[serviceKey] = serviceValue
		}

		if p.Networks != nil {
			for networkKey, networkValue := range p.Networks {
				if resultCompose.Networks == nil {
					resultCompose.Networks = make(map[string]Network)
				}

				networkValue.Name = projectNames[i]+"_"+networkKey
				if !p.Networks[networkKey].External {
					resultCompose.Networks[projectNames[i]+"_"+networkKey] = networkValue
				}
			}
		}

		if p.Volumes != nil {
			for volumeKey, volumeValue := range p.Volumes {
				if resultCompose.Volumes == nil {
					resultCompose.Volumes = make(map[string]Volume)
				}
				resultCompose.Volumes[volumeKey] = volumeValue
			}
		}

		if p.Secrets != nil {
			for secretKey, secretValue := range p.Secrets {
				if resultCompose.Secrets == nil {
					resultCompose.Secrets = make(map[string]Secret)
				}
				resultCompose.Secrets[secretKey] = secretValue
			}
		}

		if p.Configs != nil {
			for configKey, configValue := range p.Configs {
				if resultCompose.Configs == nil {
					resultCompose.Configs = make(map[string]Config)
				}
				resultCompose.Configs[configKey] = configValue
			}
		}
	}

	fmt.Println(resultCompose)

	resultComposeFile, err := os.OpenFile("docker-compose_full.yml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Error opening/creating file: %v", err)
	}
	defer resultComposeFile.Close()

	enc := yaml.NewEncoder(resultComposeFile)
	err = enc.Encode(resultCompose)
	if err != nil {
		log.Fatal(err)
	}
}

func isProjectPrefixPresent(projectNames []string, networkKey string) bool {
	for _, projectName := range projectNames {
		if strings.HasPrefix(networkKey, projectName+"_") {
			return true
		}
	}
	return false
}

func getProjectName(filename string) string {
	parts := strings.Split(filename, "_")
	return strings.Replace(parts[len(parts)-1], ".yml", "", -1)
}
