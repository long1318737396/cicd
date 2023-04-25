package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
    "bufio"
)

func main() {
    // 创建文件
    file, err := os.Create("test.txt")
    if err != nil {
        fmt.Println("创建文件失败：", err)
        return
    }
    defer file.Close()

    // 启动写入子进程
    writerCmd := exec.Command("go", "run", "writer.go")
    writerCmd.Stdout = os.Stdout
    writerCmd.Stderr = os.Stderr
    writerCmd.Start()

    // 启动读取子进程
    readerCmd := exec.Command("go", "run", "reader.go")
    readerCmd.Stdout = os.Stdout
    readerCmd.Stderr = os.Stderr
    readerCmd.Stdin = bufio.NewReader(file) 
    readerCmd.Start()

    // 等待子进程完成
       writerCmd.Wait()
     readerCmd.Wait()

    for {
    time.Sleep(time.Second*2)
}
    fmt.Println("程序运行完毕.")
}

