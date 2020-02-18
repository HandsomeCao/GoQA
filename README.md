### Docker创建镜像
```
docker build -t xgfy:v1 .
```

### 创建容器
```
docker run --name="xgfy_app" -p 8080:8080 xgfy:v1
```

### 访问
```
localhost:8080
```
