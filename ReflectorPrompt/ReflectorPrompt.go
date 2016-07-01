package ReflectorPrompt

import (
    "fmt"
    "reflect"
    "strings"
    "bufio"
    "os"
)

type ErrPointerRequired struct{}
type ErrUndefinedType struct{}

func (e ErrPointerRequired) Error() string {
    return fmt.Sprint("passed variable is not a pointer")
}

func (e ErrUndefinedType) Error() string {
    return fmt.Sprintf("unknown type passed as variable pointer")
}

func Prompter(prompt string, variable interface{}) error {
    fmt.Print(prompt)
    // Determine the type of the variable
    v := reflect.ValueOf(variable)
    // Ensure variable is a pointer
    if v.Kind().String() != "ptr" {
        return ErrPointerRequired{}
    }

    s := v.Type().String()

    // Test for type of pointer and create appropriate variable and receiver format string
    switch {
    case strings.Contains(s, "float"):
        recv := new(float64)
        formatStr := "%f "
        fmt.Scanf(formatStr, recv)
        v.Elem().SetFloat(*recv)
    case strings.Contains(s, "int"):
        recv := new(int64)
        formatStr := "%d "
        fmt.Scanf(formatStr, recv)
        v.Elem().SetInt(*recv)
    case strings.Contains(s, "string"):
        rd := bufio.NewReader(os.Stdin)
        str, err := rd.ReadString('\n')
        if err != nil {
            return err
        }
        v.Elem().SetString(strings.TrimSpace(str))
    default:

        return ErrUndefinedType{}
    }

    return nil
}
