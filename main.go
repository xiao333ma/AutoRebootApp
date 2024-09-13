package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// 定义应用程序的结构体
type App struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// 定义配置结构体
type Config struct {
	Apps []App `json:"apps"`
}

// 从 JSON 文件加载配置
func loadConfig(filename string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filename) // 使用 os.ReadFile 替代 ioutil.ReadFile
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}

// 检查应用程序是否在运行
func isAppRunning(appName string) bool {
	out, err := exec.Command("tasklist").Output()
	if err != nil {
		fmt.Println("Error checking task list:", err)
		return false
	}
	return strings.Contains(string(out), appName)
}

// 启动应用程序
func startApp(appPath string) {
	cmd := exec.Command(appPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting application:", err)
	}
}

// 监控所有应用程序
func monitorApplications(apps []App) {
	// 创建一个全局 Ticker，每隔 5 秒触发一次
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop() // 程序退出时停止 ticker

	for range ticker.C {
		for _, app := range apps {
			// 检查每个应用程序是否在运行
			if !isAppRunning(app.Name) {
				fmt.Printf("%s 崩溃了，重新启动...\n", app.Name)
				startApp(app.Path)
			}
		}
	}
}

func main() {
	// 加载配置文件
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// 启动监控
	monitorApplications(config.Apps)

	// 防止主线程退出
	select {}
}
