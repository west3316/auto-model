package {{ .PackageName }}

{{ if .Enums }}
const({{ range $item := .Enums }}
    // {{ $item.Comment }}{{ range $val := $item.Values }}
    {{ $item.Prefix }}{{ $val }} = "{{ $val }}"{{ end }}
    
    {{ end }}
)
{{ range $item := .Enums }}
// IsValid{{$item.Prefix}} 检测枚举，{{ $item.Comment }}
func IsValid{{$item.Prefix}}(v string) bool {
    return map[string]bool{ {{ range $val := $item.Values }}
        {{ $item.Prefix }}{{ $val }}: true,{{ end }}
    }[v]
}
{{ end }}
{{ end }}

// {{ .StructName }} 表名: {{ .TableName }}
type {{ .StructName }} struct {
    {{ range .Fields }}
    // {{ .Comment }}
    {{ .Name }} {{ .Type }} {{ .Tags }}
    {{ end }}
}

// TableName 表名
func (*{{ .StructName }}) TableName() string {
    return "{{ .TableName }}"
}
