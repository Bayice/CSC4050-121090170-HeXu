{{define "secEventStruct"}}
{{range $key, $values := .Structs }}
struct {{.Name}} {
    {{range $key, $value := .Members }}
    {{$value}} {{$key}};{{end}}
};
{{if .Members}}const struct {{.Name}} *unused_{{.Name}} __attribute__((unused));{{end}}{{end}}{{end}}


// 定义 BPF 映射，用于存储时间戳
BPF_MAP_DEF(timestamps) = {
    .map_type = BPF_MAP_TYPE_HASH,
    .key_size = sizeof(int),
    .value_size = sizeof(u64),
    .max_entries = 1,
};
// 将 BPF 映射导出到用户空间
BPF_MAP_ADD(timestamps);