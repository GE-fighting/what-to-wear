#!/bin/bash

echo "Starting What-to-Wear Server in Development Mode..."

# 设置开发环境变量
export GIN_MODE=debug
export LOG_LEVEL=debug
export LOG_FORMAT=text
export LOG_OUTPUT=stdout
export ENABLE_DETAILED_LOGGING=true
export ENABLE_SQL_LOGGING=true

# 启动服务器
go run main.go
