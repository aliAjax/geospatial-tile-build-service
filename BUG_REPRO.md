# Bug Reproduction

使用零值依赖调用健康检查、限流和认证包装器时会直接 panic；默认配置校验也会失败。红灯验证真实报告了 `panic: nil health`、`panic: nil limits`、`panic: nil auth` 以及默认配置无效，堆栈落在 `internal/platform/http` 和 `internal/platform/config` 的处理函数。

