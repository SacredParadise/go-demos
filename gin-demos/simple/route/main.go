package main

import (
	"fmt"
	"gin-demos/simple/route/app/blog"
	"gin-demos/simple/route/app/shop"
	"gin-demos/simple/route/routers"
)

func main() {
	routers.Include(blog.Routers, shop.Routers)
	r := routers.Init()

	if err := r.Run(); err != nil {
		fmt.Println("Startup service failed, err:%v \n", err)
	}
}
