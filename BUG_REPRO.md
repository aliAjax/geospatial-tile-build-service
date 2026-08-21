# Bug Reproduction

同一频道连续发布新版本后，接口显示成功但当前版本仍指向旧版本，历史列表与状态摘要也不一致。运行 `TestPublicationCurrentReleaseActive` 会失败，表现为最新成功发布没有成为当前版本。

