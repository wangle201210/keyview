package main

import (
	"embed"
	_ "embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
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
		Name: "KeyView",
		Services: []application.Service{
			application.NewService(appService),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: false,
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	// 创建主窗口
	mainWindow := wailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "KeyView",
		Width:  1280,
		Height: 800,
		URL:    "/",
	})

	// 监听窗口关闭事件，改为最小化窗口而不是关闭
	mainWindow.RegisterHook(events.Common.WindowClosing, func(event *application.WindowEvent) {
		mainWindow.Minimise()
		event.Cancel()
	})

	// 点击程序坞图标时重新显示窗口
	wailsApp.Event.OnApplicationEvent(events.Mac.ApplicationShouldHandleReopen, func(appEvent *application.ApplicationEvent) {
		mainWindow.Show()
	})

	// 启动应用
	if err := wailsApp.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
