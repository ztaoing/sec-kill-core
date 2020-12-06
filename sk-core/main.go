/**
* @Author:zhoutao
* @Date:2020/7/6 下午11:17
 */

package main

import "github.com/ztaoing/sec-kill-core/sk-core/setup"

func main() {
	setup.InitZk()
	setup.InitRedis()
	setup.RunService()
}
