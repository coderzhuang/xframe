一个简单的，基于gin的，开箱即用的api框架

框架分层：

* 接入层
    * 对应目录为 access，接入层是变化层，目前支持grpc、http
    * 请求转发到服务层或接收服务层返回，会进行对象转换
* 服务层
    * 对应目录为 service，服务层实现具体的业务逻辑
    * 服务层被接入层和基建层依赖
* 基建层
    * 对应目录为 infrastructure，这里维护着底层服务，比如：mysql。
    * 请求转发到服务层或接收服务层返回，会进行对象转换
    * 为保证服务层的稳定，基建层需要实现对应的接口
    
todo
- [x] 平滑重启
- [x] 配置加载
- [x] 集成swag
- [x] debug模式
- [x] 数据库链接数
- [x] 集成job
- [x] docker部署
- [x] 完善启动过程
- [x] 普罗米修斯 
  
  依赖 https://github.com/chenjiandongx/ginprom
  
  grafana Go Metrics ID: 10826
  
- [x] 完善追踪

