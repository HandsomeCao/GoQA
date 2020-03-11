## Go问答Demo
### 简介
一个使用Go实现的简单问答项目。
其中问题检索和推荐使用基础的BOW(Bag of words)求余弦相似度。
后台框架使用的`gin`

### 准备数据
将数据存储以下格式保存为'data.json'，建议换成数据库存储
```json
{
	"question1": "answer1",
	"question2": "answer2",
	........
}
```

### Go环境运行
1. 安装go 1.13
2. 安装g++, gcc(`gojieba`需要)
3. 运行`go run main.go`，会根据`go.mod`自动安装依赖

### Docker运行

1. Docker创建镜像
```
docker build -t xgfy:v1 .
```

2. 创建容器
```
docker run --name="xgfy_app" -p 8080:8080 xgfy:v1
```

### 访问
```
localhost:8080
```

