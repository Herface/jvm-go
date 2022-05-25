package main

import (
	"fmt"
	"jvm-go/classpath"
	"jvm-go/cmd"
	"jvm-go/interpreter"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
	"regexp"
	"runtime"
	"strings"
	"net"
)

type A interface {
	doA()
}
type B interface {
	doA()
	doB()
}


func main() {
	parseCmd := cmd.ParseCmd()
	if parseCmd.Help {
		cmd.PrintUsage()
	} else if parseCmd.Version {
		fmt.Println("java version \"1.8.0_221\"\nJava(TM) SE Runtime Environment (build 1.8.0_221-b11)\nJava HotSpot(TM) 64-Bit Server VM (build 25.221-b11, mixed mode)")
	} else {
		//startVm(parseCmd)
		desc :=  "(CFDZBSIJLjava/lang/String;ZBSLjava/lang/Integer;)J"
		compile, _ := regexp.Compile("[BSIJCFDZ]+[)L]")
		stringList := compile.FindAllString(desc, -1)
		primitiveTypes := make([]string, 0)
		if len(stringList) > 0 {
			primitiveTypes = strings.Split(stringList[0], "")
		}

		index := strings.Index(desc, "L")
		types := desc[index:]
		split := strings.Split(types, ";")
		fmt.Println(primitiveTypes)
		fmt.Println(split)
		runtime.resetspinning
		runtime.selectgo
		runtime.netpoll
		runtime.poll_runtime_pollOpen


	}

}

func startVm(cmd *cmd.Cmd) {
	c := classpath.NewClassPath(cmd.Xbootclasspath, cmd.Classpath)
	loader := heap.NewClassLoader(c)
	class := loader.LoadClass(cmd.Class)
	mainMethod := class.FindMainClass()
	// 启动main线程
	thread := rtda.NewThread()
	thread.PushFrame(thread.NewFrame(mainMethod))
	loadClass := loader.LoadClass("java.lang.Thread")
	jlThread := loadClass.NewObject()
	thread.SetJLThread(jlThread)
	interpreter.Interpret(thread)
}
