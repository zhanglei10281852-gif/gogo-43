# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

路由任务列表中任意位置出现 nil job 时，调用会在完成输入拒绝前发生空指针 panic。请修复为稳定返回明确的 nil-job 输入错误，确保检查发生在排序等解引用之前，同时保持正常任务路由和拒绝统计行为不变，并保证全量测试通过。

## 含 Bug 版本

- 仓库：zhanglei10281852-gif/gogo-43
- 仓库地址：https://github.com/zhanglei10281852-gif/gogo-43.git
- parent SHA：9be4dbb196c7600fa08dd24f4efa8d6cb5a9358a

## 复现步骤

```bash
git clone -- https://github.com/zhanglei10281852-gif/gogo-43.git bug-repro
cd bug-repro
git checkout --detach 9be4dbb196c7600fa08dd24f4efa8d6cb5a9358a
go test ./internal/query -run "^TestRouteReturnsErrorForNilJob$" -count=1 -v
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/query -run "^TestRouteReturnsErrorForNilJob$" -count=1 -v
=== RUN   TestRouteReturnsErrorForNilJob
--- FAIL: TestRouteReturnsErrorForNilJob (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x50 pc=0x506b29]

goroutine 18 [running]:
testing.tRunner.func1.2({0x51dbe0, 0x639720})
	/usr/local/go/src/testing/testing.go:1631 +0x24a
testing.tRunner.func1()
	/usr/local/go/src/testing/testing.go:1634 +0x377
panic({0x51dbe0?, 0x639720?})
	/usr/local/go/src/runtime/panic.go:770 +0x132
QueueForge/internal/query.Route.sortJobs.func1(0xc000126090?, 0x53f70d?)
	/app/internal/query/query.go:222 +0x69
sort.insertionSort_func({0xc00010ede0?, 0xc00014c040?}, 0x0, 0x2)
	/usr/local/go/src/sort/zsortfunc.go:12 +0xa7
sort.stable_func({0xc00010ede0?, 0xc00014c040?}, 0x2)
	/usr/local/go/src/sort/zsortfunc.go:343 +0x75
sort.SliceStable({0x514800?, 0xc000126090?}, 0xc00010ede0)
	/usr/local/go/src/sort/slice.go:44 +0xb0
QueueForge/internal/query.sortJobs(...)
	/app/internal/query/query.go:213
QueueForge/internal/query.Route({0xc00010ef50, 0x2, 0x61d208?}, {{0x53f4c0, 0x6}, {0xc000155f30, 0x1, 0x1}, 0x0, {0x0, ...}}, ...)
	/app/internal/query/query.go:327 +0x277
QueueForge/internal/query.TestRouteReturnsErrorForNilJob(0xc00013e4e0)
	/app/internal/query/query_test.go:14 +0x245
testing.tRunner(0xc00013e4e0, 0x54b668)
	/usr/local/go/src/testing/testing.go:1689 +0xfb
created by testing.(*T).Run in goroutine 1
	/usr/local/go/src/testing/testing.go:1742 +0x390
FAIL	QueueForge/internal/query	0.005s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/query -run "^TestRouteReturnsErrorForNilJob$" -count=1 -v
=== RUN   TestRouteReturnsErrorForNilJob
--- FAIL: TestRouteReturnsErrorForNilJob (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x50 pc=0x10e68c]

goroutine 33 [running]:
testing.tRunner.func1.2({0x125b00, 0x243800})
	/usr/local/go/src/testing/testing.go:1631 +0x1c4
testing.tRunner.func1()
	/usr/local/go/src/testing/testing.go:1634 +0x33c
panic({0x125b00?, 0x243800?})
	/usr/local/go/src/runtime/panic.go:770 +0x124
QueueForge/internal/query.Route.sortJobs.func1(0x0?, 0x1?)
	/app/internal/query/query.go:222 +0x6c
sort.insertionSort_func({0x4000074dd8?, 0x40001aa000?}, 0x0, 0x2)
	/usr/local/go/src/sort/zsortfunc.go:12 +0xc0
sort.stable_func({0x4000074dd8?, 0x40001aa000?}, 0x2)
	/usr/local/go/src/sort/zsortfunc.go:343 +0x6c
sort.SliceStable({0x11c7e0?, 0x400019a030?}, 0x4000074dd8)
	/usr/local/go/src/sort/slice.go:44 +0xac
QueueForge/internal/query.sortJobs(...)
	/app/internal/query/query.go:213
QueueForge/internal/query.Route({0x4000074f48, 0x2, 0x40001a4638?}, {{0x1473c8, 0x6}, {0x4000133f28, 0x1, 0x1}, 0x0, {0x0, ...}}, ...)
	/app/internal/query/query.go:327 +0x1ac
QueueForge/internal/query.TestRouteReturnsErrorForNilJob(0x4000194000)
	/app/internal/query/query_test.go:14 +0x190
testing.tRunner(0x4000194000, 0x153590)
	/usr/local/go/src/testing/testing.go:1689 +0xec
created by testing.(*T).Run in goroutine 1
	/usr/local/go/src/testing/testing.go:1742 +0x318
FAIL	QueueForge/internal/query	0.151s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

nil 位于任务列表任意位置时均先返回错误而不 panic；不会在错误返回前排序或执行拒绝流程；正常路由行为不回归；双架构定向、全量、build/vet 通过。
