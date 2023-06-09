package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	guess1()
}

///猜數字遊戲版本1
func guess1() {
	rand.Seed(time.Now().UnixNano()) ///亂數種子
	ans := rand.Intn(100) + 1        ///隨機1~100
	var userScan int                 ///設定使用者輸入變數(int)
	var win bool
	fmt.Println("Guess a number between 1~100, you have 10 time can try")

	for i := 10; i > 0; i-- {
		fmt.Scan(&userScan) ///使用者輸入
		switch {
		case userScan > 100 || userScan < 1: ///輸入值超出範圍
			fmt.Println("Please guess between 1~100!")
			i++
			break
		case userScan > ans: ///猜太大
			fmt.Printf("This number is too big! You have %d time\n", i)
		case userScan < ans: ///猜太小
			fmt.Printf("This number is too small! You have %d time\n", i)
		case userScan == ans: ///猜中答案，結束循環
			fmt.Println("You got it!")
			i = 0
			win = true
		}
	}
	if win != true { ///猜失敗，顯示答案
		fmt.Println("You are lose.... Answer is ", ans)
	}
}

///猜數字遊戲版本2
func guess2() {
	rand.Seed(time.Now().UnixNano()) ///亂數種子(奈米秒)
	ans := rand.Intn(100) + 1        ///隨機1~100
	win := false
	fmt.Println("Guess a number between 1~100, you have 10 time can try")

	reader := bufio.NewReader(os.Stdin)         ///讀取鍵盤輸入(string)
	for guesses := 0; guesses < 10; guesses++ { ///10次循環
		fmt.Println("You have", 10-guesses, "guesses left")
		fmt.Print("Make you guess:")
		input, err := reader.ReadString('\n') ///讀取輸入到ENTER鍵停止
		if err != nil {                       ///錯誤回報
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  ///去除頭尾單元字
		guess, err := strconv.Atoi(input) ///字串轉整數
		if err != nil {
			log.Fatal(err)
		}

		if guess > ans { ///猜太大
			fmt.Println("This number is too big!")
		} else if guess < ans { ///猜太小
			fmt.Println("This number is too small!")
		} else { ///猜中
			win = true
			fmt.Println("You got it!")
			break ///中止迴圈
		}
	}
	if win != true { ///猜失敗，顯示答案
		fmt.Println("You are lose.... Answer is ", ans)
	}
}
