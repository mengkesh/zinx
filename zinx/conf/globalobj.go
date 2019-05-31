package conf

import (
	"io/ioutil"

	"encoding/json"
)

type GlobalObj struct {
	Host string //当前监听的IP
	Port int //当前的监听Port
	Name string	//当前zinxserver的名称
	Version string //当前框架的版本号
	MaxPackageSize uint32 //每次Read一次的最大长度
	WorkPoolSize uint32
	MaxWorkerTaskLen uint32
	MaxConn uint32
}

var GlobalObject *GlobalObj
func(g *GlobalObj)LoadConfig(){
	data,err:=ioutil.ReadFile("conf/zinx.json")
	if err!=nil{
		panic(err)
	}
	err=json.Unmarshal(data,&g)
	if err!=nil{
		panic(err)
	}
}
func init(){
	GlobalObject=&GlobalObj{
		Name:"ZinxServerApp",
		Host:"0.0.0.0",
		Port:8999,
		Version:"V0.4",
		MaxPackageSize:512,
		WorkPoolSize:10,
		MaxWorkerTaskLen:4096,
		MaxConn:1000,
	}
	GlobalObject.LoadConfig()
}