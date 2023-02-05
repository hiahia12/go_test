package main

import (
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"go_test/user/internal/config"
	"go_test/user/internal/handler"
	"go_test/user/internal/svc"
)

var configFile = flag.String("f", "user/etc/user-api.yaml", "the config file")

func main() {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{ // nacos服务的ip和端口及请求方式
			IpAddr:      "192.168.10.105",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	//注册服务名
	serviceName := "nacos-study-go-one"

	//注册实例：RegisterInstance
	flag, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "192.168.10.105",                //注册服务的ip
		Port:        8888,                            //注册服务的端口
		ServiceName: serviceName,                     //注册服务名
		Weight:      10,                              //权重
		Enable:      true,                            //是否可用
		Healthy:     true,                            //健康状态
		Ephemeral:   true,                            //零时节点（服务下线之后，nacos 上注册信息会删除）
		Metadata:    map[string]string{"name": "go"}, //元数据
		ClusterName: "DEFAULT",                       // 默认值DEFAULT  集群名称
		GroupName:   "DEFAULT_GROUP",                 // 默认值DEFAULT_GROUP 组名称
	})

	if flag {
		fmt.Printf("服务：%s注册成功", serviceName)
	} else {
		fmt.Printf("服务：%s注册失败，错误信息：%v", serviceName, err)
	}

	fmt.Println()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	//// 注销实例：DeregisterInstance
	//flag1, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
	//	Ip:          "127.0.0.1",
	//	Port:        8089,
	//	ServiceName: serviceName,
	//	Ephemeral:   true,
	//	Cluster:     "DEFAULT",       // 默认值DEFAULT
	//	GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	//})
	//
	//if flag1 {
	//	fmt.Printf("服务：%s注销成功", serviceName)
	//} else {
	//	fmt.Printf("服务：%s注销失败，错误信息：%v", serviceName, err)
	//}
	//
	//time.Sleep(time.Minute)
}
