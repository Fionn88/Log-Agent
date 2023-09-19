# Log-Agent

## issue 

- 傳送 Struct 發生錯誤
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