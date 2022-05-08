package main

import (
	"fmt"
	"jvm-go/classfile"
	"jvm-go/classpath"
	"jvm-go/cmd"
	"jvm-go/rtda"
	"math"
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
	class, _, err := c.ReadClass(cmd.Class)
	if err != nil {
		panic(err)
	}
	//fmt.Println(len(class))
	_, err = classfile.Parse(class)
	if err != nil {
		panic(err)
	}
	//fmt.Println(classdata)
	testStackFrame()

}

func testStackFrame() {
	localVars := rtda.NewLocalVars(4)
	stack := rtda.NewOperandStack(4)

	stack.PushLong(math.MaxInt64 - 1)
	a := stack.PopLong()
	localVars.SetLong(0, a)

	stack.PushLong(1)
	b := stack.PopLong()
	localVars.SetLong(2, b)

	a = localVars.GetLong(0)
	stack.PushLong(a)
	b = localVars.GetLong(2)
	stack.PushLong(b)

	a = stack.PopLong()
	b = stack.PopLong()
}
