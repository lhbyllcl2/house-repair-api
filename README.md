#### 常用命令
```
 //生成模型
 goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3310)/gozero_sys" -table="*"  -dir="./model/sysmodel" --style="go_zero"
 
 //创建api
 goctl api go -api admin.api -style go_zero -dir ../
 
 //创建rpc

 goctl rpc proto -src sys.proto -style go_zero -dir . 
```