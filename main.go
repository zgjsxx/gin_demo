package main

import (
	"fmt"
	"gin_demo/app/blog"
	"gin_demo/app/shop"
	"gin_demo/routers"
)

func main() {
	// 家在多个APP的路由配置
	routers.Include(shop.Routers,blog.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}