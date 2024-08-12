package genstruct

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/cnf"
	"github.com/xxjwxc/gormt/data/view/generate"
	"github.com/xxjwxc/gormt/data/view/genfunc"
)

// SetName Setting element name.设置元素名字
func (e *GenElement) SetName(name string) {
	e.Name = name
}

// SetType Setting element type.设置元素类型
func (e *GenElement) SetType(tp string) {
	e.Type = tp
}

// SetNotes Setting element notes.设置注释
func (e *GenElement) SetNotes(notes string) {
	e.Notes = strings.Replace(notes, "\n", ",", -1)
}

// AddTag Add a tag .添加一个tag标记
func (e *GenElement) AddTag(k string, v string) {
	if e.Tags == nil {
		e.Tags = make(map[string][]string)
	}
	e.Tags[k] = append(e.Tags[k], v)
}

// Generate Get the result data.获取结果数据
func (e *GenElement) Generate() string {
	tag := ""
	if e.Tags != nil {
		var ks []string
		for k := range e.Tags {
			ks = append(ks, k)
		}
		sort.Strings(ks)

		var tags []string
		for _, v := range ks {
			tags = append(tags, fmt.Sprintf(`%v:"%v"`, v, strings.Join(e.Tags[v], ";")))
		}
		tag = fmt.Sprintf("`%v`", strings.Join(tags, " "))
	}

	var p generate.PrintAtom
	if len(e.Notes) > 0 {
		p.Add(e.Name, e.Type, tag, "// "+e.Notes)
	} else {
		p.Add(e.Name, e.Type, tag)
	}

	return p.Generates()[0]
}

func (e *GenElement) GenerateCreate() string {
	tag := ""
	if e.Tags != nil {
		var ks []string
		for k := range e.Tags {
			ks = append(ks, k)
		}
		sort.Strings(ks)

		var tags []string
		for _, v := range ks {
			if strings.ToLower(v) != "json" {
				continue
			}
			if e.Tags[v][0] == "-" {
				e.Tags[v][0] = strings.ToLower(e.Name)
			}
			tags = append(tags, fmt.Sprintf(`%v:"%v"`, v, strings.Join(e.Tags[v], ";")))
		}
		tag = fmt.Sprintf("`%v`", strings.Join(tags, " "))
	}

	var p generate.PrintAtom
	if len(e.Notes) > 0 {
		p.Add(e.Name, e.Type, tag, "// "+e.Notes)
	} else {
		p.Add(e.Name, e.Type, tag)
	}

	return p.Generates()[0]
}

func (e *GenElement) GenerateDelete() string {
	tag := ""
	if e.Tags != nil {
		var ks []string
		for k := range e.Tags {
			ks = append(ks, k)
		}
		sort.Strings(ks)

		var tags []string
		for _, v := range ks {
			if strings.ToLower(v) != "json" {
				continue
			}
			if e.Tags[v][0] == "-" {
				e.Tags[v][0] = strings.ToLower(e.Name)
			}
			tags = append(tags, fmt.Sprintf(`%v:"%v"`, v, strings.Join(e.Tags[v], ";")))
		}
		tag = fmt.Sprintf("`%v`", strings.Join(tags, " "))
	}

	var p generate.PrintAtom
	if len(e.Notes) > 0 {
		p.Add(e.Name, e.Type, tag, "// "+e.Notes)
	} else {
		p.Add(e.Name, e.Type, tag)
	}

	return p.Generates()[0]
}

func (e *GenElement) GenerateUpdate() string {
	tag := ""
	if e.Tags != nil {
		var ks []string
		for k := range e.Tags {
			ks = append(ks, k)
		}
		sort.Strings(ks)

		var tags []string
		for _, v := range ks {
			if strings.ToLower(v) != "json" {
				continue
			}
			if e.Tags[v][0] == "-" {
				e.Tags[v][0] = strings.ToLower(e.Name)
			}
			tags = append(tags, fmt.Sprintf(`%v:"%v"`, v, strings.Join(e.Tags[v], ";")))
		}
		tag = fmt.Sprintf("`%v`", strings.Join(tags, " "))
	}

	var p generate.PrintAtom
	if len(e.Notes) > 0 {
		p.Add(e.Name, e.Type, tag, "// "+e.Notes)
	} else {
		p.Add(e.Name, e.Type, tag)
	}

	return p.Generates()[0]
}

func (e *GenElement) GenerateList() string {
	tag := ""
	if e.Tags != nil {
		var ks []string
		for k := range e.Tags {
			ks = append(ks, k)
		}
		sort.Strings(ks)

		var tags []string
		for _, v := range ks {
			if strings.ToLower(v) != "json" {
				continue
			}
			if e.Tags[v][0] == "-" {
				e.Tags[v][0] = strings.ToLower(e.Name)
			}
			tags = append(tags, fmt.Sprintf(`%v:"%v"`, v, strings.Join(e.Tags[v], ";")))
		}
		tag = fmt.Sprintf("`%v`", strings.Join(tags, " "))
	}

	var p generate.PrintAtom
	if len(e.Notes) > 0 {
		p.Add(e.Name, e.Type, tag, "// "+e.Notes)
	} else {
		p.Add(e.Name, e.Type, tag)
	}

	return p.Generates()[0]
}

// GenerateColor Get the result data.获取结果数据
func (e *GenElement) GenerateColor() string {
	tag := ""
	if e.Tags != nil {
		var ks []string
		for k := range e.Tags {
			ks = append(ks, k)
		}
		sort.Strings(ks)

		var tags []string
		for _, v := range ks {
			tags = append(tags, fmt.Sprintf(`%v:"%v"`, v, strings.Join(e.Tags[v], ";")))
		}
		tag = fmt.Sprintf("`%v`", strings.Join(tags, " "))
	}

	var p generate.PrintAtom
	if len(e.Notes) > 0 {
		p.Add(e.Name, "\033[32;1m "+e.Type+" \033[0m", "\033[31;1m "+tag+" \033[0m", "\033[32;1m // "+e.Notes+" \033[0m")
	} else {
		p.Add(e.Name, "\033[32;1m "+e.Type+" \033[0m", "\033[31;1m "+tag+" \033[0m")
	}

	return p.Generates()[0]
}

//////////////////////////////////////////////////////////////////////////////
// struct
//////////////////////////////////////////////////////////////////////////////

// SetCreatTableStr Set up SQL create statement, backup use setup create statement, backup use.设置创建语句，备份使用
func (s *GenStruct) SetCreatTableStr(sql string) {
	s.SQLBuildStr = sql
}

// SetTableName Setting the name of struct.设置struct名字
func (s *GenStruct) SetTableName(name string) {
	s.TableName = name
}

// SetStructName Setting the name of struct.设置struct名字
func (s *GenStruct) SetStructName(name string) {
	s.Name = name
}

// SetNotes set the notes.设置注释
func (s *GenStruct) SetNotes(notes string) {
	if len(notes) == 0 {
		notes = "[...]" // default of struct notes(for export ).struct 默认注释(为了导出注释)
	}

	notes = s.Name + " " + notes

	a := strings.Split(notes, "\n")
	var text []string

	for _, v := range a {
		// if len(v) > 0 {
		text = append(text, "// "+v)
		// }
	}
	s.Notes = strings.Join(text, ";")
}

// AddElement Add one or more elements.添加一个/或多个元素
func (s *GenStruct) AddElement(e ...GenElement) {
	s.Em = append(s.Em, e...)
}

// GenerateTableName generate table name .生成表名
func (s *GenStruct) GenerateTableName() []string {
	tmpl, err := template.New("gen_tnf").Parse(genfunc.GetGenTableNameTemp())
	if err != nil {
		panic(err)
	}
	var data struct {
		TableName  string
		StructName string
	}
	data.TableName, data.StructName = s.TableName, s.Name
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	return []string{buf.String()}
}

// GenerateColumnName generate column name . 生成列名
func (s *GenStruct) GenerateColumnName() []string {
	tmpl, err := template.New("gen_tnc").Parse(genfunc.GetGenColumnNameTemp())
	if err != nil {
		panic(err)
	}
	var data struct {
		StructName string
		Em         []struct {
			ColumnName string
			StructName string
		}
	}
	data.StructName = s.Name
	for _, v := range s.Em {
		if strings.EqualFold(v.Type, "gorm.Model") { // gorm model
			data.Em = append(data.Em, []struct {
				ColumnName string
				StructName string
			}{
				{ColumnName: "id", StructName: "ID"},
				{ColumnName: "created_at", StructName: "CreatedAt"},
				{ColumnName: "updated_at", StructName: "UpdatedAt"},
				{ColumnName: "deleted_at", StructName: "DeletedAt"},
			}...)
		} else if len(v.ColumnName) > 0 {
			data.Em = append(data.Em, struct {
				ColumnName string
				StructName string
			}{ColumnName: v.ColumnName,
				StructName: v.Name,
			})
		}

	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	return []string{buf.String()}
}

// Generates Get the result data.获取结果数据
func (s *GenStruct) Generates() []string {
	var p generate.PrintAtom
	if config.GetIsOutSQL() {
		p.Add("/******sql******")
		p.Add(s.SQLBuildStr)
		p.Add("******sql******/")
	}
	p.Add(s.Notes)
	p.Add("type", s.Name, "struct {")
	mp := make(map[string]bool, len(s.Em))
	for _, v := range s.Em {
		if !mp[v.Name] {
			mp[v.Name] = true
			p.Add(v.Generate())
		}
	}
	p.Add("}")

	return p.Generates()
}

func (s *GenStruct) GeneratesReqCreate() []string {
	var p generate.PrintAtom
	//增加表和表之间的间隔
	p.Add("// " + centerString("", "=", 80))
	p.Add("// " + centerString(strings.ToLower(s.TableName)+"表", "-", 80))
	p.Add("// " + centerString("", "=", 80))
	p.Add("\n")
	p.Add("// " + s.Name + "Create " + s.Name + "表创建请求参数")
	p.Add("type", s.Name+"Create", "struct {")
	//加上Base类型
	p.Add("BaseRequest      `json:\"-\"`")
	p.Add("BaseTokenRequest      `json:\"-\"`")
	mp := make(map[string]bool, len(s.Em))
	for _, v := range s.Em {

		//创建无需id，uuid,created_time,updated_time,is_deleted字段
		if v.Name == "ID" || v.Name == "UUID" || v.Name == "CreatedTime" || v.Name == "UpdatedTime" || v.Name == "IsDeleted" {
			continue
		}

		if !mp[v.Name] {
			mp[v.Name] = true
			p.Add(v.GenerateCreate())
		}
	}
	p.Add("}")

	return p.Generates()
}
func (s *GenStruct) GeneratesRespCreate() []string {
	var p generate.PrintAtom
	p.Add("// " + centerString("", "=", 80))
	p.Add("// " + centerString(strings.ToLower(s.TableName)+"表", "-", 80))
	p.Add("// " + centerString("", "=", 80))
	p.Add("\n")
	p.Add("// " + s.Name + "Create " + s.Name + "表创建返回参数")
	p.Add("type", s.Name+"Create", "struct {}")
	return p.Generates()
}

func (s *GenStruct) GeneratesReqDelete() []string {
	var p generate.PrintAtom
	p.Add("// " + s.Name + "Delete " + s.Name + "表删除请求参数")
	p.Add("type", s.Name+"Delete", "struct {")
	//加上Base类型
	p.Add("BaseRequest      `json:\"-\"`")
	p.Add("BaseTokenRequest      `json:\"-\"`")
	mp := make(map[string]bool, len(s.Em))
	for _, v := range s.Em {

		//删除只需要uuid字段
		if v.Name == "UUID" {
		} else {
			continue
		}

		if !mp[v.Name] {
			mp[v.Name] = true
			p.Add(v.GenerateDelete())
		}
	}
	p.Add("}")

	return p.Generates()
}
func (s *GenStruct) GeneratesRespDelete() []string {
	var p generate.PrintAtom
	p.Add("// " + s.Name + "Delete " + s.Name + "表删除返回参数")
	p.Add("type", s.Name+"Delete", "struct {")
	p.Add("}")
	return p.Generates()
}
func (s *GenStruct) GeneratesReqUpdate() []string {
	var p generate.PrintAtom
	p.Add("// " + s.Name + "Update " + s.Name + "表更新请求参数")
	p.Add("type", s.Name+"Update", "struct {")
	//加上Base类型
	p.Add("BaseRequest      `json:\"-\"`")
	p.Add("BaseTokenRequest      `json:\"-\"`")
	mp := make(map[string]bool, len(s.Em))
	for _, v := range s.Em {

		//更新不需要id,is_deleted,created_time,updated_time字段,同时其他类型改为指针类型
		if v.Name == "ID" || v.Name == "IsDeleted" || v.Name == "CreatedTime" || v.Name == "UpdatedTime" {
			continue
		} else {
			v.Type = "*" + v.Type
		}

		if !mp[v.Name] {
			mp[v.Name] = true
			p.Add(v.GenerateUpdate())
		}
	}
	p.Add("}")

	return p.Generates()
}
func (s *GenStruct) GeneratesRespUpdate() []string {
	var p generate.PrintAtom
	p.Add("// " + s.Name + "Update " + s.Name + "表更新返回参数")
	p.Add("type", s.Name+"Update", "struct {")
	p.Add("}")
	return p.Generates()
}

func (s *GenStruct) GeneratesReqList() []string {
	var p generate.PrintAtom
	p.Add("// " + s.Name + "List " + s.Name + "表列表请求参数")
	p.Add("type", s.Name+"List", "struct {")
	//加上Base类型
	p.Add("BaseRequest      `json:\"-\"`")
	p.Add("BasePagination      `json:\"-\"`")
	p.Add("BaseTokenRequest      `json:\"-\"`")
	p.Add("}")

	return p.Generates()
}
func (s *GenStruct) GeneratesRespList() []string {
	var p generate.PrintAtom
	p.Add("// " + s.Name + "List " + s.Name + "表列表返回参数")
	p.Add("type", s.Name+"List", "struct {")
	p.Add("BasePagination      `json:\"-\"`")
	mp := make(map[string]bool, len(s.Em))
	for _, v := range s.Em {

		//列表不需要is_deleted,updated_time字段
		if v.Name == "IsDeleted" || v.Name == "UpdatedTime" {
			continue
		}

		if !mp[v.Name] {
			mp[v.Name] = true
			p.Add(v.GenerateList())
		}
	}
	p.Add("}")

	return p.Generates()
}

func (s *GenStruct) GeneratesService() []string {
	var p generate.PrintAtom
	p.Add(s.Name + "Create" + "(info request." + s.Name + "Create)(out response." + s.Name + "Create, code commons.ResponseCode, err error)")
	p.Add(s.Name + "Delete" + "(info request." + s.Name + "Delete)(out response." + s.Name + "Delete, code commons.ResponseCode, err error)")
	p.Add(s.Name + "Update" + "(info request." + s.Name + "Update)(out response." + s.Name + "Update, code commons.ResponseCode, err error)")
	p.Add(s.Name + "List" + "(info request." + s.Name + "List)(out response." + s.Name + "List, code commons.ResponseCode, err error)")

	return p.Generates()
}
func (s *GenStruct) GeneratesImpl() []string {
	var p generate.PrintAtom
	p.Add("// " + centerString("", "=", 80))
	p.Add("// " + centerString(s.Name+" service layer implementation", "-", 80))
	p.Add("// " + centerString("", "=", 80))
	p.Add("\n")
	p.Add("func (g genServiceImp) " + s.Name + "Create" + "(info request." + s.Name + "Create)(out response." + s.Name + "Create, code commons.ResponseCode, err error) {")
	p.Add("\t//todo")
	p.Add("\treturn")
	p.Add("}")

	p.Add("func (g genServiceImp) " + s.Name + "Delete" + "(info request." + s.Name + "Delete)(out response." + s.Name + "Delete, code commons.ResponseCode, err error) {")
	p.Add("\t//todo")
	p.Add("\treturn")
	p.Add("}")

	p.Add("func (g genServiceImp) " + s.Name + "Update" + "(info request." + s.Name + "Update)(out response." + s.Name + "Update, code commons.ResponseCode, err error) {")
	p.Add("\t//todo")
	p.Add("\treturn")
	p.Add("}")

	p.Add("func (g genServiceImp) " + s.Name + "List" + "(info request." + s.Name + "List)(out response." + s.Name + "List, code commons.ResponseCode, err error) {")
	p.Add("\t//todo")
	p.Add("\treturn")
	p.Add("}")

	return p.Generates()
}

// \033[3%d;%dm -%d;%d-colors!\033[0m\n
// GeneratesColor Get the result data on color.获取结果数据 带颜色
func (s *GenStruct) GeneratesColor() []string {
	var p generate.PrintAtom
	if config.GetIsOutSQL() {
		p.Add("\033[32;1m /******sql******\033[0m")
		p.Add(s.SQLBuildStr)
		p.Add("\033[32;1m ******sql******/ \033[0m")
	}
	p.Add("\033[32;1m " + s.Notes + " \033[0m")
	p.Add("\033[34;1m type \033[0m", s.Name, "\033[34;1m struct \033[0m {")
	mp := make(map[string]bool, len(s.Em))
	for _, v := range s.Em {
		if !mp[v.Name] {
			mp[v.Name] = true
			p.Add(" \t\t" + v.GenerateColor())
		}
	}
	p.Add(" }")

	return p.Generates()
}

//////////////////////////////////////////////////////////////////////////////
// package
//////////////////////////////////////////////////////////////////////////////

// SetPackage Defining package names.定义包名
func (p *GenPackage) SetPackage(pname string) {
	p.Name = pname
}

// AddImport Add import by type.通过类型添加import
func (p *GenPackage) AddImport(imp string) {
	if p.Imports == nil {
		p.Imports = make(map[string]string)
	}
	p.Imports[imp] = imp
}

// AddStruct Add a structure.添加一个结构体
func (p *GenPackage) AddStruct(st GenStruct) {
	p.Structs = append(p.Structs, st)
}

// Generate Get the result data.获取结果数据
func (p *GenPackage) Generate() string {
	p.genimport() // auto add import .补充 import

	var pa generate.PrintAtom
	pa.Add("package", p.Name)
	// add import
	if p.Imports != nil {
		pa.Add("import (")
		for _, v := range p.Imports {
			pa.Add(v)
		}
		pa.Add(")")
	}
	// -----------end
	// add struct
	for _, v := range p.Structs {
		for _, v1 := range v.Generates() {
			pa.Add(v1)
		}

		if config.GetIsTableName() { // add table name func
			for _, v1 := range v.GenerateTableName() {
				pa.Add(v1)
			}
		}

		if config.GetIsColumnName() {
			for _, v2 := range v.GenerateColumnName() { // add column list
				pa.Add(v2)
			}
		}
	}

	////添加增删改查结构体
	//for _, v := range p.Structs {
	//	for _, v1 := range v.GeneratesReqCreate() {
	//		pa.Add(v1)
	//	}
	//}
	//for _, v := range p.Structs {
	//	for _, v1 := range v.GeneratesRespCreate() {
	//		pa.Add(v1)
	//	}
	//}
	//for _, v := range p.Structs {
	//	for _, v1 := range v.GeneratesReqDelete() {
	//		pa.Add(v1)
	//	}
	//}
	//for _, v := range p.Structs {
	//	for _, v1 := range v.GeneratesRespDelete() {
	//		pa.Add(v1)
	//	}
	//}
	//for _, v := range p.Structs {
	//	for _, v1 := range v.GeneratesReqUpdate() {
	//		pa.Add(v1)
	//	}
	//}
	//for _, v := range p.Structs {
	//	for _, v1 := range v.GeneratesRespUpdate() {
	//		pa.Add(v1)
	//	}
	//}
	//for _, v := range p.Structs {
	//	for _, v1 := range v.GeneratesReqList() {
	//		pa.Add(v1)
	//	}
	//}
	//for _, v := range p.Structs {
	//	for _, v1 := range v.GeneratesRespList() {
	//		pa.Add(v1)
	//	}
	//}
	// -----------end

	// add func
	for _, v := range p.FuncStrList {
		pa.Add(v)
	}
	// -----------end

	// output.输出
	strOut := ""
	for _, v := range pa.Generates() {
		strOut += v + "\n"
	}

	return strOut
}
func (p *GenPackage) GenerateReqFile() string {
	p.genimport() // auto add import .补充 import

	var pa generate.PrintAtom
	pa.Add("package", "request")
	// add import
	if p.Imports != nil {
		pa.Add("import (")
		pa.Add(")")
	}

	//添加增删改查结构体
	for _, v := range p.Structs {

		for _, v1 := range v.GeneratesReqCreate() {
			pa.Add(v1)
		}

		for _, v1 := range v.GeneratesReqDelete() {
			pa.Add(v1)
		}

		for _, v1 := range v.GeneratesReqUpdate() {
			pa.Add(v1)
		}

		for _, v1 := range v.GeneratesReqList() {
			pa.Add(v1)
		}
	}

	// output.输出
	strOut := ""
	for _, v := range pa.Generates() {
		strOut += v + "\n"
	}

	return strOut
}
func (p *GenPackage) GenerateRspFile() string {
	p.genimport() // auto add import .补充 import

	var pa generate.PrintAtom
	pa.Add("package", "response")
	// add import
	if p.Imports != nil {
		pa.Add("import (")
		pa.Add("`time`")
		pa.Add(")")
	}

	for _, v := range p.Structs {
		for _, v1 := range v.GeneratesRespCreate() {
			pa.Add(v1)
		}
		for _, v1 := range v.GeneratesRespDelete() {
			pa.Add(v1)
		}
		for _, v1 := range v.GeneratesRespUpdate() {
			pa.Add(v1)
		}
		for _, v1 := range v.GeneratesRespList() {
			pa.Add(v1)
		}
	}

	// output.输出
	strOut := ""
	for _, v := range pa.Generates() {
		strOut += v + "\n"
	}

	return strOut
}
func (p *GenPackage) GenerateServices() string {
	p.genimport() // auto add import .补充 import

	var pa generate.PrintAtom
	pa.Add("package", "services")
	// add import
	if p.Imports != nil {
		pa.Add("import (")
		pa.Add("\"github.com/qiafan666/gotato/commons\"")
		pa.Add("\"sync\"")
		pa.Add(")")
	}

	pa.Add("\n")
	pa.Add("// " + "GenService service layer interface")
	pa.Add("type", "GenService", "interface {")
	for _, v := range p.Structs {
		for _, v1 := range v.GeneratesService() {
			pa.Add(v1)
		}
	}
	pa.Add("}")

	pa.Add("\n")
	pa.Add("var genServiceIns *genServiceImp")
	pa.Add("var genServiceInitOnce sync.Once")
	pa.Add("\n")
	pa.Add("func NewGenServiceInstance() GenService {")
	pa.Add("\n")
	pa.Add("	genServiceInitOnce.Do(func() {")
	pa.Add("		genServiceIns = &genServiceImp{")
	pa.Add("			dao: dao.Instance(),")
	pa.Add("		}")
	pa.Add("	})")
	pa.Add("\n")
	pa.Add("	return genServiceIns")
	pa.Add("}")
	pa.Add("\n")
	pa.Add("type genServiceImp struct {")
	pa.Add("	dao dao.Dao")
	pa.Add("}")
	// -----------end

	//接口实现类
	pa.Add("\n")

	for _, v := range p.Structs {
		for _, v1 := range v.GeneratesImpl() {
			pa.Add(v1)
		}
	}

	// output.输出
	strOut := ""
	for _, v := range pa.Generates() {
		strOut += v + "\n"
	}

	return strOut
}

// AddFuncStr add func coding string.添加函数串
func (p *GenPackage) AddFuncStr(src string) {
	p.FuncStrList = append(p.FuncStrList, src)
}

// compensate and import .获取结果数据
func (p *GenPackage) genimport() {
	for _, v := range p.Structs {
		for _, v1 := range v.Em {
			if v2, ok := cnf.EImportsHead[v1.Type]; ok {
				if len(v2) > 0 {
					p.AddImport(v2)
				}
			}
		}
	}
}
func centerString(content, boundaryChar string, totalLength int) string {
	// 确保 content 的长度不超过 totalLength
	if len(content) > totalLength {
		content = content[:totalLength]
	}

	// 计算左右边界的填充长度
	paddingLength := (totalLength - len(content)) / 2

	// 确保左右边界相等，如果总长度为奇数，右边界多一个字符
	leftPadding := strings.Repeat(boundaryChar, paddingLength)
	rightPadding := strings.Repeat(boundaryChar, totalLength-len(content)-paddingLength)

	// 拼接字符串
	return leftPadding + content + rightPadding
}
