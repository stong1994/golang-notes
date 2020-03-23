# 读取配置文件的仓库

1. ```go
   "gopkg.in/yaml.v2"
   ```

   支持yaml文件的读取

2. ```go
   "github.com/spf13/viper"
   ```

   1. 支持多种文件格式: JSON,TOML,YAML,HCL...
   2. 监听配置变化
   3. 支持对环境变量的自动绑定

3. ```go
   gopkg.in/ini.v1
   ```

   1. 支持注释读写操作
   2. 内容到结构体的双向映射
   
   
参考:[ini-设计讲解](https://docs.google.com/presentation/d/1ruDCybd9v7oWBx3WoIC-G26PjfnOdu3GwfSphED8UlM/edit#slide=id.g7b6b4008ad_0_76)