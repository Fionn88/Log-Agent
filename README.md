# Log-Agent

## 正在處理的 issue 

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

- [Ref](https://github.com/Fionn88/Log-Agent/issues/1) 無法送訊息到 Kafka，創立 Topic 可以
   - 換一個 Docker Source 來源即可

## 補充知識

- [Ref](https://ithelp.ithome.com.tw/m/articles/10187265) 變數的可視範圍，字母大寫會被exported，小寫則反之
- Run Kafka Consumer
  - `docker run --rm --network <KAFKA_SAME_NETWORK> bitnami/kafka kafka-console-consumer.sh --topic web_log --from-beginning --bootstrap-server <YOUR_KAFKA_NAME>:9092`