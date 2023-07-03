package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	// read konfigurasi
	viper.AddConfigPath("belajar")
	viper.SetConfigName("properties")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	log.Print("Success Read Config viper ")
	// get konfigurasi from properties.yaml
	belajar := viper.Get("data.belajar")
	vipe := viper.Get("data.materi")
	log.Println(belajar)
	log.Println(vipe)

}
