package main

import (
	"strings"
	"os/exec"
	"regexp"
	"log"
	"flag"

)

/**
 * Sequence of things to happen at the openshift cli (oc):
 * - oc login -u xyz -p xyz
 * -
 */

var queryName string
var appName string
var userName string
var password string

type QueryArgs struct {
	queryName string
	appName string
	userName string
	password string
}


func constructArgs(args [] string) QueryArgs {
	queryArgs := new(QueryArgs)
	queryArgs.queryName = args[0]
	queryCommand := flag.NewFlagSet(queryName, flag.ExitOnError)

	appNamePtr := queryCommand.String("appname", "pet-clinic", "name of openshift application")
	userNamePtr := queryCommand.String("username", "dev", "user name")
	passwordPtr := queryCommand.String("password", "dev", "password")

	queryCommand.Parse(args[1:])

	queryArgs.appName = *appNamePtr
	queryArgs.userName = *userNamePtr
	queryArgs.password = *passwordPtr

	return *queryArgs
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


