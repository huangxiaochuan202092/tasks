
# 项目介绍

本项目实现了登录即注册的功能，并在此基础上依次实现了用户功能模块、任务功能模块以及博客功能模块。

## 项目部署流程

### 1. 克隆项目
使用 `git clone` 命令克隆项目到本地：
```bash
git clone <项目地址>
```
请将 `<项目地址>` 替换为实际的项目仓库地址。

### 2. 修改配置
需要修改以下配置项：
- **Redis 和 MySQL 连接配置**：在项目配置文件中更新 Redis 和 MySQL 的连接信息，如地址、端口、用户名、密码等。
- **邮箱及邮箱授权码**：设置用于系统通知、验证等功能的邮箱地址及其授权码。

### 3. 初始化项目
在项目根目录下执行以下命令初始化项目依赖：
```bash
go mod tidy
```

### 4. 运行项目
配置修改完成后，进入 `resul` 目录并执行以下命令启动项目：
```bash
cd resul
go run main.go
```

### 5. 访问项目
项目启动成功后，通过 `IP:Port` 地址访问项目，其中 `IP` 是服务器的 IP 地址，`Port` 是项目配置的端口号。

## 视频资源
### 哔哩哔哩播放
[点击观看](https://space.bilibili.com/3546867629558058?spm_id_from=333.337.0.0)

### 抖音视频播放
[点击观看](https://www.douyin.com/user/MS4wLjABAAAAutuiF-v06OCpXGOjaUDTGT6u4WG4kadCuRbZEvLRY1s?from_tab_name=main)

## 系列视频内容
### 【第 1 期】
Go 语言快速上手 Web 应用项目之登录即注册，Gin 是 Web 开发的涡轮增压引擎，技术要点包括 Gin + Gorm + Redis。

### 【第 2 期】
Go 语言快速上手 Web 应用项目之用户管理模块。

### 【第 3 期】
Go 语言快速上手 Web 应用项目之任务管理模块。

### 【第 4 期】
Go 语言快速上手 Web 应用项目之博客管理模块。

### 【第 5 期】
Go 语言快速上手 Web 应用项目之问卷管理模块。

## English Version

# Project Introduction

This project has implemented a feature where logging in also serves as registration. Based on this, the user function module, the task function module, and the blog function module have been successively implemented.

## Project Deployment Process

### 1. Clone the Project
Use the `git clone` command to clone the project to your local machine:
```bash
git clone <Project Address>
```
Replace `<Project Address>` with the actual address of the project repository.

### 2. Modify the Configuration
You need to modify the following configuration items:
- **Redis and MySQL Connection Configuration**: Update the connection information for Redis and MySQL, such as address, port, username, and password, in the project configuration file.
- **Email and Email Authorization Code**: Set the email address and its authorization code for system notifications, verification, and other functions.

### 3. Initialize the Project
Execute the following command in the project root directory to initialize project dependencies:
```bash
go mod tidy
```

### 4. Run the Project
After modifying the configuration, navigate to the `resul` directory and execute the following command to start the project:
```bash
cd resul
go run main.go
```

### 5. Access the Project
After the project starts successfully, access the project via the `IP:Port` address, where `IP` is the server's IP address and `Port` is the port number configured for the project.

## Video Resources
### Bilibili Playback
[Click to Watch](https://space.bilibili.com/3546867629558058?spm_id_from=333.337.0.0)

### Douyin Video Playback
[Click to Watch](https://www.douyin.com/user/MS4wLjABAAAAutuiF-v06OCpXGOjaUDTGT6u4WG4kadCuRbZEvLRY1s?from_tab_name=main)

## Series Video Content
### [Episode 1]
Quickly get started with the "Login is Registration" feature in the Go web application project. Gin is a turbocharged engine for web development, and the technical points include Gin + Gorm + Redis.

### [Episode 2]
Quickly get started with the user management module in the Go web application project.

### [Episode 3]
Quickly get started with the task management module in the Go web application project.

### [Episode 4]
Quickly get started with the blog management module in the Go web application project.

### [Episode 5]
Quickly get started with the questionnaire management module in the Go web application project. 

<img width="1245" alt="image" src="https://github.com/user-attachments/assets/453ec66d-178a-44b0-9ee2-4afe8785d99c" />
