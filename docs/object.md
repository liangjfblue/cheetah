model设计


调度器集群
- id
- ip                机器ip
- status            状态(1-leader 2-follower)


执行器组worker group
- id
- name              执行器组名
- workers           执行器列表, 包含id和name, 显示name, id用于查询详细信息. 支持单个执行器增加, 按执行器类型tag增加 
- status            状态, 正常-组中全部执行器心跳正常, 异常-组中有执行器下线或者异常. 定时更新

执行器worker
- id
- host              机器ip
- name              机器名
- srvName           服务名, 用于向注册中心注册
- tag               服务类,可选,用于调度时匹配特定机器)


任务job
- id
- jobId             全局唯一, uuid
- name              任务名字
- groupId           所属执行器组, 可选, 可自行指定执行器组
- tag               所属执行器类型, 用于调度器下发任务时匹配特定执行器
- cron              cron表达式
- type              任务类型, 包括命令行, shell脚本, code[golang,python], http
- target            目标任务, 命令行-命令, shell脚本-shell代码, code-源代码, http-http接口
- ip                任务执行所在执行器ip
- status            执行状态, 1-执行中 2-完成  3-未知  4-错误
- times             耗时
- remark            备注
- username          创建人userId,显示username
- createTime        创建时间Unix()
- startTime         开始时间Unix()
- endTime           结束时间Unix()
- result            返回结果, 1-ok 2-failed


执行日志记录
- id
- jobName           任务名
- workerGroupName   执行器组名
- workerName        执行器名
- workerIp          执行器ip
- cron              cron表达式
- type              任务类型, 包括命令行, shell脚本, code[golang,python], http
- status            执行状态, 1-执行中 2-完成  3-未知  4-错误
- result            返回结果, 1-ok 2-failed
- times             耗时
- username          创建人userId,显示username
- createTime        创建时间Unix()


## web api
### 执行器页面
- 新增    [POST]      /v1/workers
- 删除    [DELETE]    /v1/workers/:id
- 查找    [GET]       /v1/workers/:id
- 更新    [PUT]       /v1/workers/:id
- 列表    [GET]       /v1/workers

###执行器组页面
- 新增    [POST]      /v1/worker-groups
- 删除    [DELETE]    /v1/worker-groups/:id
- 查找    [GET]       /v1/worker-groups/:id
- 更新    [PUT]       /v1/worker-groups/:id
- 列表    [GET]       /v1/worker-groups

###任务页面
- 新增    [POST]      /v1/jobs
- 删除    [DELETE]    /v1/jobs/:id
- 查找    [GET]       /v1/jobs/:id
- 更新    [PUT]       /v1/jobs/:id
- 列表    [GET]       /v1/jobs

###调度器集群页面
- 列表    [GET]       /v1/masters

###调度日志页面
- 列表    [GET]       /v1/scheduler-logs

###调度
- 开始调度    [POST]      /v1/schedulers         支持批量,延迟,定时
- 停止调度    [DELETE]    /v1/schedulers         支持批量






