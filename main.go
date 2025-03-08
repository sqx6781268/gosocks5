package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/armon/go-socks5"
)

type Config struct {
	Port     string
	User     string
	Password string
}

var (
	confPath = flag.String("c", "", "配置文件路径")
	port     = flag.String("p", "1080", "监听端口")
	user     = flag.String("u", "", "用户名")
	password = flag.String("P", "", "密码")
)

// parseFlags 解析命令行参数并返回配置信息和配置文件路径
// 输入参数：无
// 输出参数：Config 配置信息，包含端口、用户名和密码；string 配置文件路径
func parseFlags() (Config, string) {
	flag.Parse()
	return Config{
		Port:     *port,
		User:     *user,
		Password: *password,
	}, *confPath
}

// loadConfiguration 加载配置信息
// 输入参数：无
// 输出参数：Config 配置信息，包含端口、用户名和密码
func loadConfiguration() Config {
	config, confPath := parseFlags()
	return loadConfig(config, confPath)
}

// createServer 创建一个 SOCKS5 服务器实例
// 输入参数：Config 配置信息，包含端口、用户名和密码
// 输出参数：*socks5.Server SOCKS5 服务器实例；error 错误信息
func createServer(config Config) (*socks5.Server, error) {
	var authMethods []socks5.Authenticator
	if config.User != "" && config.Password != "" {
		creds := socks5.StaticCredentials{
			config.User: config.Password,
		}
		auth := socks5.UserPassAuthenticator{Credentials: creds}
		authMethods = append(authMethods, auth)
	} else {
		noAuth := socks5.NoAuthAuthenticator{}
		authMethods = append(authMethods, noAuth)
	}
	return socks5.New(&socks5.Config{
		AuthMethods: authMethods,
		Logger:      log.New(log.Writer(), "", log.LstdFlags),
	})
}

// startServer 启动 SOCKS5 服务器
// 输入参数：*socks5.Server SOCKS5 服务器实例；Config 配置信息，包含端口、用户名和密码
// 输出参数：error 错误信息
func startServer(server *socks5.Server, config Config) error {
	log.Printf("正在启动 SOCKS5 服务器，端口为 :%s，用户名为 %s，密码为 %s", config.Port, config.User, config.Password)
	return server.ListenAndServe("tcp", ":"+config.Port)
}

// loadConfig 加载配置文件并合并配置信息
// 输入参数：Config 配置信息，包含端口、用户名和密码；string 配置文件路径
// 输出参数：Config 合并后的配置信息，包含端口、用户名和密码
func loadConfig(cfg Config, confPath string) Config {
	if confPath != "" {
		file, err := os.Open(confPath)
		if err != nil {
			log.Fatalf("无法打开配置文件: %v", err)
		}
		defer file.Close()

		var fileCfg Config
		if err := json.NewDecoder(file).Decode(&fileCfg); err != nil {
			log.Fatalf("配置文件解析错误: %v", err)
		}

		// 合并配置（文件配置作为默认值）
		if cfg.Port == "1080" && fileCfg.Port != "" {
			cfg.Port = fileCfg.Port
		}
		if cfg.User == "" && fileCfg.User != "" {
			cfg.User = fileCfg.User
		}
		if cfg.Password == "" && fileCfg.Password != "" {
			cfg.Password = fileCfg.Password
		}
	}

	// if cfg.User == "" || cfg.Password == "" {
	// 	log.Fatal("必须提供用户名和密码")
	// }
	return cfg
}

// main 程序入口函数，加载配置信息，创建并启动 SOCKS5 服务器
// main函数是程序的入口点。
func main() {
	// 加载配置信息，这里假设配置信息对于服务器的创建和启动是必需的。
	config := loadConfiguration()

	// 创建服务器实例，传入配置信息。
	// 如果创建过程中遇到错误，则记录错误信息并终止程序。
	server, err := createServer(config)
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务器，传入服务器实例和配置信息。
	// 如果服务器启动过程中遇到错误，则记录错误信息并终止程序。
	if err := startServer(server, config); err != nil {
		log.Fatal(err)
	}
}
