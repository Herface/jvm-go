package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	return &MethodDescriptor{}
}
