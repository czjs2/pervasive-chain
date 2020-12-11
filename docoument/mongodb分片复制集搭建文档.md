### mongodb分片集群搭建
官方文档: https://docs.mongodb.com/manual/tutorial/deploy-shard-cluster/

#### 分片集群复制集搭建搭建
搭建一个分片复制集集群示例

编写配置文件,多个实例类似

    # mongod.conf
    # where to write logging data.
    # Where and how to store data.
    storage:
    dbPath: /mongodb/shard/rs1/mongodb27017/data
    journal:
        enabled: true
    #  engine:
    #  wiredTiger:

    # how the process runs
    processManagement:
      fork: true  # fork and run in background
      pidFilePath: /mongodb/shard/rs1/mongodb27017/mongod.pid  # location of pidfile
      timeZoneInfo: /usr/share/zoneinfo

    # network interfaces
    net:
      port: 27017
      bindIp: 0.0.0.0
    #Enter 0.0.0.0,:: to bind to all IPv4 and IPv6 addresses or, alternatively, use the net.bindIpAll setting.


    security:
      keyFile: /mongodb/etc/keyFile.file
      authorization: enabled
    #operationProfiling:

    replication:
      replSetName: shard01

    sharding:
      clusterRole: shardsvr
    ## Enterprise-Only Options

    ---------------------------------------------

    // 启动分片集合
    mongod -f /mongodb/shard/rs1/mongodb27017/mongod.conf
    mongod -f /mongodb/shard/rs1/mongodb27018/mongod.conf
    mongod -f /mongodb/shard/rs1/mongodb27019/mongod.conf


    // 连接 mongo shell 
     mongo --port 27017


    // 初始化复制集
    rs.initiate(
    {
        _id : "shard01",
        members: [
        { _id : 0, host : "118.24.168.230:27017" },
        { _id : 1, host : "118.24.168.230:27018" },
        { _id : 2, host : "118.24.168.230:27019" }
        ]
    }
    )

 
    // 创建管理用户
    use admin
    db.createUser(
        {
            user:"root",
            pwd:"xjrw2020",
            roles:["root"]
        }
    )   

    


#### 配置服务复制集搭建
配置文件

    # mongod.conf

    # for documentation of all options, see:
    #   http://docs.mongodb.org/manual/reference/configuration-options/

    # where to write logging data.
    systemLog:
      destination: file
      logAppend: true
      path: /mongodb//configs/configs27025/mongod.log

    # Where and how to store data.
    storage:
      dbPath: /mongodb/configs/configs27025/data
    journal:
        enabled: true
    #  engine:
    #  wiredTiger:

    # how the process runs
    processManagement:
      fork: true  # fork and run in background
      pidFilePath: /mongodb/configs/configs27025/mongod.pid  # location of pidfile
      timeZoneInfo: /usr/share/zoneinfo

    # network interfaces
    net:
      port: 27025
      bindIp: 0.0.0.0
    #Enter 0.0.0.0,:: to bind to all IPv4 and IPv6 addresses or, alternatively, use the net.bindIpAll setting.


    security:
      keyFile: /mongodb/etec/keyFile.file
      authorization: enabled
    #operationProfiling:
    sharding:
      clusterRole: configsvr

    replication:
      replSetName: configs
    #sharding:

    ## Enterprise-Only Options

    // 连接mongo shell,
    mongo --port 27023


    // 初始化复制集
    rs.initiate(
    {
        _id: "configs",
        configsvr: true,
        members: [
        { _id : 0, host : "118.24.168.230:27023" },
        { _id : 1, host : "118.24.168.230:27024" },
        { _id : 2, host : "118.24.168.230:27025" }
        ]
    }
    )

    // 创建管理用户
    use admin
    db.createUser(
        {
            user:"configroot",
            pwd:"xjrw2020",
            roles:["root"]
        }
    )   


#### mongos 路由服务
配置文件

    # where to write logging data.
    systemLog:
      destination: file
      logAppend: true
      path: /mongodb/mongos/mongod.log

    # Where and how to store data.
    #storage:
    #  dbPath: /mongodb/shard/rs1/mongodb27019/data
    #  journal:
    #    enabled: true
    #  engine:
    #  wiredTiger:

    # how the process runs
    processManagement:
      fork: true  # fork and run in background
      pidFilePath: /mongodb/mongos/mongod.pid  # location of pidfile
      timeZoneInfo: /usr/share/zoneinfo

    # network interfaces
    net:
      port: 27026
      bindIp: 0.0.0.0
    #Enter 0.0.0.0,:: to bind to all IPv4 and IPv6 addresses or, alternatively, use the net.bindIpAll setting.

    #security:
    #  keyFile: /mongodb/etec/keyFile.file
    #  authorization: enabled
    #operationProfiling:

    #replication:
    #  replSetName: myset
    sharding:
      configDB: configs/127.0.0.1:27023,127.0.0.1:27024,127.0.0.1:27025
    ## Enterprise-Only Options

    #auditLog:

    // 启动
    mongos -f   /mongodb/mongos/mongod.cnf  

    // 连接 shell
    mongo --host 127.0.0.1 --port 27026

    // 验证，这个配置服务器设置的用户
    use admin
    db.auth("root","xjrw2020")

    

    // 添加分片集群一
    sh.addShard( "shard01/118.24.168.230:27017,118.24.168.230:27018,118.24.168.230:27019")
    // 分片集群二
    sh.addShard( "shard02/139.186.84.15:27017,139.186.84.15:27018,139.186.84.15:27019")

    // 创建通过 mongos 连接的用户 
    use pynxtest
    db.createUser({
        user:"pynxtest",
        pwd:"xjrw2020",
        customData:{
            name:'xjrw',
        },
        roles:[
            {role:"readWrite",db:"pynxtest"},
        ]
    })  

     // 确定那个数据库启用分片 
    sh.enableSharding("pynxtest")

    // 设置分片键 hash和等差两种
    sh.shardCollection("pynxtest.block", { height : 1} )

    // mongod url
    mongodb://pynxtest:xjrw2020@127.0.0.127026/pynxtest?authSource=pynxtest








参考: 
1. https://www.jianshu.com/p/94dd3c6b2cbb
2. 官方: https://docs.mongodb.com/manual/tutorial/deploy-shard-cluster/

