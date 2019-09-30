package utils

import (
	"bytes"
	"fmt"
	"github.com/tealeg/xlsx"
	"net/http"
	"reflect"
	"strings"
	"time"
)

//创建Excel数据
//函数传入一个[]interface{}数组，必须在传入之前就已经转换成[]interface{},否则报错
//cells :excel第一行数据，标题
//sn: sheet页的名字，如果不写的话，默认是Sheet1
//inters: Excel的数据
func CreateExcel(cells []string, sn string, inters []interface{}) (file *xlsx.File) {
	var row *xlsx.Row
	var cell *xlsx.Cell
	//创建Excel第一条数据：标题
	file = xlsx.NewFile()
	var sheetName string
	if len(strings.TrimSpace(sn)) == 0 {
		sheetName = "Sheet1"
	} else {
		sheetName = sn
	}
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		fmt.Println("创建sheet出错：", err)
	}
	row = sheet.AddRow()
	for _, v := range cells {
		cell = row.AddCell()
		cell.Value = v
	}

	//创建Excel主体
	if len(inters) > 0 {
		for _, v := range inters {
			//excel添加一个新行
			row = sheet.AddRow()
			//利用反射判断出v的类型，有可能v是基本类型，所以只是简单的判断了一下四个基本类型：string,float64,float32,int
			//其他类型默认是struct
			getType := reflect.TypeOf(v)
			strGetType := fmt.Sprintf("%s", getType)
			getValue := reflect.ValueOf(v)
			if strGetType == "string" {
				//添加一个新的单元格
				cell = row.AddCell()
				cell.String()
				cell.Value = fmt.Sprintf("%s", v)
			} else if strGetType == "float64" {
				cell = row.AddCell()
				cell.String()
				cell.Value = fmt.Sprintf("%1.2f", v)
			} else if strGetType == "float32" {
				cell = row.AddCell()
				cell.String()
				cell.Value = fmt.Sprintf("%1.2f", v)
			} else if strGetType == "int" {
				cell = row.AddCell()
				cell.String()
				cell.Value = fmt.Sprintf("%d", v)
			} else if strGetType == "map[string]interface {}" {
				keys := getValue.MapKeys()
				for _, v := range keys {
					mapValue := getValue.MapIndex(v)
					mapValueType := fmt.Sprintf("%s", mapValue.Kind())
					//添加新的单元
					cell = row.AddCell()
					cell.String()
					switch {
					case mapValueType == "string":
						cell.Value = fmt.Sprintf("%s", mapValue)
					case mapValueType == "float64":
						cell.Value = fmt.Sprintf("%1.2f", mapValue)
					case mapValueType == "float32":
						cell.Value = fmt.Sprintf("%1.2f", mapValue)
					case mapValueType == "int":
						cell.Value = fmt.Sprintf("%d", mapValue)
					case mapValueType == "interface":
						mapV := fmt.Sprintf("%s", mapValue.Interface())
						if strings.Contains(mapV, "%") {
							mapV = fmt.Sprintf("%1.2f", mapValue)
						}
						if strings.Contains(mapV, "%") {
							mapV = fmt.Sprintf("%d", mapValue)
						}
						cell.Value = mapV
					default:
						cell.Value = fmt.Sprintf("%s", mapValue)
					}
				}
			} else {
				//利用反射判断struct有几个属性，创建几个cell,判断struct中的属性是什么类型，并把数据填写到excel中
				for i := 0; i < getType.NumField(); i++ {
					value := getValue.FieldByName(getType.Field(i).Name)
					fileType := getType.Field(i).Type
					sFileType := fmt.Sprintf("%s", fileType)
					cell = row.AddCell()
					switch {
					case sFileType == "string":
						cell.Value = fmt.Sprintf("%s", value)
					case sFileType == "float64":
						cell.Value = fmt.Sprintf("%1.2f", value)
					case sFileType == "float32":
						cell.Value = fmt.Sprintf("%1.2f", value)
					case sFileType == "int":
						cell.Value = fmt.Sprintf("%d", value)
					default:
						cell.Value = fmt.Sprintf("%s", value)
					}

				}
			}
		}
	}
	return
}

//导出Excel
// 文件名称以毫秒数命名
//函数传入一个[]interface{}数组，必须在传入之前就已经转换成[]interface{},否则报错
//cells :excel第一行数据，标题
//sheetname: sheet页的名字，如果不写的话，默认是Sheet1
//inters: Excel的数据
//write 页面的responseWrite
//request  页面的http.request请求

func ExportExcel(cells []string, sheetname string, inters []interface{}, write http.ResponseWriter, request *http.Request) {
	file := CreateExcel(cells, sheetname, inters)
	filename := fmt.Sprintf("%d.xlsx", time.Now().Unix()) //文件名是毫秒数
	//导出excel
	write.Header().Add("Content-Disposition", "attachment;filename="+filename)
	write.Header().Add("Content-Type", "application/vnd.ms-excel")
	write.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	var buffer bytes.Buffer
	if err := file.Write(&buffer); err != nil {
		panic(err)
		return
	}
	r := bytes.NewReader(buffer.Bytes())

	http.ServeContent(write, request, filename, time.Now(), r)
}
