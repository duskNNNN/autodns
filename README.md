# autodns
> 适用于家庭带宽 

自动获取本地ipv6地址，并修改对应腾讯云域名AAAA类型的记录值

## 配置文件
> 创建Conf/conf.yml

```shell
# 填写腾讯云API密钥
dnspod:
  id: 
  token: 
# 填写邮箱信息
mail:
  adminEmail:  # 此处为接收者邮箱
  host: 
  port: 
  sendEmail: # 此处为发送者邮箱
  token:  # 发送者的密钥
```

## PS
代码写的过于垃圾，大佬不喜轻喷