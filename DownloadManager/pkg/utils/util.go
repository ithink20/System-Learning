package utils

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path"
)

func GetURLFromUser() (*url.URL, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the file URL to download: ")
	scanner.Scan()
	input := scanner.Text()

	//convert to URL
	userURL, err := url.Parse(input)
	if err != nil {
		return nil, err
	}
	return userURL, nil
}

func ExtractFileName(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	fileName := path.Base(parsedURL.Path)
	if fileName == "/" || fileName == "." {
		return "", fmt.Errorf("unable to extract file name from URL: %s", urlStr)
	}

	return fileName, nil
}
