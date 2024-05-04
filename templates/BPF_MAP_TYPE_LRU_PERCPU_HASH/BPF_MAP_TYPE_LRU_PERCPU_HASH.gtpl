{{define "secMap"}}
{{range $key, $values := .Maps }}
{{if eq $values.Type "BPF_MAP_TYPE_LRU_PERCPU_HASH"}}
struct {
    __uint(type, {{.Type}});
    __type(key, {{.Key}});
    __type(value, {{.Value}});
    __uint(max_entries, {{.MaxEntries}});
} {{.Name}} SEC(".maps");
{{end}}
