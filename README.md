#  LET’S TRY LOGINING!

This is a simple demo to help people write better login&auth system. It illustrate some features about cookie/CORS,and also provide some sample implementations of basic authentication. 

## HOW TO USE IT

1. server the html in somewhere (You can use JetBrains IDE or VSCode Live Server plugin)
2. change `config.yaml`, set `RemoteOrigin` to your remote page location
3. we use go for our backend, make sure you have installed go 1.18+
4. run `go mod tidy` to install dependency
5. run `go run main` to start backend server



Some trick to see what happened in it:
**[WARN:pictures out of date]**

1. use developer tools of your browser to track request

   <img src="https://s2.loli.net/2023/07/30/mJeErKp5A6yFZoM.png" alt="image-20230730014241578" style="zoom: 33%;" />

2. check cookies in browser

   <img src="https://s2.loli.net/2023/07/30/Lpr74zNifsH3IQY.png" alt="image-20230730014455446" style="zoom: 25%;" />

3. see logs of gin

   <img src="https://s2.loli.net/2023/07/30/6dXD92oeIp7lq3y.png" alt="image-20230730014722511" style="zoom: 50%;" />

   

## GET STARTED!

- [x] cookie

打开`sample.html`，启动后端 server。

点击各个按钮，在控制台中查看输出，以及每个请求的具体信息。

后端负责 cookie 的接口在`server/cookie.go`文件中，你可以按自己的想法更改设置并尝试结果。



- [x] cors

打开`sample.html`，启动后端 server。

点击各个按钮，在控制台中查看输出，并查看后端 server 的 log（浏览器似乎不会显示 OPTIONS 请求，只能在后端这边查看）。

后端负责 cors 的中间件在`server/cors.go`文件中，其中的`CORSMidware`实现了 CORS 相关功能，你可以按自己的想法更改它们后尝试结果。

前面的所有 cookie 接口同样使用了这个中间件，你可以在更改后再去尝试浏览器如何在 CORS 规则下处理 cookie。



- [ ] session 实现
- [ ] jwt 实现
- [ ] oauth 实现(google为例?)
- [ ] CSRF & XSF，倒是找到一些例子，但是要用 IE 演示（谁电脑上还有IE啊x），在想办法

