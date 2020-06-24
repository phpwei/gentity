<h1 align="center">《一键生成Go-orm结构体模型工具》V1.01</h1>

## 环境要求

* Go版本 >= 1.13
* Go Module 模式

## 安装步骤

*  go get -u github.com/phpwei/gentity
* 项目根目录执行命令 go mod tidy
* go环境安装 可以参考 https://learnku.com/articles/38410


## 工具命令汇总

* 【一键生成指定表】 go run plustool/main.go db -m  输入表名
* 【一键生成全部表】 go run plustool/main.go db -m  all
* 【一键生成全部表自动生成 无需提醒】 go run plustool/main.go db -m  all -y

## 注意事项

* 生成的model结构体默认在根目录 app/models/ 下面
* 数据库配置文件在根目录下 application.yaml

## 效果展示

![](https://cdn.learnku.com/uploads/images/202005/21/26539/TM3Q6KwiKq.png!large)

![](https://cdn.learnku.com/uploads/images/202005/21/26539/MuyWi1U3El.png!large)

![](https://cdn.learnku.com/uploads/images/202005/21/26539/Vdh6j75Ruq.png!large)