package main

import (
	"os"
	"fmt"
	"time"
)

func main() {

	if len(os.Args) == 1 {
		args := []string{"-h"}
		constructArgs(args)
		os.Exit(-1)
	}
	queryArgs := constructArgs(os.Args[1:])

	if "waitForBuild" == queryArgs.queryName {
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

func init() {

}