# Bug Reproduction

并发构建时，存储层复用可变输入缓冲区，读取与写入同时发生会触发 race，并可能读到半截数据。运行 `TestBuildInputSnapshotUnderConcurrentAccess` 的 race 检查可稳定看到 `WARNING: DATA RACE`，堆栈指向 `internal/storage/infrastructure/memory.go` 的存储读写路径。

