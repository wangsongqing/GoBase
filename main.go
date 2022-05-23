package main

import (
	"GoBase/app/controllers"
	"flag"
)

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
	}

}
