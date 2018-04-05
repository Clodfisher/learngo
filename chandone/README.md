# done   

### 说明   

此函数主要用于测试，对于结构体中对于函数类型的声明。将系统的包对象由简单的传参方式，改为传对象函数类型。    

### 流程    

创建10个chan，分两次对这10个chan分进行读写操作，同时实现对于并发任务的等待操作。       

### 输出    

```
Worker 3 received d
Worker 4 received e
Worker 0 received a
Worker 0 received A
Worker 7 received h
Worker 9 received j
Worker 1 received b
Worker 8 received i
Worker 2 received c
Worker 5 received f
Worker 6 received g
Worker 9 received J
Worker 1 received B
Worker 2 received C
Worker 3 received D
Worker 4 received E
Worker 7 received H
Worker 8 received I
Worker 6 received G
Worker 5 received F
```   
