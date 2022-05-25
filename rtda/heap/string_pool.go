package heap

import "unicode/utf16"

// stringTable 字符串常量池
var stringTable = map[string]*Object{}

// JString  从常量池存入stringtable
func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := stringTable[goStr]; ok {
		return internedStr
	}
	chars := stringToUtf16(goStr)
	jChars := &Object{class:loader.LoadClass("[C"), data:chars}
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	stringTable[goStr] = jStr
	return jStr
}
// stringToUtf16 go 字符串 转utf16 char数组
func stringToUtf16(str string) []uint16 {
	return utf16.Encode([]rune(str))
}


func GoString(object *Object) string{
	return string(utf16.Decode(object.data.([]uint16)))
}

func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := stringTable[goStr]; ok {
		return internedStr
	}
	stringTable[goStr] = jStr
	return jStr
}