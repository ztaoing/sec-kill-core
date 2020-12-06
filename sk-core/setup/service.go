/**
* @Author:zhoutao
* @Date:2020/7/6 下午10:09
 */

package setup

import (
	"fmt"
	"github.com/ztaoing/sec-kill-core/sk-core/service/svc_redis"
	register "github.com/ztaoing/sec-kill-pkg/pkg/discover"
	"os"
	"os/signal"
	"syscall"
)

func RunService() {
	//启动处理线程
	svc_redis.RunProcess()
	errChan := make(chan error)

	go func() {
		//服务注册
		register.Register()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	//等待
	error := <-errChan
	//服务注销
	register.Deregister()
	fmt.Println(error)
}
