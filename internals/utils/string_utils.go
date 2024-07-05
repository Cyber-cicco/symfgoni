package utils

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
