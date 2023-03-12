# gin-curd

### 数据库初始化(docker)
本地的3309端口挂载一个名为mysql，root用户名密码为123456的MySQL容器环境:
```
docker run --name mysql -p 3309:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:8.0.19
--name: 容器名称
-e    : 设置环境
-d    : 守护线程，后台运行
```



另外启动一个MySQL Client连接上面的MySQL环境，密码为上一步指定的密码123456:
```
docker run -it --network host --rm mysql mysql -h127.0.0.1 -P3309 --default-character-set=utf8mb4 -uroot -p
```
注： --rm参数会使容器退出时删除容器
如需多次使用，请去除 --rm
进入容器命令
```
docker exec -it <容器ID> mysql -h127.0.0.1 -P3309 --default-character-set=utf8mb4 -uroot -p
```

