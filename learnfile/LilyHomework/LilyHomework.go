package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'lilysHomework' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

//  func lilysHomework(arr []int32) int32 {
// 	var ansS, ansB int32
// 	long := len(arr)
// 	arr2 := append(arr[:0:0], arr...)
// 	for i := 0; i <= long-2; i++ {
// 		k, l := i, i
// 		for j := i + 1; j <= long-1; j++ {
// 			if arr[j] < arr[i] && arr[j] < arr[k] {
// 				k = j
// 			}
// 			if arr2[j] > arr2[i] && arr2[j] > arr2[l] {
// 				l = j
// 			}
// 			if j == long-1 {
// 				if arr[k] < arr[i] {
// 					arr[i], arr[k] = arr[k], arr[i]
// 					ansS++
// 				}
// 				if arr2[l] > arr2[i] {
// 					arr[i], arr[k] = arr[k], arr[i]
// 					ansB++
// 				}
// 			}
// 		}
// 	}
// 	if ansB < ansS {
// 		return ansB
// 	}
// 	return ansS
// }

func lilysHomework(arr []int32) int32 {
	var As, Des int32
	ansA, ansD := make(chan int32), make(chan int32)
	go AscendingSort(arr, ansA)
	go DescenderSort(arr, ansD)
	As = <-ansA
	Des = <-ansD
	if As < Des {
		return As
	}
	return Des
}

func AscendingSort(arr []int32, channel chan int32) {
	var ansA int32
	for i := 0; i <= len(arr)-2; i++ {
		k := i
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[j] < arr[i] && arr[j] < arr[k] {
				k = j
			}
			if j == len(arr)-1 {
				if arr[k] < arr[i] {
					arr[i], arr[k] = arr[k], arr[i]
					ansA++
				}
			}
		}
	}
	channel <- ansA
}

func DescenderSort(arr []int32, channel chan int32) {
	var ansD int32
	for i := 0; i <= len(arr)-2; i++ {
		k := i
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[j] > arr[i] && arr[j] > arr[k] {
				k = j
			}
			if j == len(arr)-1 {
				if arr[k] > arr[i] {
					arr[i], arr[k] = arr[k], arr[i]
					ansD++
				}
			}
		}
	}
	channel <- ansD
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := lilysHomework(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
