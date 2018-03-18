package main

import (
	"fmt" 
	"os/exec"
	"strings"
	"regexp"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Run as: openshift_query queryName appName")
		fmt.Println(" where: queryName = {waitForBuild}")
		fmt.Println("        appName = name of openshift application")
		return
	}

	queryName := os.Args[1]
	appName := os.Args[2]

	if "waitForBuild" == queryName {
		for {
			status := getBuildStatus(appName)
			fmt.Println(status)
			if "Complete" == status {
				fmt.Println("Build completes.")
				break
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func getBuildStatus(appName string) string {
	lines := getBuildDescription(appName)
	status := ""

	for _, line := range lines {
		if strings.HasPrefix(line, "Status:") {
			status = getAlphabetOnlyString(line[7:len(line)])
			status = status[2:len(status)-2]
		}
	}
	return status
}

func getBuildDescription(appName string) []string {
	cmdOut, _ := exec.Command("oc", "describe", "build", appName).Output()
	outputString := string(cmdOut)
	return strings.Split(outputString, "\n")
}

func getAlphabetOnlyString(src string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
    if err != nil {
        log.Fatal(err)
    }
    return reg.ReplaceAllString(src, "")
}