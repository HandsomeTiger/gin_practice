# gin
> github.com/gin-gonic/gin

## 安装gin
1. 获取并安装

    `go get -u github.com/gin-gonic/gin`

2. 创建项目并引入gin包
    
    `import "github.com/gin-gonic/gin"`
    
3. 使用govendor创建依赖管理

## 快速开始   
    package main
    
    import "github.com/gin-gonic/gin"
    
    func main() {
        r := gin.Default()
        r.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "pong",
            })
        })
        r.Run() // 在 0.0.0.0:8080 上监听并服务
    }

## gin.Default()
    func Default() *Engine {
    	debugPrintWARNINGDefault()
    	engine := New()
    	engine.Use(Logger(), Recovery())
    	return engine
    }
> gin.Default 实际上是实例化了一个-Engine,并加载了Logger()日志中间件，和Recovery()恢复中间件

## Engine（翻译：引擎）
* func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes 
* func (engine *Engine) Run(addr ...string) (err error)

## Context（翻译：上下文）
* func (c *Context) Copy() *Context   

        func (c *Context) Copy() *Context {
                    	var cp = *c
                    	cp.writermem.ResponseWriter = nil
                    	cp.Writer = &cp.writermem
                    	cp.index = abortIndex
                    	cp.handlers = nil
                    	cp.Keys = map[string]interface{}{}
                    	for k, v := range c.Keys {
                    		cp.Keys[k] = v
                    	}
                    	paramCopy := make([]Param, len(cp.Params))
                    	copy(paramCopy, cp.Params)
                    	cp.Params = paramCopy
                    	return &cp
        }
                    
* func (c *Context) Abort() 
* func (c *Context) JSON(code int, obj interface{})
* func (c *Context) AbortWithStatusJSON(code int, jsonObj interface{}) 
* func (c *Context) Set(key string, value interface{}) 
* func (c *Context) Get(key string) 
* func (c *Context) Param(key string) string
* func (c *Context) Query(key string) string 
* func (c *Context) PostForm(key string) 
* func (c *Context) FormFile(name string) (*multipart.FileHeader, error) 
* func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error 
* func (c *Context) Bind(obj interface{}) error 
* func (c *Context) BindJSON(obj interface{}) error 
* func (c *Context) ShouldBindJSON(obj interface{}) error 
* func (c *Context) ShouldBindQuery(obj interface{}) error 
* func (c *Context) ShouldBind(obj interface{}) error 
* func (c *Context) Redirect(code int, location string)                    

## 日志

## 路由
> golang中路由使用**github.com/julienschmidt/httprouter**包,该包所有的特性和方法，gin都支持
### RouterGroup
* func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes 
* func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup
* func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes
* ... 

## 绑定模型和验证
> Gin 使用 go-playground/validator.v8 进行验证。
* Must bind
* Should bind

## Json
