db:
  driver: mysql
  dsn: root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local
  #生成文件目录
  path: /app/models/
  #设置表前缀,没有可以为空
  prefix: