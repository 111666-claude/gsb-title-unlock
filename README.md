# title-unlock

称号解锁判定。判定逻辑在库 module `example.com/titlekit`，命令行在 `example.com/titlecli`，用 `go.work` 串起来。

```
go test example.com/titlekit/... example.com/titlecli/...   # 或在 lib/ 与 app/ 里各跑 go test ./...
go vet example.com/titlekit/... example.com/titlecli/...
go run ./app --sample=achv
go run ./app --sample=prereq
go run ./app --sample=event
go run ./app --sample=replay
go run ./app --sample=work
```

## 口径

- **同时满足**：等级 ≥ `need_level`、成就数 ≥ `need_achv`、在线小时 ≥ `need_hours`，三条都要满足。
- **前置**：`needs` 里的称号必须都已经解锁。
- **覆盖层**：`event` 为真的活动称号不受等级门槛限制，其余条件照旧。
- **幂等**：同一个 `req_id` 只结算一次，重复请求返回首次结果且不再登记解锁。

## 不变量

- 解锁成功的判定必须三条门槛与前置全部满足。
- 已解锁列表里不出现重复称号。
- `scanned` 不随判定次数放大：四千次判定的 `scanned` 不超过 9000。

## 输出契约

每次判定打印一行 JSON：`title`、`status`（`unlocked` / `locked`）、`unlocked`、`scanned`。
`--sample=work` 只打印 `{"judgements": N, "scanned": N}`。
