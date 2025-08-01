@echo off
echo Starting What-to-Wear Server in Development Mode...

REM 设置开发环境变量
set GIN_MODE=debug
set LOG_LEVEL=debug
set LOG_FORMAT=text
set LOG_OUTPUT=stdout
set ENABLE_DETAILED_LOGGING=true
set ENABLE_SQL_LOGGING=true

REM 启动服务器
go run main.go

pause
