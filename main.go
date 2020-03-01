package main

import (
	"encoding/base64"
	"fmt"
	"github.com/ASVBPREAUBV/orthanc-drop/cmd"
	"github.com/spf13/viper"
)

type orthancConfig struct {
	dir      string
	user     string
	password string
	url      string
}

func getAuthheader(user, pass string) string {
	auth := user + ":" + pass
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
	readConfig()
	cmd.Execute()
	//upload()
}

func readConfig() {
	viper.SetDefault("dir", ".")
	viper.SetDefault("user", "")
	viper.SetDefault("password", "")
	viper.SetDefault("url", "http://localhost:8080/")
	fmt.Println("----used config file:----")
	fmt.Println(viper.ConfigFileUsed())
	fmt.Println("-------------------------")
	config := orthancConfig{}
	viper.Unmarshal(config)

	viper.SetConfigName("orthanc-drop") // name of config file (without extension)
	viper.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")            // optionally look for config in the working directory
	err := viper.ReadInConfig()         // Find and read the config file
	if err != nil { // Handle errors reading the config file
		viper.SafeWriteConfig()
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

/*func upload() {
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

}*/
