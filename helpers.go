package gologger

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// toJSON из слайса []any формирует строку
// Формат слайса fieldsAndValues должен быть в формате:
// ["ИмяПоля1", Значение1, "ИмяПоля2", Значение2...]
//
// Желательно, чтобы значения реализовывали интерфейс fmt.Stringer
func toJSON(fieldsAndValues []any) []byte {
	buf := bytes.Buffer{}
	//builder := strings.Builder{}

	//builder.WriteString("{")
	buf.WriteRune('{')
	for i := 1; i < len(fieldsAndValues); i += 2 {
		keyField, _ := fieldsAndValues[i-1].(string)

		buf.WriteRune('"')
		//builder.WriteString(`"`)
		buf.WriteString(keyField)
		//builder.WriteString(keyField)
		buf.WriteRune('"')
		//builder.WriteString(`"`)
		//builder.WriteString(`:`)
		buf.WriteRune(':')

		switch v := fieldsAndValues[i].(type) {
		case string:
			buf.WriteRune('"')
			//builder.WriteString(`"`)

			buf.WriteString(v)
			//builder.WriteString(v)
			buf.WriteRune('"')
			//builder.WriteString(`"`)
		case fmt.Stringer:
			buf.WriteRune('"')
			//builder.WriteString(`"`)
			buf.WriteString(v.String())
			//builder.WriteString(keyField)
			//builder.WriteString(v.String())
		default:

			vB, err := json.Marshal(v)
			if err != nil {
				buf.WriteString(err.Error())
				//builder.WriteString(err.Error())
				break
			}

			// Экранируем json объект
			vb := bytes.Replace(vB, []byte(`"`), []byte(`\"`), -1)
			buf.WriteRune('"')
			buf.Write(vb)
			//builder.Write(vb)
			buf.WriteRune('"')
		}

		if i < len(fieldsAndValues)-2 {
			//builder.WriteString(",")
			buf.WriteRune(',')
		}
	}

	buf.WriteRune('}')
	//builder.WriteString("}")

	buf.WriteString("\r\n")
	return buf.Bytes()
}
