package gologger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// toJSON из слайса []any формирует строку
// Формат слайса fieldsAndValues должен быть в формате:
// ["ИмяПоля1", Значение1, "ИмяПоля2", Значение2...]
//
// Желательно, чтобы значения реализовывали интерфейс fmt.Stringer
func toJSON(fieldsAndValues []any) string {
	builder := strings.Builder{}

	builder.WriteString("{")
	for i := 1; i < len(fieldsAndValues); i += 2 {
		keyField, _ := fieldsAndValues[i-1].(string)

		builder.WriteString(`"`)
		builder.WriteString(keyField)
		builder.WriteString(`"`)
		builder.WriteString(`:`)

		switch v := fieldsAndValues[i].(type) {
		case string:
			builder.WriteString(`"`)
			builder.WriteString(v)
			builder.WriteString(`"`)
		case fmt.Stringer:
			builder.WriteString(`"`)
			builder.WriteString(keyField)
			builder.WriteString(v.String())
		default:

			vB, err := json.Marshal(v)
			if err != nil {
				builder.WriteString(err.Error())
				break
			}

			// Экранируем json объект
			vb := bytes.Replace(vB, []byte(`"`), []byte(`\"`), -1)
			builder.WriteString(`"`)
			builder.Write(vb)
			builder.WriteString(`"`)
		}

		if i < len(fieldsAndValues)-2 {
			builder.WriteString(",")
		}
	}

	builder.WriteString("}")

	return builder.String() + "\r\n"
}
