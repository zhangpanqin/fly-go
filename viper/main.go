package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	conf := viper.New()
	conf.AutomaticEnv()
	conf.SetConfigFile("./viper/config/local.yml")
	replacer := strings.NewReplacer(".", "_")
	conf.SetEnvKeyReplacer(replacer)
	err := conf.ReadInConfig()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Printf("tokens: %s\n", conf.GetStringSlice("gitlab.watchDog.token"))
	arr:=conf.GetStringSlice("gitlab.watchDog.token")
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
}
