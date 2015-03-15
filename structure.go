package structrecord

import (
	"fmt"
	"reflect"
	"strings"
)

// reister server.
// example : mysql sqlit mssql memcached redis mongodb for provider ORM
var provider map[string]interface{}

type StructureInterface interface {
	Save(model interface{}) (int, bool) //Saved back ID,otherwise return false
	Find(index int) []interface{}
	FindAll() []interface{}
	FindBy(fieldname string) []interface{}
	Update(model interface{}) (int, bool) //Updates on successful return number of affected rows,Failed to update return false
	Delete(model interface{}) bool        // successful return true,otherwise return false
	Begin()
	Commit([]interface{}) int //
	Rollback()
}

type Structure struct {
	TableName string
	Relat     map[int]string
	Target    map[int]string
	StructureInterface
}

// read struct infomation
// get congruent relationship and table name
func (s *Structure) readStructInfo(model interface{}) Structure {
	val := reflect.ValueOf(model)

	ind := reflect.Indirect(val)

	numField := ind.NumField()
	name := snakeString(ind.Type().Name())
	rela := make(map[int]string)
	tget := make(map[int]string)
	for i := 0; i < numField; i++ {
		rela[i] = ind.Type().Field(i).Tag.Get("relat")

		tget[i] = ind.Type().Field(i).Tag.Get("target")
	}
	return Structure{TableName: name, Relat: rela, Target: tget}

}

func (s *Structure) Save(model interface{}) (i int, b bool) {
	md := s.readStructInfo(model)
	fmt.Println(model)
	fmt.Println(md)
	return 1, true

}

// snake string, XxYy to xx_yy
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:len(data)]))
}

func (st *Structure) Bootstrap() Structure {
	var s Structure = Structure{}
	return s
}

func Run() StructureInterface {

	st := new(Structure)

	return st
}
