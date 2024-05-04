{{define "secMap"}}
{{range $key, $values := .Maps }}
{{if eq $values.Type "BPF_MAP_TYPE_HASH_OF_MAPS"}}
struct {
    __uint(type, {{.Type}});
    __type(key, {{.Key}});
    __type(value, {{.Value}});
    __uint(max_entries, {{.MaxEntries}});
} inner_map2 SEC(".maps");

struct {
    __uint(type, {{.Type}});
    __type(key, {{.Key}});
    __type(value, {{.Value}});
    __uint(max_entries, {{.MaxEntries}});
} inner_map2 SEC(".maps");

struct {
    __uint(type, {{.Type}});
    __type(key, {{.Key}});
    __type(value, {{.Value}});
    __uint(max_entries, {{.MaxEntries}});
} {{.Name}} SEC(".maps");
{{end}}

