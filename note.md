1. go build 命令生成.exe文件：必须要有main 包和main方法
go build halo.go

go run halo.go: 执行/运行go文件

2. import : 不能导入没有用到的包
导入的包邮多个的时候，会按照排列的顺序导入；
如果导入的包里面依赖其他包，会先导入他依赖的包；
先导入依赖包，在初始化main包里面的变量，最后执行main函数。

3. 不能自动导入自定义的包？
要把project gopath的路径设置为项目的路径:H:\gopro\go-study
