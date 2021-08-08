 1. 目录结构
    ```cassandraql
    conf：用于存储配置文件
    middleware：应用中间件
    models：应用数据库模型
    pkg：第三方包
    routers 路由逻辑处理
    runtime 应用运行时数据
    ```
 2. 执行命令
    ```shell
    # 拉取go-ini/ini的依赖包
    go get -u github.com/go-ini/ini
    go get -u github.com/jinzhu/gorm
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/unknwon/com
    go get -u github.com/astaxie/beego/validation
    go get -u github.com/dgrijalva/jwt-go
    go get -u github.com/fvbock/endless
    
    go get -u github.com/swaggo/swag/cmd/swag
    go get -u github.com/swaggo/gin-swagger
    go get -u github.com/swaggo/gin-swagger/swaggerFiles
    go get -u github.com/robfig/cron
    go get -u github.com/gomodule/redigo/redis
    ```