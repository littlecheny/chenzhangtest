#webservice

## 接口
### 提交任务
http://localhost:8080/submitasks
提交格式
{
    "Index": 0,
    "BurstTime": 10
}

### 选择调度算法
http://localhost:8080/strategy
提交格式
{
    "Algo": "FIFO"
}
默认运行FIFO调度算法

##设计模式
- 单例模式：TaskManager 作为全局唯一的任务管理器，负责维护所有用户的任务队列。通过读写锁保证并发安全;
Add方法：AddTasks(userID string, tasks []domain.Task) error 使用Lock保护任务队列的并发写入，避免竞态条件。
Snapshot方法：Snapshot(userID string) ([]domain.Task, error) 使用RLock锁保护任务队列的并发读取，返回用户当前任务的深拷贝快照，避免并发读写冲突。
ScheduleNow方法：ScheduleNow(userID string, algo domain.Schedule) (domain.SchedulerState, error) 调用指定的调度算法对用户任务进行调度，返回调度结果，更新任务状态时使用Lock锁保护，避免并发修改导致的不一致状态。

- 工厂模式：在domain层定义Schedule接口，不同的调度算法实现该接口，如FIFO、SJF等。在services层根据用户提交的调度算法，动态创建不同的调度服务。

- 依赖注入：在路由层通过构造函数注入 TaskManager，而不是直接在路由处理函数中创建实例。


保证高并发的方案B：通道 + 调度后台协程
- 核心思想：用一个或多个 goroutine 作为调度 worker，路由把任务通过 channel 投递到队列；worker 消费队列、运行算法、合并状态。
- 优点：天然异步、易做限流和队列度量；调度周期化更容易
- 缺点：对你当前的简洁结构来说复杂度增加；跨用户隔离要设计多个队列或带 userID 的消息，并增加管理组件