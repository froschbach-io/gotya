# gotya
GO Template tool, Yet Another one

This is an example implementation that shows basic usage of Go Templates. For more information, please refer to the [accompanying article](https://froschbach.io/articles/gotya-go-templates/).

## Build

Install dependencies:

```
$ go get gopkg.in/yaml.v2
```

Build binary:

```
$ go build -o gotya main.go
```

## Usage

Provide the values as YAML file and the template file (any text file):

```
$ ./gotya values.yaml template.txt
```
