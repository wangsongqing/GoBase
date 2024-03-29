package main

import (
	"GoBase/app/controllers"
	"flag"
)

// 程序执行  go run main.go --controllerName=time
func main() {
	controllerName := flag.String("controllerName", "", "")

	flag.Parse()

	switch *controllerName {
	case "array": // 数组
		controllers.RunArray()
		break
	case "slice":
		controllers.RunSlice()
		break
	case "map":
		controllers.MapSlice()
		break
	case "time":
		controllers.TimeTest()
	case "struct":
		controllers.StructTest()
	case "structfun":
		controllers.FunStructTest()
	case "ooptest":
		controllers.OopTest()
	case "interface":
		controllers.InterfaceTest()
	case "file":
		controllers.TestFiles()
	case "json":
		controllers.JsonTest()
	case "goroutine":
		controllers.GoroutineTest()
	case "channel":
		controllers.ChannelTest()
	case "string":
		controllers.StringTest()
	case "fun":
		controllers.FuncTest()
		break
	case "defer":
		controllers.DeferTest()
		break
	}

}
