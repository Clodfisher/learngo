# selecttime      

### 说明   

此函数主要用于测试select，定时器的使用情况。    

### 知识    

此样例涉及到的知识点如下所示：    

* select + default 可以构成一个非阻塞式的管道。     
* select 中每次只能执行一个case，当有多个case可以执行时，随机一个进行执行。    
* 对于为nil的chan，在select中可以正常运行，但是无法select到，无法获取到数据。  
* 定时器的三种常见使用。      

### 流程    

创建两个chan在select中通过case进行数据的收，同时将收到的数据，通过另一个case进行发送。在整个过程中，要避免收的速度过快，到时发送不能及时，从而将就收到的数据添加到缓存中，进行发送，对于缓存的数据要有个定时器每秒钟打印下缓存的长度。这个流程实现后，再添加一个10秒钟的定时器，进行退出操作。同时增加个case，每次select花费800毫秒就打印个timeout，即800毫秒内没有收到数据的打timeout。      

### 输出    
```
queue len = 3
Worker received 0
queue len = 5
Worker received 0
queue len = 5
Worker received 1
queue len = 10
Worker received 1
queue len = 10
Worker received 2
queue len = 11
Worker received 2
queue len = 12
Worker received 3
queue len = 13
Worker received 3
queue len = 14
Worker received 4
bye...
```