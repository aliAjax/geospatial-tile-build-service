# Bug Reproduction

请求取消时瓦片编码返回文本化的模糊错误，调用方无法用 `errors.Is` 判断 `context.Canceled`。运行 `TestEncoderPreservesWriteErrorChain` 及其跨层错误链测试会失败，错误信息显示底层取消原因已丢失。

