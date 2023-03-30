package routers

import (
	"fmt"
	"gateway/common/global"
)

func HttpServerRun(services []interface{}) {
	r := NewRouter(services)

	//server := &http.Server{
	//	Addr:           viper.GetString("server.port"),
	//	Handler:        ginRouter,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//err = server.ListenAndServe()
	//if err != nil {
	//	fmt.Println("绑定HTTP到 %s 失败！可能是端口已经被占用，或用户权限不足")
	//	fmt.Println(err)
	//}
	go func() {
		if err := r.Run(global.Config.Server.Port); err != nil {
			fmt.Println("gateway启动失败, err: ", err)
		}
	}()
}
