# Bug Reproduction

长批次文件存储处理会积累未释放的 reader/body，取消时又丢失最初错误。运行 `TestFileStoreCancellationReleasesHandle` 可复现句柄释放和错误链异常，取消返回值不再保留底层原因。

