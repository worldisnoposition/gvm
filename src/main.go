package main

import (
	"fmt"
)

func main() {
	cmd := parseCmd()
	//if cmd.versionFlag {
	//	fmt.Println("version 1.0.0")
	//} else if cmd.helpFlag || cmd.class == "" {
	//	printUsage()
	//} else {
	//	cmd.XjreOption = "C:\\Program Files\\Java\\jre1.8.0_271"
	//	cmd.class = "java.lang.String"
	//	startJVM(cmd)
	//}

	//cmd.XjreOption = "C:\\Program Files\\Java\\jre1.8.0_271"
	//cmd.class = "java.lang.String"
	cmd.XjreOption = "C:\\Program Files\\Java\\jre1.8.0_271"
	cmd.classpath = "D:\\coding\\java\\demo\\target\\classes"
	cmd.class = "com.example.gvm.CircleCalculator"
	startJVM(cmd)
	fmt.Println(cmd.helpFlag)
	fmt.Println(cmd.class)
}
