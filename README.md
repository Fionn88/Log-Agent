# Log-Agent

## 正在處理的 issue 

- 無法送訊息到Kafka，創立Topic 可以
1. kafka 和 zookeeper 是用 Docker Run 架設起來，以下為 docker run command
```
docker network create app-tier --driver bridge

docker run -d --name zookeeper --network app-tier -p 2181:2181 jplock/zookeeper

docker run -d --name kafka --link zookeeper --network app-tier -p 7203:7203 -p 9092:9092 -e KAFKA_ADVERTISED_HOST_NAME=zookeeper -e ZOOKEEPER_IP=zookeeper ches/kafka
```

2. go run main.go 接著 input 內容到 my.log 
```
WARN[0007] send message failed,kafka server: In the middle of a leadership election, there is currently no leader for this partition and hence it is unavailable for writes.
```

3. kafka 的 log
```
[2023-09-21 06:51:34,199] INFO Topic creation {"version":1,"partitions":{"0":[0]}} (kafka.admin.AdminUtils$)
[2023-09-21 06:51:34,213] INFO [KafkaApi-0] Auto creation of topic web_log with 1 partitions and replication factor 1 is successful (kafka.server.KafkaApis)
```

## 已解決的 issue
- [Ref](https://stackoverflow.com/questions/76143322/golang-shadowing-variable)以為是傳送 Struct 發生錯誤，是 Variable shadowing 的問題。
```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x10 pc=0x12532f6]

goroutine 1 [running]:
main.run()
        /Users/FionnKuo/Documents/Developer/git/github/Log-Agent/main.go:30 +0x36
main.main()
        /Users/FionnKuo/Documents/Developer/git/github/Log-Agent/main.go:74 +0x28a
exit status 2
```

## 補充知識

- [Ref](https://ithelp.ithome.com.tw/m/articles/10187265) 變數的可視範圍，字母大寫會被exported，小寫則反之