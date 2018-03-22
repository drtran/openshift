package main

import (
	"fmt"
	"os"
	"time"
	"strings"
	"os/exec"
	"regexp"
	"log"
)

/**
 * Sequence of things to happen at the openshift cli (oc):
 * - oc login -u xyz -p xyz
 * -
 */
var queryName string
var appName string

func verifyArguments(osArgs []string) bool {
	if len(osArgs) == 0 {
		fmt.Println("Run as: openshift_query queryName appName")
		fmt.Println(" where: queryName = {waitForBuild}")
		fmt.Println("        appName = name of openshift application")
		return false
	}
	queryName = osArgs[0]
	appName = osArgs[1]
	return true
}

func main() {
	if !verifyArguments(os.Args[1:]) {
		return
	}

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
	return execOC("describe", "build", appName)
}

func execOC(action string, objectType string, objectName string) [] string {
	cmdOut, _ := exec.Command("oc", action, objectType, objectName).Output()
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
