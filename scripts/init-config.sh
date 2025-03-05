#!/bin/bash

# 设置工作目录
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# 创建配置文件函数
create_config() {
    local config_file=$1
    cat > "$config_file" << 'EOF'
# Aquila 服务配置文件

# 应用基础配置
app:
  env: local      # 运行环境：local, dev, prod
  port: 9090      # 服务端口
  app_name: aquila # 应用名称
  app_url: ''     # 应用URL
  db_type: pgsql  # 数据库类型：mysql, pgsql

# JWT认证配置
jwt:
  signing-key: fdc9e9a7-6db6-4e53-bd22-f3410065fe17  # JWT签名密钥
  expires-time: 7d    # Token过期时间
  buffer-time: 1d     # Token缓冲时间
  issuer: qmPlus      # Token签发者

# MySQL数据库配置
mysql:
  host: 127.0.0.1     # 数据库主机地址
  port: "3306"        # 数据库端口
  config: charset=utf8mb4&parseTime=True&loc=Local  # 数据库连接参数
  db_name: aquila     # 数据库名
  username: root      # 数据库用户名
  password: "123456"  # 数据库密码
  prefix: "t_"        # 表前缀
  singular: false     # 是否单数表名
  engine: ""          # 数据库引擎
  max_idle_conns: 10  # 最大空闲连接数
  max_open_conns: 100 # 最大打开连接数
  log_mode: error     # 日志模式
  log_zap: false      # 是否启用zap日志
EOF

    # 根据不同的配置文件添加不同的数据库配置
    if [[ "$config_file" == *"release"* ]]; then
        cat >> "$config_file" << 'EOF'

# PostgreSQL数据库配置
pgsql:
  prefix: ""          # 表前缀
  host: 127.0.0.1     # 数据库主机地址
  port: 15791         # 数据库端口
  config: sslmode=disable TimeZone=Asia/Shanghai  # 数据库连接参数
  db_name: aquila     # 数据库名
  username: postgres  # 数据库用户名
  password: postgresJmx1122@  # 数据库密码
  path: 127.0.0.1     # 数据库路径
EOF
    else
        cat >> "$config_file" << 'EOF'

# PostgreSQL数据库配置
pgsql:
  prefix: ""          # 表前缀
  host: 127.0.0.1     # 数据库主机地址
  port: "5432"        # 数据库端口
  config: sslmode=disable TimeZone=Asia/Shanghai  # 数据库连接参数
  db_name: aquila     # 数据库名
  username: postgres  # 数据库用户名
  password: pg123456  # 数据库密码
  path: 127.0.0.1     # 数据库路径
EOF
    fi

    # 添加剩余的通用配置
    cat >> "$config_file" << 'EOF'
  engine: ""          # 数据库引擎
  log_mode: error     # 日志模式
  max_idle_conns: 10  # 最大空闲连接数
  max_open_conns: 100 # 最大打开连接数
  singular: false     # 是否单数表名
  log-zap: false      # 是否启用zap日志

# 七牛云存储配置
qiniu:
  zone: ZoneHuaDong   # 存储区域
  bucket: ""          # 存储桶名称
  img-path: ""        # 图片路径
  access-key: ""      # 访问密钥
  secret-key: ""      # 秘密密钥
  use-https: false    # 是否使用HTTPS
  use-cdn-domains: false  # 是否使用CDN

# Redis配置
redis:
  addr: 127.0.0.1:6379  # Redis地址
  password: ""          # Redis密码
  db: 0                 # 数据库索引
  useCluster: false     # 是否使用集群

# 系统配置
system:
  db-type: mysql        # 数据库类型
  oss-type: local      # 对象存储类型
  router-prefix: ""     # 路由前缀
  addr: 8888           # 系统地址
  iplimit-count: 15000 # IP限制计数
  iplimit-time: 3600   # IP限制时间
  use-multipoint: false # 是否启用多点登录
  use-redis: false     # 是否使用Redis

# Zap日志配置
zap:
  level: info          # 日志级别
  prefix: '[aquila/server]'  # 日志前缀
  format: console      # 日志格式
  director: log        # 日志目录
  encode_level: LowercaseColorLevelEncoder  # 日志级别编码
  stacktrace_key: stacktrace  # 堆栈跟踪键
  max_age: 0          # 最大保留天数
  show_line: true     # 显示行号
  log_in_console: true  # 控制台输出
EOF

    # 只在非release配置中添加minio配置
    if [[ "$config_file" != *"release"* ]]; then
        cat >> "$config_file" << 'EOF'

# MinIO对象存储配置
minio:
  endpoint:           # MinIO服务端点
  access-key:         # 访问密钥
  secret-key:         # 秘密密钥
  use-ssl: false     # 是否使用SSL
  bucket:            # 存储桶名称
  location:          # 区域位置
  use-cdn-domains: false  # 是否使用CDN
  use-local: false   # 是否使用本地存储
  use-https: false   # 是否使用HTTPS
EOF
    fi
}

# 创建配置文件
echo "Creating config files..."
create_config "${PROJECT_ROOT}/config.yaml"
create_config "${PROJECT_ROOT}/config.release.yaml"

# 设置权限
chmod 644 "${PROJECT_ROOT}/config.yaml"
chmod 644 "${PROJECT_ROOT}/config.release.yaml"

echo "Configuration files have been created successfully!"
