// by gpt4o mini @ 240801
// https://chatgpt.com/share/943bfc66-0ba6-45aa-b680-9abbfbc02d36

package debug

import (
	"fmt"
	"log"
	"reflect"
	"runtime"

	"github.com/eh-web-viewer/eh-web-viewer/Tools/orderedmap"
	"github.com/fatih/color"
)

func OrderedMap(o *orderedmap.OrderedMap) {
	if o == nil {
		fmt.Printf("%v\n", o)
	} else {
		orderedMap(*o, "")
	}
}

func orderedMap(o orderedmap.OrderedMap, intent string) {
	for _, k := range o.Keys() {
		v, _ := o.Get(k)
		if oo, ok := v.(orderedmap.OrderedMap); ok {
			fmt.Printf("%s%s : \n", intent, k)
			orderedMap(oo, intent+"  ")
		} else if ol, ok := v.([]any); ok {
			fmt.Printf("%s%s : \n", intent, k)
			for _, olv := range ol {
				if oo, ok := olv.(orderedmap.OrderedMap); ok {
					orderedMap(oo, intent+"  ")
				} else {
					fmt.Printf("%s  %v\n", intent, olv)
				}
			}
		} else {
			fmt.Printf("%s%s : %v\n", intent, k, v)
		}
	}
}

func DeepPrint(v any, indent string) {
	rv := reflect.ValueOf(v)

	if !rv.IsValid() {
		fmt.Println(indent + "Invalid")
		return
	}

	// Handle pointers by recursively dereferencing
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			fmt.Println(indent + "Nil Pointer")
			return
		}
		rv = rv.Elem() // Dereference the pointer
	}

	switch rv.Kind() {
	case reflect.Struct:
		fmt.Println(indent + "Struct:")
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			fieldName := rv.Type().Field(i).Name
			fmt.Printf("%sField %s: ", indent, fieldName)
			DeepPrint(field.Interface(), indent+indent)
		}
	case reflect.Slice:
		fmt.Println(indent + "Slice:")
		for i := 0; i < rv.Len(); i++ {
			fmt.Printf("%sElement %d:\n", indent, i)
			DeepPrint(rv.Index(i).Interface(), indent+indent)
		}
	case reflect.Map:
		fmt.Println(indent + "Map:")
		for _, key := range rv.MapKeys() {
			value := rv.MapIndex(key)
			fmt.Printf("%sKey %v: ", indent, key.Interface())
			DeepPrint(value.Interface(), indent+indent)
		}
	default:
		fmt.Println(rv.Interface())
	}
}

type logLevel int

const (
	Trace logLevel = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

func levelToString(level logLevel) string {
	switch level {
	case Trace:
		return logFunctions[Trace]("TRACE")
	case Debug:
		return logFunctions[Debug]("DEBUG")
	case Info:
		return logFunctions[Info]("INFO")
	case Warn:
		return logFunctions[Warn]("WARN")
	case Error:
		return logFunctions[Error]("ERROR")
	case Fatal:
		return logFunctions[Fatal]("Fatal")
	default:
		return logFunctions[Fatal]("UNKNOWN")
	}
}

var LogLevel logLevel = Trace
var logFunctions = map[logLevel]func(a ...interface{}) string{
	Trace: color.New(color.FgBlack, color.BgCyan, color.Bold).SprintFunc(),
	Debug: color.New(color.FgBlack, color.BgBlue, color.Bold).SprintFunc(),
	Info:  color.New(color.FgBlack, color.BgGreen, color.Bold).SprintFunc(),
	Warn:  color.New(color.FgBlack, color.BgYellow, color.Bold).SprintFunc(),
	Error: color.New(color.FgBlack, color.BgRed, color.Bold).SprintFunc(),
	Fatal: color.New(color.FgBlack, color.BgRed, color.Bold).SprintFunc(),
}

func T(tag any, msg ...any) {
	if LogLevel > Trace {
		return
	}
	log.Printf("[%s] %s %s", tag, levelToString(Trace), fmt.Sprint(msg...))
}

func D(tag any, msg ...any) {
	if LogLevel > Debug {
		return
	}
	log.Printf("[%s] %s %s", tag, levelToString(Debug), fmt.Sprint(msg...))
}

func I(tag any, msg ...any) {
	if LogLevel > Info {
		return
	}
	log.Printf("[%s] %s %s", tag, levelToString(Info), fmt.Sprint(msg...))
}

func W(tag any, msg ...any) {
	if LogLevel > Warn {
		return
	}
	log.Printf("[%s] %s %s", tag, levelToString(Warn), fmt.Sprint(msg...))
}

func E(tag any, msg ...any) {
	if LogLevel > Error {
		return
	}
	log.Printf("[%s] %s %s", tag, levelToString(Error), fmt.Sprint(msg...))
}

func F(tag any, msg ...any) {
	log.Printf("[%s] %s %s", tag, levelToString(Fatal), fmt.Sprint(msg...))
}

func Print(v ...any) {
	// 获取调用栈信息
	_, file, line, ok := runtime.Caller(1)
	if ok {
		D(fmt.Sprintf("%s:%d", file, line), v...)
	}
}
