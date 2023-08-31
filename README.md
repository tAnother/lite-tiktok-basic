## 目录说明

### 核心服务结构：

1. controller/

    - protocols/: 存放网络通信结构

2. service/：业务逻辑封装

3. repository/：crud操作

4. model/: 存放实体类定义（数据库依赖于这些定义）

### 其他：
1. config/：redis相关初始化，mysql相关初始化，服务器地址端口设置

2. middleware/：拦截器

3. utils/：utils

## 注意事项：
1. 需在config中设置正确的mysql数据库端口，密码，用户名
2. 设置正确的redis数据库端口，密码，用户名

<br>

## TODO
### 未解决问题:
1. 没有生成视频封面，也没有设置封面正确的路径
2. 暂无固定后端域名orIP，所以运行程序时需要手动输入服务端地址和端口，以获取正确的视频路径

### 优化方向
[ ] 1. 优化数据存储（尝试mongoDB）

[ ] 2. 优化依赖注入（别放router.go里）

[ ] 3. 找个好用的log库

[ ] 4. 考虑是否有并发风险

[ ] 5. protocols应该放在哪里？token相关的东西应该单独出来还是放在service层？controller应该负责construct response吗？