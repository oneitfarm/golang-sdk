# golang version sdk for Oneitfarm 

## Usage:

```golang

// 以gin框架举例:

import (
    "github.com/gin-gonic/gin"
    // 导入golang sdk
    sdk "github.com/oneitfarm/golang-sdk"
)

server := gin.Default()

// 使用sdk.Gin wrapper来包装handler
server.GET("/test", sdk.Gin(func(ctx *gin.Context, apiCtx *sdk.ApiContext) {
    // handler内部会得到 ApiContext对象
    // 通过ApiContext对象完成一系列功能

    // 调用其它服务
    result := apiCtx.Get("other_service", "/xxx_api")

    // post表单
    result2 := apiCtx.PostForm("other_service2", "xxx_api", map[string]string{"key": "value"})


    // handler处理结束， 返回结果
    apiCtx.ReturnOk("success result")
}))


```
