[README-ENG](./README-ENG.md)

## 目录说明

### 核心服务结构：

1. controller/

2. service/：业务逻辑封装

3. repository/：crud操作

4. model/: 存放实体类定义（数据库依赖于这些定义）

5. proto/：网络通信结构

### 其他：
1. config/：redis相关初始化，mysql相关初始化，服务器地址端口设置

2. middleware/：目前只有一个拦截器

<br>

## 注意事项：
1. 需在config中设置正确的mysql数据库端口，密码，用户名
2. 设置正确的redis数据库端口，密码，用户名

<br>

## TODO
### 未解决问题:
1. 没有生成视频封面，也没有设置封面正确的路径
2. 暂无固定后端域名orIP，所以运行程序时需要手动输入服务端地址和端口，以获取正确的视频路径

### 优化方向

- [ ] 1. 优化数据存储   

    - 本想用mongodb来存视频信息，这样似乎比较适合现有的api的要求（返回Video时不必有过多额外操作就能带全Author信息），但想了想还是不太合适，现实情况下，Author信息理应是能够随时更新的，用mongodb的话反而让更新变得麻烦。本质是视频信息也依然是关系型的。在找到更好的方案前继续用mysql。此外，gorm提供[预加载](https://gorm.io/zh_CN/docs/preload.html)，算是能弥补一些不便。

- [ ] 2. 优化依赖注入（别放router.go里）

- [ ] 3. 找个好用的log库，优化debug print和error output

- [ ] 4. 考虑是否有并发风险

- [ ] 5. token相关的东西应该单独出来还是放在service层？controller应该负责construct response吗？

- [ ] 6. 设置token过期时间
