# gin学习 (学而时习之不亦说乎)
* gin.Logger() 实际上是使用的标准库里的 log包
* serve.Run() 实际上是调用的http包里的ListenAndServe()
* ListenAndServe() 实际上是调用的net包里的Listen() 和http包里的Serve()
* router是使用的julienschmidt / httprouter 包
* binding 是使用的validate.v8

## 参考资料
* [ Gin 文档中文翻译](https://learnku.com/docs/gin-gonic/2018/gin-readme/3819)
* [开始学习 Go Gin web 框架，同时写了一份教程](https://www.v2ex.com/t/589372)
* [Gin框架开发与实战](https://www.chaindesk.cn/witbook/19/329)
* [Gin框架源码解析](https://www.cnblogs.com/yjf512/p/9670990.html)
