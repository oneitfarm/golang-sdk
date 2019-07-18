package cienv

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

const EnvsBasePath = ""

var envCache = map[string]string{}

var lock = sync.Mutex{}

func GetEnv(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return ""
	}

	lock.Lock()
	defer lock.Unlock()

	value, exits := envCache[name]
	if exits {
		return value
	}

	value, exits = getEnvFromFile(name)
	if exits {
		envCache[name] = value
		return value
	}

	return os.Getenv(name)
}

func getEnvFromFile(name string) (string, bool) {
	filePath := fmt.Sprintf("%s/%s", EnvsBasePath, name)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", false
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", false
	}

	return string(data), true
}

func GetMysqlConfig(name string) *MysqlConfig {
	host := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_HOST_WRITE", name))
	port := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_PORT", name))
	intPort, _ := strconv.Atoi(port)
	db := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_NAME", name))
	user := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_USERNAME", name))
	password := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_PASSWORD", name))
	return &MysqlConfig{
		Host:     host,
		Port:     intPort,
		DB:       db,
		User:     user,
		Password: password,
	}
}
