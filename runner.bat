@echo off
setlocal enabledelayedexpansion

:: 设置默认参数值
set "env=dev"
set "sys=win64"
set "show_help=0"

:: 解析命令行参数
:parse_args
if "%1"=="" goto end_parse
if "%1"=="--env" (
    set "env=%2"
    shift
    shift
    goto parse_args
)
if "%1"=="--sys" (
    set "sys=%2"
    shift
    shift
    goto parse_args
)
if "%1"=="--help" (
    set "show_help=1"
    shift
    goto parse_args
)
shift
goto parse_args

:end_parse

:: 显示帮助信息
if "%show_help%"=="1" (
    echo 用法: %0 [OPTIONS]
    echo.
    echo OPTIONS:
    echo   --env dev^|prod    设置运行环境，默认为dev
    echo   --sys win64^|linux 设置目标系统，默认为win64
    echo   --help             显示此帮助信息
    echo.
    echo 示例:
    echo   %0 --env dev --sys win64
    echo   %0 --env prod --sys linux
    echo   %0 --help
    goto :eof
)

:: 参数验证
if not "%env%"=="dev" if not "%env%"=="prod" (
    echo 错误: env参数必须是dev或prod
    goto :eof
)

if not "%sys%"=="win64" if not "%sys%"=="linux" (
    echo 错误: sys参数必须是win64或linux
    goto :eof
)

echo 当前配置:
echo   环境: %env%
echo   系统: %sys%
echo.

:: 开发环境处理
if "%env%"=="dev" (
    echo 检查GoFrame依赖...
    
    :: 检查go.mod中是否包含goframe
    findstr /i "goframe" go.mod >nul 2>&1
    if %errorlevel% equ 0 (
        echo 检测到GoFrame依赖，使用gf run...
        gf run main.go
    ) else (
        echo 未检测到GoFrame依赖，使用go run...
        go run main.go
    )
    goto :eof
)

:: 生产环境处理 - 打包
if "%env%"=="prod" (
    echo 开始构建生产环境包...
    
    if "%sys%"=="win64" (
        echo 目标系统: Windows 64位
        set "CGO_ENABLED=0"
        set "GOOS=windows"
        set "GOARCH=amd64"
        echo 设置环境变量: CGO_ENABLED=0, GOOS=windows, GOARCH=amd64
        echo.
        echo 执行: go build
        set CGO_ENABLED=0
        set GOOS=windows
        set GOARCH=amd64
        go build
        if %errorlevel% equ 0 (
            echo 构建成功！生成文件: app.exe
        ) else (
            echo 构建失败！
        )
    )
    
    if "%sys%"=="linux" (
        echo 目标系统: Linux
        set "CGO_ENABLED=0"
        set "GOOS=linux"
        set "GOARCH=amd64"
        echo 设置环境变量: CGO_ENABLED=0, GOOS=linux, GOARCH=amd64
        echo.
        echo 执行: go build
        set CGO_ENABLED=0
        set GOOS=linux
        set GOARCH=amd64
        go build
        if %errorlevel% equ 0 (
            echo 构建成功！生成文件: app
        ) else (
            echo 构建失败！
        )
    )
    goto :eof
)

endlocal