package main

import (
    "fmt"
    "os"
    "time"
    "bufio"
)

func main() {
    // 打开文件
    time.Sleep(time.Second *10)
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("无法打开文件：", err)
        return
    }
    defer file.Close()

    // 读取数据
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    time.Sleep(time.Second *2)

    fmt.Println("数据读取完毕.")
   }
}

