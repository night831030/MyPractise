package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1]) //開啟輸入名稱的檔案 //go run main.go xxx.txt <-檔案名在後
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) //讀取檔案內容
	for scanner.Scan() {
		fmt.Println(scanner.Text()) //以迴圈輸出
	}

	// io.Copy(os.Stdout, file) //複製檔案內容輸出至終端
}
