package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viperInit()

}
func viperInit() {
	viper.SetConfigFile("./configure.yaml")
	if err:=viper.ReadInConfig();err!=nil{
		log.Println(err)
	}
	s ,ok:= viper.Get("mysql.user").(string)
	if ok{
		fmt.Println(s)
	}
	

	// fmt.Println(viper.Get("mysql.user"))
	// fmt.Println(viper.Get("shit.purpose"))
	
	// fmt.Printf("fp(fmt.Sprint(viper.Get(\"mysql.user\")), fmt.Sprint(viper.Get(\"shit.purpose\"))): %v\n", fp(fmt.Sprint(viper.Get("mysql.user")), fmt.Sprint(viper.Get("shit.purpose"))))
}

func fp(a,b string)string{
	return a+b
}