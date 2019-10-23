package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)


type Config struct {
	configFilePath                string
	DEBUG                         bool
	VERSION                       string
	APPNAME                       string
	LOGDIR                        string
	PORT                          string
}

var conf Config

func init(){
	defer func() { r := recover() ; if r != nil {	log.Print("Possible .env error:" , r)} }()
	flag.StringVar(&conf.configFilePath, "config", ".env", "config file path") //Get the .env file for settings
	flag.Parse()
	log.Println("Config File:",conf.configFilePath)
	checkErr(godotenv.Load(conf.configFilePath))
	conf = Config{
				conf.configFilePath,
				getEnvAsBool("DEBUG",true),
				getEnv("VERSION","<!--version-->") ,
				getEnv("APPNAME","AENConnect-API"),
				getEnv("LOGDIR","/var/log/connect-api"),
				getEnv("PORT","8000"),
		}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsInt64(name string, defaultVal int64) int64 {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseInt(valueStr,10,64); err == nil {
		return value
	}
	return defaultVal
}
