{{define "secFunction"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeKprobeTime"}}
// 定义 BPF 程序的入口点
SEC("kprobe/{{$element.Name}}")
int kprobe_sys_read(struct pt_regs *ctx) {
    // 增加计数器
    bpf_map_increment_elem(&function_count_{{index}}, &zero);
    return 0;
}
    {{end}}
{{end}}
{{end}}