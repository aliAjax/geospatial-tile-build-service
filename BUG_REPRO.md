# Bug Reproduction

非法数据集 JSON 经过 adapter 和 application 层后，原始解析错误没有保留在错误链中。运行 `TestDecodeDatasetPreservesJSONError` 会失败，报告无法用 `errors.Is` 或 `errors.As` 识别底层 JSON 错误。

