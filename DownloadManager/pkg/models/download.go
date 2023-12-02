package models

import (
	"fmt"
	"golang/pkg/worker-pool/DownloadManager/pkg/http"
	"io"
	"os"
)

type Download struct {
	URL        string
	FileName   string
	Chunks     int
	ChunkSize  int64
	TotalSize  int64
	HttpClient *http.HttpClient
}

// SplitIntoChunks splits total size into n chunks and respective start and end ranges
func (d *Download) SplitIntoChunks() [][2]int64 {
	arr := make([][2]int64, d.Chunks)
	for i := 0; i < d.Chunks; i++ {
		if i == 0 {
			arr[i][0] = 0
			arr[i][1] = d.ChunkSize
		} else if i == d.Chunks-1 {
			arr[i][0] = arr[i-1][1] + 1
			arr[i][1] = d.TotalSize - 1
		} else {
			arr[i][0] = arr[i-1][1] + 1
			arr[i][1] = arr[i][0] + d.ChunkSize
		}
	}
	return arr
}

func (d *Download) Download(index int, byteChunk [2]int64) error {
	//fmt.Printf("Downloading Chunk - %v [%d, %d]\n", index, byteChunk[0], byteChunk[1])
	// make GET request with range
	headers := map[string]string{
		"User-Agent": "Downloader",
		"Range":      fmt.Sprintf("bytes=%v-%v", byteChunk[0], byteChunk[1]),
	}
	resp, err := d.HttpClient.ExecuteCustomRequest("GET", d.URL, headers)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Chunk fail %v", resp.StatusCode))
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("can't process, server response is %v", resp.StatusCode)
	}
	// open a file to write the body to
	fname := fmt.Sprintf("tempFile-%v.tmp", index)
	file, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Can't create a file %v", fname))
	}
	defer file.Close()

	// write to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	//log.Println(fmt.Sprintf("wrote chunk %v to file", index))
	return nil
}

func (d *Download) MergeDownloads() error {
	// create output file
	out, err := os.Create(d.FileName)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer out.Close()

	// append each chunk to final file
	for idx := 0; idx < d.Chunks; idx++ {
		fname := fmt.Sprintf("tempFile-%v.tmp", idx)
		in, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("failed to open chunk file %s: %v", fname, err)
		}
		defer in.Close() //Todo: Fix this

		_, err = io.Copy(out, in)
		if err != nil {
			return fmt.Errorf("failed to merge chunk file %s: %v", fname, err)
		}
	}

	//fmt.Println("File chunks merged successfully...")
	return nil
}

func (d *Download) CleanupTmpFiles() error {
	//log.Println("Starting to clean tmp downloaded files...")

	// delete each chunk file
	for idx := 0; idx < d.Chunks; idx++ {
		fname := fmt.Sprintf("tempFile-%v.tmp", idx)
		err := os.Remove(fname)
		if err != nil {
			return fmt.Errorf("failed to remove chunk file %s: %v", fname, err)
		}
	}

	return nil
}
