# Bug Reproduction

清单请求第一次超时后，后续请求会继承已经结束的 deadline。运行 `TestManifestRequestContextDoesNotLeak` 会失败，表现为后续请求被错误取消；问题经过 HTTP、manifest 和压缩调用链传播。

