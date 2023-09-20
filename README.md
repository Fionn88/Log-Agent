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