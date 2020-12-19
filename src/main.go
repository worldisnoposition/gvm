package main

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
	//cmd.XjreOption = "C:\\Program Files\\Java\\jdk-11.0.9\\lib\\src\\java.base"
	cmd.classpath = "D:\\code\\go\\gvm\\java\\classes"
	cmd.class = "com.example.gvm.CircleCalculator"
	cmd.class = "com.example.gvm.GaussTest"
	cmd.class = "com.example.gvm.MyObject"
	cmd.class = "com.example.gvm.InvokeDemo"
	cmd.class = "com.example.gvm.FibonacciTest"
	cmd.cpOption = "D:\\code\\go\\gvm\\java\\classes"
	cmd.verboseInstFlag = true
	startJVM(cmd)
}
