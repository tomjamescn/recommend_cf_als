# recommend_cf_als
推荐系统_协同过滤_als算法

# 使用方法
## 安装
```sh
go get -v github.com/tomjamescn/recommend_cf_als/...
```

## 解压预先训练的模型
> 此模型精度不高，作为演示够用了

```sh
cd $GOPATH/src/github.com/tomjamescn/recommend_cf_als/als
tar xvf model.tar.gz
```

## 运行服务
```sh
cd $GOPATH/src/github.com/tomjamescn/recommend_cf_als/service
go run main
```

## 访问服务
http://127.0.0.1:9999/recommend_cf_asl?movie_id=111146
