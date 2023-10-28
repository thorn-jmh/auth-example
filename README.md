#  LET’S TRY LOGINING!

This is a simple demo to help people write better login&auth system. It illustrated some features about cookie/CORS,and also provide some sample implementations of basic authentication. 

## HOW TO USE IT

1. server the html in somewhere (You can use JetBrains IDE or VSCode Live Server plugin)
2. change `config.yaml`, set `RemoteOrigin` to your remote page location. If you want to change `ServePort`, also modify `server` in `sample.js`. 
3. we use go for our backend, make sure you have installed go 1.18+
4. run `go mod tidy` to install dependency
5. run `go run main` to start backend server


> **Please use localhost instead of '127.0.0.1' in every place. CORS and cookie domain will not parse your domain name.**

Some trick to see what happened:

1. use developer tools of your browser to track request in **network**, and see js output in console

   <img src="https://s2.loli.net/2023/07/30/mJeErKp5A6yFZoM.png" alt="image-20230730014241578" style="zoom: 33%;" />

2. check cookies in browser (you can also find request cookie and response cookie in a request in network)

   <img src="https://s2.loli.net/2023/07/30/Lpr74zNifsH3IQY.png" alt="image-20230730014455446" style="zoom: 25%;" />

3. see logs of gin

   <img src="https://s2.loli.net/2023/07/30/6dXD92oeIp7lq3y.png" alt="image-20230730014722511" style="zoom: 50%;" />

   

## GET STARTED!

打开`sample.html`，启动后端 server。



- [x] **Cookie**

点击各个按钮，在控制台中查看输出，以及每个请求的具体信息。

后端负责 cookie 的接口在`server/cookie.go`文件中，你可以按自己的想法更改设置并尝试结果。



- [x] **CORS**

点击各个按钮，在控制台中查看输出，并查看后端 server 的 log（浏览器有可能不会显示 OPTIONS 请求，但是在后端这边一定能看到）。

后端负责 CORS 的中间件在`server/cors.go`文件中，其中的`CORSMidware`实现了 CORS 相关功能，你可以按自己的想法更改它们后尝试结果。如果需要关闭 CORS 处理，可以直接 comment `server/init.go` 中的 `r.Use(CORSMidware())`。

前面的所有 cookie 接口同样使用了这个中间件，你可以在更改后再去尝试浏览器如何在 CORS 规则下处理 cookie。所有前端请求同意使用 `sample.js`文档开头定义的 post/get 封装，你可以更改其中 fetchAPI 的参数（比如`mode`和`credential`）尝试效果。



- [x] session 实现

输入用户名后点击登录，之后再点击 auth 查看 console 观察结果。

并没有要求输入密码，只是做了 session 和 username 的对应，此处仅仅用于理解 login & authentication 两个步骤，请勿真的模仿实现登录功能。

所有登录有关的代码在 `sever/login.go`中。







- [ ] jwt 实现
- [ ] oauth 实现 (google为例?)
- [ ] CSRF & XSF，倒是找到一些例子，但是要用 IE 演示，母鸡有没有通用方法

