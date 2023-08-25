[toc]

# gin 实现增删改查

# 启动

```sh
go run .
```

# 测试接口 (烂尾，未写完)

```sh
# 首页
curl http://localhost:8080

# 用户登录接口
curl -X POST -d '{"username":"root","password":"admin"}'http://localhost:8080/loginJSON
curl -X POST --data-urlencode 'username=root' --data-urlencode 'password=admin' 'http://localhost:8080/loginForm'
# curl -X http://localhost:8080/loginJSON?username=root&password=admin


# 获取所有动物园
curl http://localhost:8080/zoos
# 新增动物园
curl -X POST -d ``  http://localhost:8080/zoos
```
