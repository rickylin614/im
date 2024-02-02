
# Prometheus

## 自定義監控

| Metric Name     | Description                                            |
|-----------------|--------------------------------------------------------|
| online_user_num | Custom metric representing the number of online users. |

## gorm

| Metric Name               | Description                                                        |
|---------------------------|--------------------------------------------------------------------|
| gorm_dbstats_max_open_connections   | Maximum number of open connections to the database.                |
| gorm_dbstats_open_connections       | The number of established connections both in use and idle.        |
| gorm_dbstats_in_use                | The number of connections currently in use.                        |
| gorm_dbstats_idle                  | The number of idle connections.                                    |
| gorm_dbstats_wait_count            | The total number of connections waited for.                        |
| gorm_dbstats_wait_duration         | The total time blocked waiting for a new connection.               |
| gorm_dbstats_max_idle_closed       | The total number of connections closed due to SetMaxIdleConns.     |
| gorm_dbstats_max_lifetime_closed   | The total number of connections closed due to SetConnMaxLifetime.  |
| gorm_dbstats_max_idletime_closed   | The total number of connections closed due to SetConnMaxIdleTime.  |


## redis

| Name                      | Type           | Description                                                                 |
|---------------------------|----------------|-----------------------------------------------------------------------------|
| `pool_hit_total`          | Counter metric | number of times a connection was found in the pool                          |
| `pool_miss_total`         | Counter metric | number of times a connection was not found in the pool                      |
| `pool_timeout_total`      | Counter metric | number of times a timeout occurred when getting a connection from the pool  |
| `pool_conn_total_current` | Gauge metric   | current number of connections in the pool                                   |
| `pool_conn_idle_current`  | Gauge metric   | current number of idle connections in the pool                              |
| `pool_conn_stale_total`   | Counter metric | number of times a connection was removed from the pool because it was stale |

## 預設

### CPU Metrics:

| Metric Name               | Description                                                  | Type    |
| ------------------------- | ------------------------------------------------------------ | ------- |
| go_gc_duration_seconds    | A summary of the pause duration of garbage collection cycles. | summary |
| go_goroutines             | Number of goroutines that currently exist.                   | gauge   |
| process_cpu_seconds_total | Total user and system CPU time spent in seconds.             | counter |

### Memory Metrics:

| Metric Name                      | Description                                                  | Type    |
| -------------------------------- | ------------------------------------------------------------ | ------- |
| go_memstats_alloc_bytes          | Number of bytes allocated and still in use.                  | gauge   |
| go_memstats_alloc_bytes_total    | Total number of bytes allocated, even if freed.              | counter |
| go_memstats_heap_alloc_bytes     | Number of heap bytes allocated and still in use.             | gauge   |
| go_memstats_heap_idle_bytes      | Number of heap bytes waiting to be used.                     | gauge   |
| go_memstats_heap_inuse_bytes     | Number of heap bytes that are in use.                        | gauge   |
| go_memstats_heap_objects         | Number of allocated objects.                                 | gauge   |
| go_memstats_heap_released_bytes  | Number of heap bytes released to the OS.                     | gauge   |
| go_memstats_heap_sys_bytes       | Number of heap bytes obtained from the system.               | gauge   |
| go_memstats_last_gc_time_seconds | Number of seconds since 1970 of the last garbage collection. | gauge   |
| go_memstats_mallocs_total        | Total number of mallocs.                                     | counter |
| go_memstats_mcache_inuse_bytes   | Number of bytes in use by mcache structures.                 | gauge   |
| go_memstats_mcache_sys_bytes     | Number of bytes used for mcache structures obtained from the system. | gauge   |
| go_memstats_mspan_inuse_bytes    | Number of bytes in use by mspan structures.                  | gauge   |
| go_memstats_mspan_sys_bytes      | Number of bytes used for mspan structures obtained from the system. | gauge   |
| go_memstats_next_gc_bytes        | Number of heap bytes when next garbage collection will take place. | gauge   |
| go_memstats_other_sys_bytes      | Number of bytes used for other system allocations.           | gauge   |
| go_memstats_stack_inuse_bytes    | Number of bytes in use by the stack allocator.               | gauge   |
| go_memstats_stack_sys_bytes      | Number of bytes obtained from the system for stack allocator. | gauge   |
| go_memstats_sys_bytes            | Number of bytes obtained from the system.                    | gauge   |
| process_max_fds                  | Maximum number of open file descriptors.                     | gauge   |
| process_open_fds                 | Number of open file descriptors.                             | gauge   |
| process_resident_memory_bytes    | Resident memory size in bytes.                               | gauge   |
| process_virtual_memory_bytes     | Virtual memory size in bytes.                                | gauge   |