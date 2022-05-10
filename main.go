package main

import (
	"fmt"
	"jvm-go/classpath"
	"jvm-go/cmd"
	"jvm-go/interpreter"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

func main() {
	parseCmd := cmd.ParseCmd()
	if parseCmd.Help {
		cmd.PrintUsage()
	} else if parseCmd.Version {
		fmt.Println("java version \"1.8.0_221\"\nJava(TM) SE Runtime Environment (build 1.8.0_221-b11)\nJava HotSpot(TM) 64-Bit Server VM (build 25.221-b11, mixed mode)")
	} else {
		startVm(parseCmd)
	}
}

func startVm(cmd *cmd.Cmd) {
	c := classpath.NewClassPath(cmd.Xbootclasspath, cmd.Classpath)
	loader := heap.NewClassLoader(c)
	class := loader.LoadClass(cmd.Class)
	main := class.FindMainClass()
	thread := rtda.NewThread()
	interpreter.Interpret(thread, main)

}

func testStackFrame() {

}
