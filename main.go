package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"path"

	"github.com/JacksonGariety/new-left-news/app/utils"
)

var Env = os.Getenv("env")
var Port = os.Getenv("port")

func main() {
	config := make(map[string]map[string]string)
	data, _ := ioutil.ReadFile(path.Join(utils.BasePath, "db/dbconf.yml"))
	_ = yaml.Unmarshal([]byte(data), &config)

	utils.InitTemplates()
	utils.InitDB(config[Env]["open"])

	log.Println("Leftisting...")
	log.Fatal(http.ListenAndServe(":" + Port, NewRouter()))
}
