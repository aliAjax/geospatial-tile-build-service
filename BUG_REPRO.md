# Bug Reproduction

连续过滤两批 GeoJSON 时，第二批结果会带上第一批的属性或 geometry 变化。运行 `TestFilterDoesNotMutateOriginalFeatures` 会失败，表现为输入 feature、属性 map 或切片被原地修改。

