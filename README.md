<h1 align="center">《一键生成Go-orm结构体模型工具》V1.01</h1>

## 环境要求

* Go版本 >= 1.13
* Go Module 模式

## 安装步骤

*  go get -u github.com/phpwei/go-Gentity
* 项目根目录执行命令 go mod tidy
* go环境安装 可以参考 https://learnku.com/articles/38410


## 工具命令汇总

* 【一键生成全部表】 go run plustool/main.go db -m  all


## 注意事项

* 生成的model结构体默认在根目录 app/models/ 下面
* 数据库配置文件在根目录下 application.yaml
