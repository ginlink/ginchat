## docker-compose创建mongo集群

docker-compose.yaml
```yaml
version: '3.7'
services:
  mongo1:
      image: mongo:5.0.18
      restart: always
      ports:
        - 27017:27017
      volumes:
        - ./data/mongo1:/data/db
      command: mongod --replSet "rs0" --bind_ip_all

  mongo2:
    image: mongo:5.0.18
    restart: always
    ports:
      - 27018:27017
    volumes:
      - ./data/mongo2:/data/db
    command: mongod --replSet "rs0" --bind_ip_all
```

建立目录
```sh
mkdir -p data/mongo1 data/mongo2
```


启动
```sh
docker-compose up -d
```

查看运行状态
```sh
docker-compose ps
```

连接mongo1，去设置集群
```sh
docker-compose exec mongo1 bash

mongo
```

执行以下命令来配置复制集，注意ip
```sh
rs.initiate({_id: "rs0", members: [
  {_id: 0, host: "192.168.1.2:27017"},
  {_id: 1, host: "192.168.1.2:27018"},
]});
```

链接uri
```sh
mongodb://192.168.1.2:27017,192.168.1.2:27018/test?replicaSet=rs0
```
