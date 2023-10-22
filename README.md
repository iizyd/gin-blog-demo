## 项目简介

用 `Go(Gin)` 实现了一个简单的博客后台 demo，实现了登录、标签管理、文章管理、文件上传等功能

对应的前端项目地址 https://github.com/iizyd/express-blog ，项目内含后台管理前端和博客前端，同时还有另一个版本的后端，用 `Express.js` 实现，接口一致

## 技术栈

- Web框架 `gin`
- 参数校验 `govalidator` 
- 数据库 `gorm`、`mysql`
- 鉴权 `jwt-go` 
- 配置读取 `yaml.v3`
- 日志 `zap`、`lumberjack`

## 开发计划

- [x] 日志记录
- [x] 配置读取
- [x] 参数校验
- [x] 文章管理
- [x] 文章管理
- [x] 标签管理
- [x] 图片上传
- [x] 登录
- [x] 博客前端接口

## 目录结构

|  文件夹 | 描述  |
|  :----:  | :----: |
| backend-go  | 后端代码 |
| v1.0.0-version | 旧版本 `go` 和 `nestjs` 代码 |

## 截图
![blog-1](https://raw.githubusercontent.com/iizyd/express-blog/main/pic/blog-1.png)

![blog-2](https://raw.githubusercontent.com/iizyd/express-blog/main/pic/blog-2.png)

![blog-3](https://raw.githubusercontent.com/iizyd/express-blog/main/pic/blog-3.png)

![frontend-1](https://raw.githubusercontent.com/iizyd/express-blog/main/pic/frontend-1.png)

![frontend-2](https://raw.githubusercontent.com/iizyd/express-blog/main/pic/frontend-2.png)

![frontend-3](https://raw.githubusercontent.com/iizyd/express-blog/main/pic/frontend-3.png)

![frontend-4](https://raw.githubusercontent.com/iizyd/express-blog/main/pic/frontend-4.png)
