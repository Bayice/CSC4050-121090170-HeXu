
# Example Usage

```shell
go run cmd/kb/main.go generate --tmpl templates/kprobe.c.tmpl,templates/load.go.tmpl --conf configs/default.toml --header ./headers --output-dir ./output
```

作用：基于给定的这两个template，以及对应的配置文件，生成hook到配置文件里指定的内核函数，并统计他们的执行时间

如何运行生成的ebpf代码：

```
cd output
sudo /usr/local/go/bin/go run *.go 
```

此时执行 `sudo cat /sys/kernel/debug/tracing/trace_pipe | grep "executed in"`，可以观察到打印出来的日志



go run cmd/kb/main.go generate --tmpl templates/kprobe.c.gtpl,templates/load.go.gtpl --conf configs/default.toml --header ./headers --output-dir ./output --execute load.go.gtpl

go run cmd/kb/main.go generate --tmpl templates/kprobe.c.gtpl,templates/load.go.gtpl --conf configs/default.toml --header ./headers --output-dir ./output --execute kprobe.c.gtpl,load.go.gtpl

go run cmd/kb/main.go generate --tmpl examples/dynamic_uprobe/template/uprobe.c.gtpl --conf configs/default.toml --header ./headers --output-dir ./output --execute uprobe.c.gtpl



go run cmd/kb/main.go generate --tmpl templates/kprobe.c.gtpl,templates/load.go.gtpl --conf configs/default.toml --header ./headers --output-dir ./output --execute kprobe.c.gtpl

go run cmd/kb/main.go generate --tmpl templates/uprobe.c.gtpl --conf configs/default.toml --header ./headers --output-dir ./output


go run cmd/kb/main.go generate --tmpl examples/dynamic_uprobe/template/uprobe.c.gtpl --conf examples/dynamic_uprobe/conf/default.toml --headr ./headers --output-dir examples/dynamic_uprobe/output/


go run cmd/kb/main.go generate --tmpl examples/dynamic_uprobe/template/uprobe.c.gtpl --conf examples/dynamic_uprobe/conf/default.toml --header ./headers --output-dir examples/dynamic_uprobe/output/

 go run cmd/kb/main.go --help

 go run cmd/kb/main.go generate --help


go run cmd/kb/main.go generate --tmpl templates/kernel_c/base.c.gtpl,templates/kernel_c/secEventStruct.c.gtpl,templates/kernel_c/secFunction.c.gtpl,templates/kernel_c/secInclude.c.gtpl,templates/kernel_c/secLicense.c.gtpl,templates/kernel_c/secMap.c.gtpl --conf configs/default.toml --header ./headers --output-dir ./output --execute base.c.gtpl

go run cmd/kb/main.go generate --tmpl templates/kernel_c/base.c.gtpl,templates/kernel_c/secEventStruct.c.gtpl,templates/kernel_c/secFunction.c.gtpl,templates/kernel_c/secInclude.c.gtpl,templates/kernel_c/secLicense.c.gtpl,templates/kernel_c/secMap.c.gtpl --conf configs/default.toml --header ./headers --output-dir ./output --execute base.c.gtpl