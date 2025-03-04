# Gomall - 抖音商城

![GitHub](https://img.shields.io/github/license/weeadd/Gomall)
![Static Badge](https://img.shields.io/badge/collaborator-4-lightblue)
![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)
![GitHub Repo stars](https://img.shields.io/github/stars/weeadd/Gomall)

## ✨ 简介

Gomall 是一个基于 Go 语言开发的抖音商城项目，使用 Kitex 和 Hertz 框架构建。项目采用微服务架构，涵盖了用户管理、商品展示、购物车、订单处理、支付等多个核心功能模块。本项目根据字节跳动青训营（2025春）官方要求和模板制作。

## 🛠 技术栈
<div align="center">

| [Kitex](https://www.cloudwego.io/zh/docs/kitex/) | [Hertz](https://www.cloudwego.io/zh/docs/hertz/) | [GORM](https://gorm.io/) |
|:---:|:--:|:--:|
| [<img src="https://avatars.githubusercontent.com/u/79236453?s=200&v=4" alt="kitex" height="100px"/>](https://www.cloudwego.io/zh/docs/kitex/) | [<img src="https://avatars.githubusercontent.com/u/79236453?s=200&v=4" alt="hertz" height="100px"/>](https://www.cloudwego.io/zh/docs/hertz/) | [<img src="https://gorm.io/favicon.ico" alt="gorm" height="100px"/>](https://gorm.io/) |

| [MySQL](https://www.mysql.com/) | [Consul](https://www.consul.io/) | [Prometheus](https://prometheus.io/) |
|:---:|:--:|:--:|
| [<img src="https://www.mysql.com/common/logos/logo-mysql-170x115.png" alt="mysql" height="100px" width="100px"/>](https://www.mysql.com/) | [<img src="https://www.consul.io/_next/static/media/consul-community_on-light.e8e68a5f.svg" alt="consul" height="100px" width="100px"/>](https://www.consul.io/) | [<img src="https://prometheus.io/assets/prometheus_logo_grey.svg" alt="prometheus" height="100px"/>](https://prometheus.io/) |

</div>

## 🗄️ 数据库

本项目使用 `MySQL` 作为主要的关系型数据库，用于存储用户信息、商品信息、订单信息等结构化数据。

## 🚀 运行

在执行后续步骤之前，请确保您已配置好以下环境：

- Go ≥ 1.24
- MySQL ≥ 8.0
- Docker
  
### 1. 克隆项目
首先，克隆本仓库到本地：
```bash
git clone https://github.com/weeadd/Gomall.git
cd gomall
```

### 2. 配置数据库
每个微服务都支持通过 `.env` 文件配置 MySQL 连接信息。在每个微服务的目录下（如 `app/product`），创建 `.env` 文件并填写以下内容：
```env
MYSQL_USER=root
MYSQL_PASSWORD=your_password
MYSQL_HOST=your_host_ip
MYSQL_DATABASE=database_name
```
请根据您的 MySQL 配置修改上述值。

### 3. 启动依赖服务
使用 Docker 快速启动 Consul（服务发现与配置管理）和 MySQL：
```bash
docker-compose -f docker/docker-compose.yaml up -d
```

### 4. 启动微服务
在项目根目录下运行以下命令，启动所有微服务：
```bash
make run
```
或者，您可以进入每个微服务目录（如 `app/product`），手动运行：
```bash
go run main.go
```

### 5. 访问服务
- API 服务默认运行在 `8080` 端口，您可以通过 `http://localhost:8080` 结合路由访问。（如 获取商品信息 `localhost:8080/product/get?id=1`）
- 使用 Prometheus 和 Grafana 监控服务状态（默认端口：9090 和 3000）。

---

通过以上步骤，您可以快速启动并体验 Gomall 项目。如果有任何问题，请参考项目的 [Issues](https://github.com/weeadd/Gomall/issues) 或提交新的问题。

## ⚙️ 架构设计

Gomall 采用微服务架构，基于 Kitex 和 Hertz 框架构建，整体架构分为三层：

### 1. **API 网关层**
   - 使用 **Hertz** 作为 HTTP 服务器，负责接收用户请求并进行路由转发。
   - 通过 RPC 调用后端微服务，完成业务逻辑处理。
   - 提供统一的 API 入口，支持用户认证、请求限流等功能。

### 2. **微服务层**
   - 核心业务拆分为多个独立的微服务，每个服务专注于单一职责：
     - **User**：用户管理（注册、登录、信息维护）。
     - **Auth**：身份认证与权限管理（JWT 鉴权）。
     - **Product**：商品管理（商品展示、库存管理）。
     - **Cart**：购物车管理（商品添加、删除、修改）。
     - **Order**：订单管理（订单创建、状态更新）。
     - **Payment**：支付服务（支付流程处理）。
     - **Checkout**：结算服务（整合订单、支付、库存等逻辑）。
   - 微服务之间通过 **Kitex** 进行 RPC 通信，保证高性能和低延迟。

### 3. **数据存储与基础设施层**
   - **MySQL**：存储用户、商品、订单等核心业务数据。
   - **Consul**：提供服务注册与发现、配置管理功能，支持动态扩展。
   - **Prometheus + Grafana**：监控系统性能，实时采集并展示服务指标（如请求延迟、错误率等）。
   - **Docker**：容器化部署，确保环境一致性。

### 架构特点
- **高可用性**：通过 Consul 实现服务动态注册与发现，支持弹性扩展。
- **可观测性**：集成 Prometheus 监控系统，实时跟踪服务状态。
- **松耦合**：微服务独立部署，便于维护和扩展。

---

该架构设计确保了 Gomall 的高性能、可扩展性和易维护性，能够灵活应对业务需求的变化。架构方案来源于字节跳动CloudWeGo (https://github.com/cloudwego) 官方biz-demo，项目团队根据实际需求和工期进行了部分裁剪和调整，未来将进一步加入更多组件与功能。

## 🤝 团队成员

本项目由以下开发者共同完成（顺序不分先后）：

- [Douglas](https://github.com/108474839)
- [weeadd](https://github.com/weeadd)
- [XYsiren](https://github.com/XYsiren)
- [ysqfirmament](https://github.com/ysqfirmament)

## ✍️ 写在最后

项目制作不易，如果它对你有帮助的话，请务必给作者点一个免费的⭐，万分感谢!🙏🙏🙏
