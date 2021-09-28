package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ASVBPREAUBV/orthanc-drop/cmd"
	"github.com/imroc/req"
	"github.com/spf13/viper"
)

type orthancConfig struct {
	Dir      string
	User     string
	Password string
	Url      string
}

func getAuthHeader(user, pass string) string {
	auth := user + ":" + pass
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
	readConfig()
	cmd.Execute()
}

func readConfig() {

	viper.SetConfigName("odrop")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("dir", ".")
	viper.SetDefault("user", "user")
	viper.SetDefault("password", "pass")
	viper.SetDefault("url", "http://localhost:8080/instances")

	viper.ReadInConfig()

	config := orthancConfig{}
	viper.Unmarshal(&config)

	fmt.Println(config)

	err := viper.ReadInConfig()
	if err != nil {
		os.Create("./odrop.yml")
		viper.WriteConfig()
		fmt.Println("-------------------------")
		fmt.Println("no config file available, i created one for you. (./odrop.yml)")
		fmt.Println("-------------------------")
	} else {

		fmt.Println("----used config file:----")
		fmt.Println(viper.ConfigFileUsed())
		fmt.Println("-------------------------")
		upload(config)

	}

}

func upload(config orthancConfig) {
	dirToUpload := config.Dir
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
		info, _ := os.Stat(file)
		if !info.IsDir() {
			header := req.Header{
				"Accept":        "application/json",
				"Authorization": getAuthHeader(config.User, config.Password),
				"Content-Type":  "application/dicom",
			}

			file, _ := os.Open(file)
			fmt.Println("send file:", file)
			res, err := req.Post(config.Url, req.FileUpload{
				File: file,
			}, header)

			fmt.Print(res, err)
		}

	}

}
