# Bug Reproduction

用默认或 nil 缩放输入枚举瓦片时可能直接发生 nil pointer dereference，空路径和边界坐标还会产生空列表或不稳定结果。红灯输出真实报告 `runtime error: invalid memory address or nil pointer dereference`，堆栈指向 `internal/tiling/application/service.go` 的可选枚举路径。

