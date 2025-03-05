# 设置工作目录
$scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
$PROJECT_ROOT = (Get-Item $scriptPath).Parent.FullName

# 创建配置文件函数
function Create-Config {
    param (
        [string]$configFile
    )
    
    # 确保目标目录存在
    $directory = Split-Path -Parent $configFile
    if (!(Test-Path -Path $directory)) {
        New-Item -ItemType Directory -Path $directory | Out-Null
    }

    # 使用 UTF8 编码创建文件
    $Utf8NoBomEncoding = New-Object System.Text.UTF8Encoding $False
    
    # 创建配置内容（与bash脚本中的内容相同，只是去掉了EOF标记）
    $configContent = @"
# Aquila 服务配置文件

# ...existing code...
"@

    # 根据是否为release版本添加不同的PostgreSQL配置
    if ($configFile -like "*release*") {
        $pgsqlConfig = @"
# PostgreSQL数据库配置（生产环境）
pgsql:
  prefix: ""          # 表前缀
  host: 127.0.0.1     # 数据库主机地址
  port: 15791         # 数据库端口
  config: sslmode=disable TimeZone=Asia/Shanghai  # 数据库连接参数
  db_name: aquila     # 数据库名
  username: postgres  # 数据库用户名
  password: postgresJmx1122@  # 数据库密码
  path: 127.0.0.1     # 数据库路径
"@
    } else {
        $pgsqlConfig = @"
# PostgreSQL数据库配置（开发环境）
pgsql:
  prefix: ""          # 表前缀
  host: 127.0.0.1     # 数据库主机地址
  port: "5432"        # 数据库端口
  config: sslmode=disable TimeZone=Asia/Shanghai  # 数据库连接参数
  db_name: aquila     # 数据库名
  username: postgres  # 数据库用户名
  password: pg123456  # 数据库密码
  path: 127.0.0.1     # 数据库路径
"@
    }

    # 合并所有配置内容
    $fullConfig = $configContent + "`n" + $pgsqlConfig

    # 添加通用配置
    $commonConfig = @"
# ...existing code...
"@
    
    $fullConfig = $fullConfig + "`n" + $commonConfig

    # 如果不是release版本，添加MinIO配置
    if ($configFile -notlike "*release*") {
        $minioConfig = @"
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
"@
        $fullConfig = $fullConfig + "`n" + $minioConfig
    }

    # 写入文件
    [System.IO.File]::WriteAllLines($configFile, $fullConfig, $Utf8NoBomEncoding)
}

# 执行配置文件创建
Write-Host "Creating config files..."
Create-Config (Join-Path $PROJECT_ROOT "config.yaml")
Create-Config (Join-Path $PROJECT_ROOT "config.release.yaml")
Write-Host "Configuration files have been created successfully!"
