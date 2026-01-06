package main

import (
	"embed"
	_ "embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wangle201210/keyview/internal/app"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 创建应用服务
	appService := app.NewAppService()

	// 初始化服务
	if err := appService.Init(); err != nil {
		log.Fatalf("Failed to initialize app service: %v", err)
	}

	// 创建 Wails 应用
	wailsApp := application.New(application.Options{
		Name:        "KeyView",
		Description: "键盘使用历史记录查看工具",
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Services: []application.Service{
			application.NewService(appService),
		},
		Mac: application.MacOptions{
			ActivationPolicy: application.ActivationPolicyRegular,
			ApplicationShouldTerminateAfterLastWindowClosed: false,
		},
	})

	// 创建主窗口
	wailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "KeyView - 键盘使用历史记录",
		Width:  1280,
		Height: 800,
		Mac: application.MacWindow{
			TitleBar: application.MacTitleBarDefault,
		},
		BackgroundColour: application.NewRGB(255, 255, 255),
		URL:              "/",
	})

	// 启动应用
	if err := wailsApp.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
