#  使用说明
打包命令请执行 linux_build.bat 或 win_build.bat 两个批处理文件默认win平台，如果代码在linux下打包，请参照打包命令便写脚本即可。

打包生成可执行文件后，config.json为配置文件，必须和可执行文件放在同一目录，其中配置说明如下：
```
{
    "port": 16666, #服务启动的端口号
    "initRes":"{\"code\":0,\"msg\":\"Success\",\"data\":\"im ok\"}", #默认返回的格式
    "initPath":"/app"  #初始化的请求路径
}
```
## 特殊说明：本程序默认只支持json数据post方式提交。其它方式需要对源码进行修改适配。