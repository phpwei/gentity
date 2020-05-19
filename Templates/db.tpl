package models
{{$ModelName:=(printf "%s" .TableName) | RemoveTablePrefix | CamelCase}}
 type {{$ModelName}}  struct {
    {{range .Data}}
        {{index . "Field" | CamelCase}} {{index . "Type" | MapType}} `gorm:"column:{{index . "Field"}};{{index . "Extra"}}" json:"{{index . "Field"}}"`
	{{end}}
 }
 {{range .Data}}
  func(this *{{$ModelName}}) With{{index . "Field" | CamelCase}}({{index . "Field" | CamelCase}} {{index . "Type" | MapType}}) *{{$ModelName}}  {
      this.{{index . "Field" | CamelCase}}={{index . "Field" | CamelCase}}
      return this
  }
  {{end}}

  {{$ModelOptions:=(printf "%sOptions" $ModelName)}}
   type {{$ModelOptions}} struct {
       apply func(*{{$ModelName}})
   }
   func New{{$ModelName}}(opts...*{{$ModelName}}Options)  *{{$ModelName}} {
       m:= &{{$ModelName}}{}
       for _,opt:=range opts{
           opt.apply(m)
       }
       return m
   }
    {{range .Data}}
   func {{$ModelName}}{{index . "Field" | CamelCase}}({{index . "Field"}} {{index . "Type" | MapType}}) *{{$ModelOptions}} {
       return &{{$ModelOptions}}{func(model *{{$ModelName}}) {
           model.{{index . "Field" | CamelCase}}={{index . "Field"}}
       }}
   }
   {{end}}









