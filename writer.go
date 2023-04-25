package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    // 打开文件
    file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("无法打开文件：", err)
        return
    }
    defer file.Close()
   
    for i := 1; i >= 1; i++ {
        line := fmt.Sprintf("这是第 %d 行数据\n", i)
    fmt.Println("数据成功写入文件.")
        _, err = file.WriteString(line)
        if err != nil {
            fmt.Println("写入文件失败：", err)
            return
        }
    time.Sleep(time.Second *1)
    }

}

