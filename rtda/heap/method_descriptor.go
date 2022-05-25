package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}
//  TODO 解析方法描述符 提取参数列表 返回类型
func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	//desc :=  "(CFDZBSIJLjava/lang/String;Ljava/lang/Integer;)J"


	return &MethodDescriptor{}
}
