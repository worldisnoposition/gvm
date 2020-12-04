package main

import (
	"flag"
	"fmt"
	"gvm/src/classfile"
	"gvm/src/classpath"
	"os"
	"strings"
)

func parseCmd() *Cmd {
	fmt.Println(flag.Args())
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "? ", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.classpath, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Print("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
	//classData, _, err := cp.ReadClass(className)
	//if err != nil {
	//	fmt.Printf("Could not find or load main class %s\n", cmd.class)
	//	return
	//}
	//fmt.Printf("class data:%v\n", classData)
	//fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags:0x%x\n", cf.AccessFlags())
	fmt.Printf("this class:%v\n", cf.ClassName())
	fmt.Printf("super class:%v\n", cf.SuperClassName())
	fmt.Printf("interfaces:%v\n", cf.InterfaceNames())
	fmt.Printf("fields count:%v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("	%s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, f := range cf.Methods() {
		fmt.Printf("	%s\n", f.Name())
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
	XjreOption  string
	classpath   string
}
