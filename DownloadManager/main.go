package main

import (
	"fmt"
	"golang/pkg/worker-pool/DownloadManager/pkg/manager"
	"golang/pkg/worker-pool/DownloadManager/pkg/utils"
	"log"
)

func main() {
	fmt.Println("...Starting Manager...")

	userURL, err := utils.GetURLFromUser()
	if err != nil {
		log.Fatal("Invalid URL : ", err)
	}
	manager.StartFast(userURL)
}
