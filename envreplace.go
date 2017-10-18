package main

import (
    "os"
    "regexp"
    "io/ioutil"
)


func main() {
    data, err := ioutil.ReadAll(os.Stdin);
    if err != nil {
        panic(err);
    }
    regex := regexp.MustCompile("\\$[\\w_]+");
    text := string(data);
    replacedText := regex.ReplaceAllStringFunc(text, func(variable string) string {
        value, defined := os.LookupEnv(variable[1:]);
        if defined {
            return value;
        }
        return variable;
    });
    os.Stdout.WriteString(replacedText);
}
