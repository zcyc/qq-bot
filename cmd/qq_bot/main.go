package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"qq-bot/common/websocket"
	"qq-bot/integration/qq"
	"qq-bot/internal/facade/at_message"
)

func main() {
	// 设置配置文件格式
	viper.SetConfigType("toml")
	// 设置配置文件目录
	viper.AddConfigPath("./conf")
	// 设置配置文件名
	viper.SetConfigName("config")
	// 加载配置
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	// 打印当前配置
	fmt.Println("当前配置: ", viper.AllSettings())

	// 初始化 GORM
	//dbConf := &infra.Dsn{
	//	Host:     viper.GetString("db.host"),
	//	User:     viper.GetString("db.user"),
	//	Password: viper.GetString("db.password"),
	//	Name:     viper.GetString("db.name"),
	//	Port:     viper.GetString("db.port"),
	//}
	//infra.InitDB(dbConf)

	// 创建 token
	ctx := context.Background()
	sandBoxMode := viper.GetBool("sand_box_mode")
	appId := viper.GetUint64("bot.app_id")
	appToken := viper.GetString("bot.token")
	botToken := qq.NewBotToken(appId, appToken)

	// 初始化 client
	botClient := qq.NewBotClient(sandBoxMode, botToken)

	// 获取网关
	ap, err := botClient.GetAccessPoint(ctx)
	if err != nil {
		panic(err)
	}

	// 注册消息处理器
	intent := at_message.NewHandler(botClient)
	if err = websocket.NewSessionManager().Start(ap, botToken, intent); err != nil {
		panic(err)
	}
}
