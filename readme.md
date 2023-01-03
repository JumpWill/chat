#

## airflow任务问题

### 每日数据未更新

每日需要更新的数据没有成功更新，任务执行失败,不是没有调度

### 执行失败总结

1. Deadlock found when trying to get lock; try restarting transaction
见alpha_daily  9f0d8fe4-f352-4cbe-bfce-f70bc8c714db at 2022-12-15, 13:00:00  与airflow连接数据库有关.可能是airflow有点慢，刷日志数据被某个事务保持    2.4.2版本修复,是否需要更新airflow, 待检验  

2. 反复在refreshing taskinstance process timeout  改excution_timeout  见strategy_daily d36   12-15     换airflow部署方式能解决，因为反复需要连接, 待检验  

3. airflow  State of this instance has been externally set to failed. Terminating instance.   将AIRFLOW__SCHEDULER__SCHEDULE_AFTER_TASK_EXECUTION为false，待检验  

4. 访问userservice服务的 http/read timeout
(userservice的服务只有两个worker,而它的探针接口较为频繁，而如果探针服务失败，会导致web服务有问题,以至于其他服务带调用出错)    userservice暂时稳定,但是timeout

5. factor_upload模块 (读取dataproxy，connection reset by peer。airflow完成新部署应该是能解决的)
ariflow的任务    12-15号基本所有任务    保证这个dataproxy能解决

6. st环境的数据有问题，用19-20年的数据进行.

7. userservice服务内存100%，且接口报错,?

            start    end

backtest
daikly

1. 工作台日志, 提交任务之前将任务清理           t0   done
2. ping mongo userservice内存泄漏相关        t2         todo  
3. prd的factor和stra t1                         done
4. 接口加上metic                            t0   done
5. 股票预览                      t0                    todo 读data_proxy慢  
6. 生产aiflow dags 同步到代码仓库中.          t2   done
7. 期货镜像里的文件 是有问题的@Xx              t0   done
8. 超时问题                                 t1         todo 截至2022年11月14日，已在starlette==0.21.0

9. 执行策略 end_date早于end_date
10. 各种任务date传值的问题(统一看下，涉及到任务里的配置)
11. env里面禁用缓存,airflow的缓存问题？       修改airflow的任务脚本
