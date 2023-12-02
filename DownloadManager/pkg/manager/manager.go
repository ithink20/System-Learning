package manager

import (
	"fmt"
	"golang/pkg/worker-pool/DownloadManager/pkg/models"
	"golang/pkg/worker-pool/DownloadManager/pkg/utils"
	"log"
	"net/url"
	"sync"

	"golang/pkg/worker-pool/DownloadManager/pkg/http"
)

const workers = 10

func StartFast(urlPointer *url.URL) {
	client := http.NewHttpClient()

	urlString := urlPointer.String()
	headers := map[string]string{
		"User-Agent": "Download-Manager",
	}
	resp, err := client.ExecuteCustomRequest("GET", urlString, headers)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Response: %v, length: %d", resp, resp.ContentLength)
	contentLength := resp.ContentLength
	chunkSize := contentLength / workers
	//fmt.Printf("Total Size: %d, Each chunk size : %d\n", contentLength, chunkSize)

	// get file name
	fname, err := utils.ExtractFileName(urlString)
	if err != nil {
		log.Fatal("Error extracting filename...")
	}
	//log.Println("Filename extracted: ", fname)

	downloadReq := models.Download{
		URL:        urlString,
		FileName:   fname,
		Chunks:     workers,
		ChunkSize:  chunkSize,
		TotalSize:  contentLength,
		HttpClient: client,
	}

	byteRangeArray := downloadReq.SplitIntoChunks()
	//fmt.Println(byteRangeArray)

	var wg sync.WaitGroup
	for index, byteChunk := range byteRangeArray {
		wg.Add(1)
		go func(index int, byteChunk [2]int64) {
			err := downloadReq.Download(index, byteChunk)
			if err != nil {
				log.Fatal(fmt.Sprintf("Failed to download chunk %v", index), err)
			}
			wg.Done()
		}(index, byteChunk)
	}
	wg.Wait()

	// merge
	err = downloadReq.MergeDownloads()
	if err != nil {
		log.Fatal("Failed merging tmp downloaded files...", err)
	}

	// cleanup
	err = downloadReq.CleanupTmpFiles()
	if err != nil {
		log.Fatal("Failed cleaning up tmp downloaded files...", err)
	}

	// final file generated
	//log.Println(fmt.Sprintf("File generated: %v\n\n", downloadReq.FileName))
}

// StartSlow Experiment
func StartSlow(urlPointer *url.URL) {
	client := http.NewHttpClient()

	urlString := urlPointer.String()
	headers := map[string]string{
		"User-Agent": "Download-Manager",
	}
	resp, err := client.ExecuteCustomRequest("GET", urlString, headers)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Response: %v, length: %d", resp, resp.ContentLength)
	contentLength := resp.ContentLength
	chunkSize := contentLength / workers
	//fmt.Printf("Total Size: %d, Each chunk size : %d\n", contentLength, chunkSize)

	// get file name
	fname, err := utils.ExtractFileName(urlString)
	if err != nil {
		log.Fatal("Error extracting filename...")
	}
	//log.Println("Filename extracted: ", fname)

	downloadReq := models.Download{
		URL:        urlString,
		FileName:   fname,
		Chunks:     1,
		ChunkSize:  chunkSize,
		TotalSize:  contentLength,
		HttpClient: client,
	}

	err = downloadReq.Download(0, [2]int64{0, contentLength - 1})
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to download chunk %v", 0), err)
	}

	// merge
	err = downloadReq.MergeDownloads()
	if err != nil {
		log.Fatal("Failed merging tmp downloaded files...", err)
	}

	// cleanup
	err = downloadReq.CleanupTmpFiles()
	if err != nil {
		log.Fatal("Failed cleaning up tmp downloaded files...", err)
	}

	// final file generated
	//log.Println(fmt.Sprintf("File generated: %v\n\n", downloadReq.FileName))
}
