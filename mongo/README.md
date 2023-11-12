## docker-compose创建mongo集群

docker-compose.yaml
```yaml
version: '3.8'

networks:
  common.network:
    driver: bridge

services:
  mongo1:
      image: mongo:5.0.18
      restart: always
      networks:
        common.network:
      ports:
        - 27017:27017
      volumes:
        - ./data/mongo1:/data/db
      entrypoint: [ "/usr/bin/mongod", "--bind_ip_all", "--replSet", "rs0" ]

  mongo2:
    image: mongo:5.0.18
    restart: always
    networks:
      common.network:
    ports:
      - 27018:27017
    volumes:
      - ./data/mongo2:/data/db
    entrypoint: [ "/usr/bin/mongod", "--bind_ip_all", "--replSet", "rs0" ]

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
设置集群
```sh
rs.initiate({_id:'rs0', members: [{_id:0, host: 'mongo1'},{_id:1, host: 'mongo2'},{_id:2, host: 'mongo3'}]})
```

设置hosts
```sh
sudo vim /etc/hosts
```

粘贴以下内容
```sh
127.0.0.1       mongo1
127.0.0.1       mongo2
```

### 其他

重新设置副本集
```sh
rs.reconfig({_id:'rs0', members: [{_id:0, host: 'mongo1'},{_id:1, host: 'mongo2'},{_id:2, host: 'mongo3'}]}, {newForce:true})
```