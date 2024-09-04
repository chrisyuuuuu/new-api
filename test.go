package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// 创建一个带有超时设置的 http.Client
	client := &http.Client{
		Timeout: 3 * time.Second, // 设置超时时间为10秒
	}

	// 发起 GET 请求
	resp, err := client.Get("https://httpbin.org/delay/5") // 一个会延迟5秒响应的测试API
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP request failed with status %d\n", resp.StatusCode)
		return
	}

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 输出响应内容
	fmt.Println("Response Body:", string(body))
}
