package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	AppName string `json:"app_name"`
	Mysql   *Mysql `json:"mysql"`
	App     *App   `json:"app"`
}

type App struct {
	BasePath string `json:"base_path"`
}

type Mysql struct {
	User   string `json:"user"`
	Dbname string `json:"dbname"`
	Pass   string `json:"pass"`
}

func main() {

	//conf := Config{
	//	AppName: "go",
	//	Mysql: &Mysql{
	//		"root",
	//		"go",
	//		"root",
	//	},
	//	App: &App{
	//		BasePath: "C:\\Users\\Admin\\go\\src\\Learn-Go\\06\\conf",
	//	},
	//}

	//b,_ := json.Marshal(conf)
	//
	//fmt.Println(string(b))
	//b,_ = json.MarshalIndent(conf,"","	")
	//fmt.Println(string(b))
	//
	//
	//file ,_ :=os.OpenFile(conf.App.BasePath+"\\config.json",os.O_WRONLY|os.O_CREATE,0666)
	//defer file.Close()

	//file.WriteString(string(b))
	//

	//file,_ := os.Open(conf.App.BasePath+"\\config.json")

	//写入
	//json.NewEncoder(file).Encode(conf)

	var conf Config
	date := []byte(`{"app_name":"go","mysql":{"user":"root","dbname":"go","pass":"root"},"app":{"base_path":"C:\\Users\\Admin\\go\\src\\Learn-Go\\06\\conf"}}`)

	json.Unmarshal(date, &conf)

	fmt.Println(conf.App.BasePath)

}
