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
	if "waitForBuildToComplete" == os.Args[1] {
		for {
			status := getBuildStatus("pet-clinic")
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
	cmdOut, _ := exec.Command("oc", "describe", "build", appName).Output()
	outputString := string(cmdOut)
	lines := strings.Split(outputString, "\n")
	status := ""

	for _, line := range lines {
		if strings.HasPrefix(line, "Status:") {
			status = getAlphabetOnlyString(line[7:len(line)])
			status = status[2:len(status)-2]
		}
	}
	return status
}

func getAlphabetOnlyString(src string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
    if err != nil {
        log.Fatal(err)
    }
    return reg.ReplaceAllString(src, "")
}