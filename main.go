package main

import (
	"encoding/base64"
	"fmt"
	"github.com/imroc/req"
	"os"
	"path/filepath"
)

var dirToUpload = "./test"
var user = "local"
var pass = "locallocal"
var hostUrl = "http://localhost:8888/instances"

func getAuthheader() string {
	auth := user + ":" + pass
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
	execute()
	upload()
}

func upload() {
	fmt.Print("DIR: ", dirToUpload)
	if _, err := os.Stat(dirToUpload); os.IsNotExist(err) {
		panic("DIR does not exist : " + dirToUpload)
	}
	var files []string

	err := filepath.Walk(dirToUpload, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		//if filepath.Ext(file) == ".dcm" {
		header := req.Header{
			"Accept":        "application/json",
			"Authorization": getAuthheader(),
			"Content-Type":  "application/dicom",
		}

		file, _ := os.Open(file)

		res, err := req.Post(hostUrl, req.FileUpload{
			File: file,
		}, header)

		fmt.Print(res, err)
		//}
	}

}
