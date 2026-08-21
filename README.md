# Geospatial Tile Build Service

纯Go地理空间矢量数据切片编译与版本分发服务，编号14。项目提供数据集与版本管理、GeoJSON校验、坐标转换、XYZ瓦片坐标、确定性构建摘要、清单接口和MVT风格二进制切片分发。默认使用线程安全内存适配器，PostgreSQL和对象存储通过接口替换。

## 启动

```bash
go build ./...
go vet ./...
go run ./cmd/server
```

默认地址为`http://localhost:18114`，支持`TILE_HTTP_ADDR`、`TILE_MAX_BODY`和`TILE_BUILD_WORKERS`环境变量。收到SIGINT/SIGTERM后5秒内优雅停机。

## 验证主流程

```bash
curl -fsS localhost:18114/healthz
curl -fsS localhost:18114/readyz
curl -X POST localhost:18114/v1/datasets -H 'content-type: application/json' -d '{"id":"city","name":"City Roads","crs":"EPSG:4326","layers":["roads"]}'
curl -X POST localhost:18114/v1/geojson/validate -H 'content-type: application/json' --data-binary @examples/sample.geojson
curl -X POST localhost:18114/v1/builds -H 'content-type: application/json' -d '{"id":"build-1","version_id":"v1","input":{"source":"sample.geojson","zoom":[0,14]}}'
curl -fsS localhost:18114/v1/builds?id=build-1
curl -fsS localhost:18114/v1/manifests/city/v1
curl -fsS localhost:18114/tiles/city/v1/0/0/0.pbf | od -An -tc
curl -fsS localhost:18114/metrics
```

## 结构与规模

顶层包含`cmd`、`internal`、`api`、`configs`、`migrations`、`deploy`和`scripts`。领域按`domain`、`application`、`adapter`、`infrastructure`分层，依赖采用接口和构造函数注入，跨模块传递`context.Context`，错误使用包装错误，日志使用`slog`。非测试Go源码目标不少于2400行；排除测试、生成代码、依赖、锁文件、迁移SQL、配置样例和构建产物，禁止重复代码凑行数。
