package common

import "jvm-go/rtda/heap"

func CheckIndex(length int, index int32) {
	if int(index) >= length {
		panic("java.lang.ArrayIndexOutOfBoundsException")
	}
}

func CheckNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
