package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/guestbook", viewHandler) //以viewhandler執行/guestbook的請求
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil) //傳遞HTTP請求
	log.Fatal(err)
}

//Guestbook is a Struct used in rendering view.html.//用在view.html的結構
type Guestbook struct {
	SignatureCount int      //簽到數
	Signatures     []string //簽到列表
}

//check calls log.Fatal on non-nil error.//err不回空值記錄錯誤訊息
func check(err error) {
	if err != nil { //檢查err值不為空的話
		log.Fatal(err) //終止程式並回傳
	}
}

//viewHandler reads guestbook signature and displays them together.//讀取簽到簿簽名並顯示
//with a count of all signature.//顯示簽名總數
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")    //從檔案讀取簽名
	html, err := template.ParseFiles("view.html") //以view.html建立模板
	check(err)                                    //確認err
	guestbook := Guestbook{                       //以guestbook變數建立結構
		SignatureCount: len(signatures), //簽名總數
		Signatures:     signatures,      //簽名
	}
	err = html.Execute(writer, guestbook) //把結果寫到ResponseWriter
	check(err)
}

//newHandler displays a form a enter a signature.//進入簽名畫面
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html") //以new.html建立模板
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

//createHandler takes a POST request with a signature to add,
//and appends it to the signature file.//新增簽名到簽名檔案
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE                     ////////////////////////////////////////
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600)) //開啟檔案。如果檔案不存在就新建，開啟成功就附加進去
	check(err)
	_, err = fmt.Fprintln(file, signature) //將表單內容加到檔案中
	check(err)
	err = file.Close() //關閉檔案
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound) //重新導向到簽到簿主頁
}

//getString returns a slice of strings read from fileName,
//one string per line. 把檔案讀取到的每一行以文字傳回到切片
func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName) //開啟檔案
	if os.IsNotExist(err) {        //如果檔案不存在
		return nil //回傳nil
	}
	check(err)
	defer file.Close()                //在最後關閉檔案
	scanner := bufio.NewScanner(file) //讀取檔案內容
	for scanner.Scan() {              //讀取直到最後一行
		lines = append(lines, scanner.Text()) //把檔案內容新增到切片
	}
	check(scanner.Err()) //檢查讀取中的錯誤
	return lines
}
