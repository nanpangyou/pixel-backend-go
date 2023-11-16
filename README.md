# pixel backend with go

1. init go program

   ```bash
   go mod init github.com/yourname/yourproject
   ```

2. go build

   ```bash
   go build .
   ```

     编译后会生成一个可执行文件，直接运行即可

3. go run

   ```bash
   go run main.go
   ```

     直接运行go文件

4. 依赖安装 `go mod tidy`

## 开发过程

### 新建文件夹，并将不同功能文件放入对应文件夹下

   cmd 作为一些功能的入口，比如启动服务，启动任务等
   internal 作为一些内部的功能，比如数据库连接，redis连接等

### 使用基础例子实现 ping pong

```go

 r.GET("/ping", controller.PingHandler)

func PingHandler(c *gin.Context) {
 c.String(http.StatusOK, "pong1")
 c.Writer.WriteString("pong2")
}
// 会返回pong1pong2
```

### 写单元测试

抄文档

测试文件均以 `_test.go` 结尾

```go

// ping_test.go
func TestPingRoute(t *testing.T) {
 router := router.New()

 w := httptest.NewRecorder()
 req, _ := http.NewRequest("GET", "/ping", nil)
 router.ServeHTTP(w, req)

 assert.Equal(t, 200, w.Code)
 assert.Equal(t, "pong1pong2", w.Body.String())
}
```
