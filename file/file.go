package file

import (
	"github.com/magiconair/properties"
	"log"
	"os"
	"strings"
)

type Config struct {
	LocalIp   string `properties:"tunnel.ip,default=127.0.0.1"`
	LocalPort int    `properties:"tunnel.port,default=8080"`
	RemoteIp  string `properties:"remote.ip,default=106.12.160.188"`
}

var cfg Config

func init() {
	filePath := os.Getenv("config.file.path")
	if strings.Compare(filePath, "") == 0 {
		filePath = "config.properties"
	} else {
		if !strings.HasSuffix(filePath, "properties") {
			suffix := "config.properties"
			if !strings.HasSuffix(filePath, "\\/") {
				suffix = "\\/" + suffix
			}
			filePath += suffix
		}
	}
	config, loadErr := properties.LoadFile(filePath, properties.UTF8)
	if loadErr != nil {
		log.Println(" The system cannot find the file specified")
	} else {
		if err := config.Decode(&cfg); err != nil {
			log.Fatal(err)
		}
	}
}
func GetEnvParam() Config {
	return cfg
}
