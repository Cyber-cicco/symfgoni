package utils

import (
	"bytes"
	"unicode"
)

func ConvertPathToNameSpace(path string) string {
    strLen := getStrLen(path) 
    buff := make([]byte, strLen)
    for i := 0; i < strLen; i++ {
        if path[i] == '/' {
            buff[i] = '\\'
            continue
        }
        buff[i] = path[i]
    }
    return "App\\" + string(buff)
}

func getStrLen(path string) int {
    if path[len(path)-1] == '/' {
        return len(path) - 1
    }
    return len(path)
}

func ConvertNameToRoute(ctrlName string) string {
    var buffer bytes.Buffer
    lastPos := 0
    for i, l := range ctrlName {
        if unicode.IsUpper(l) {
            buffer.WriteString(ctrlName[lastPos:i] + "/")
            buffer.WriteByte(byte(unicode.ToLower(l)))
            lastPos = i + 1
        }
    }
    buffer.WriteString(ctrlName[lastPos:])
    return buffer.String() 
}

func ConvertNameToRouteName(ctrlName string) string {
    var buffer bytes.Buffer
    lastPos := 0
    buffer.WriteString("app")
    for i, l := range ctrlName {
        if unicode.IsUpper(l) {
            buffer.WriteString(ctrlName[lastPos:i] + "_")
            buffer.WriteByte(byte(unicode.ToLower(l)))
            lastPos = i + 1
        }
    }
    buffer.WriteString(ctrlName[lastPos:])
    return buffer.String() 
}
