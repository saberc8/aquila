# aquila

## 项目介绍
Aquila 是一个基于 Go 语言的后台管理系统，提供了一系列 API 接口用于管理和操作数据。


## 打包
GOOS=linux GOARCH=amd64 go build -o myapp main.go

## 目录结构
```
aquila
├── README.md
├── api // api接口
├── config // 配置文件
├── constants // 常量
├── enum // 枚举
├── global // 全局变量
├── initialize // 初始化
├── log
├── middleware // 中间件
├── model // 数据库模型
├── router // 路由
├── utils // 工具
└── main.go
```

## 配置文件说明

### 配置文件初始化

项目提供了跨平台的配置文件初始化脚本，支持 Windows、Linux 和 macOS 系统。

#### Windows 系统
两种方式可以初始化配置：
1. 双击运行 `scripts/init-config.bat`
2. 在 PowerShell 中执行：
```powershell
cd scripts
Set-ExecutionPolicy Bypass -Scope Process -Force
./init-config.ps1
```

#### Linux/macOS 系统
```bash
cd scripts
chmod +x init-config.sh
./init-config.sh
```

### 配置文件结构

项目包含两个主要配置文件：
- `config.yaml`: 开发环境配置
- `config.release.yaml`: 生产环境配置

主要配置项说明：

#### 应用配置 (app)
```yaml
app:
  env: local      # 运行环境：local, dev, prod
  port: 9090      # 服务端口
  app_name: aquila # 应用名称
  db_type: pgsql  # 数据库类型：mysql, pgsql
```

#### 数据库配置
支持 PostgreSQL 和 MySQL 两种数据库：
```yaml
pgsql:
  host: 127.0.0.1     # 数据库主机地址
  port: "5432"        # 数据库端口（开发环境）
  db_name: aquila     # 数据库名
  username: postgres  # 数据库用户名

mysql:
  host: 127.0.0.1     # 数据库主机地址
  port: "3306"        # 数据库端口
  db_name: aquila     # 数据库名
```

#### 对象存储配置
支持 MinIO 和七牛云存储：
```yaml
minio:
  endpoint:           # MinIO服务端点
  access-key:         # 访问密钥
  bucket:            # 存储桶名称

qiniu:
  zone: ZoneHuaDong   # 存储区域
  bucket: ""          # 存储桶名称
```

#### 缓存配置 (Redis)
```yaml
redis:
  addr: 127.0.0.1:6379  # Redis地址
  password: ""          # Redis密码
  db: 0                 # 数据库索引
```

#### 日志配置 (Zap)
```yaml
zap:
  level: info          # 日志级别
  format: console      # 日志格式
  director: log        # 日志目录
```

### 环境差异

1. **开发环境** (`config.yaml`)
   - PostgreSQL 默认端口：5432
   - 包含 MinIO 配置
   - 简化的安全配置

2. **生产环境** (`config.release.yaml`)
   - PostgreSQL 默认端口：15791
   - 不包含 MinIO 配置
   - 增强的安全配置

### 配置文件修改建议

1. 首次使用时，复制示例配置：
   ```bash
   cp config.yaml config.local.yaml
   ```

2. 根据环境修改相应配置：
   - 数据库连接信息
   - Redis 连接信息
   - 存储配置信息

3. 注意事项：
   - 敏感信息（密码、密钥等）建议使用环境变量
   - 生产环境配置文件建议加入 .gitignore
   - 定期检查并更新 JWT 密钥