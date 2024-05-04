{{define "secMap"}}
{{range $key, $values := .Maps }}
struct {
    __uint(type, {{.Type}});
    __type(key, {{.Key}});
    __type(value, {{.Value}});
    __uint(max_entries, {{.MaxEntries}});
} {{.Name}} SEC(".maps");
{{end}}
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 1024);
} start_time_map SEC(".maps");

{{end}}
