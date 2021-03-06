package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go/format"
	//"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/btracey/su2tools/config/common"
)

var pkgName string = "config"
var optionStart string = "Option Name:"
var nameBytes []byte = []byte("Option Name:")
var typeBytes []byte = []byte("Option Type:")
var categoryBytes []byte = []byte("Option Category:")
var valuesBytes []byte = []byte("Option values:")
var defaultBytes []byte = []byte("Option default:")
var descriptionBytes []byte = []byte("Option description:")

var dynamicGeneration []byte = []byte("// This file is autogenerated using github.com/btracey/su2tools/config_writer/parser/write_options_structs.go based on the output from parse_config.py\n")

// format_file calls gofmt on the file
func format_file(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	b, err = format.Source(b)
	if err != nil {
		return err
	}
	f.Close()
	f, err = os.Create(name)
	if err != nil {
		return err
	}
	f.Write(b)
	f.Close()
	return nil
}

// CodeGenerator is something that can turn the list of options into a file
type CodeGenerator interface {
	init() error
	finalize()
	add_option(*pythonOption) error
	GetFilename() string
}

// to_camel_case turns a string ToCamelCase
func to_camel_case(s string) string {
	s = strings.ToLower(s)

	// Capitalize the first letter
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	// Go through the field and find underscores. If they exist, capitalize the next letter and remove
	// it from the string
	r2 := make([]rune, 0, len(r))
	inds := make([]int, 0)
	for i := 0; i < len(r); i++ {
		if r[i] == '_' {
			inds = append(inds, i)
			if i != len(r)-1 {
				r[i+1] = unicode.ToUpper(r[i+1])
			}
		} else {
			r2 = append(r2, r[i])
		}
	}
	s = string(r2)
	return s
}

// parseOption reads an option from the output of parse_config.py and turns it
// into a pythonOption
func parseOption(scanner *bufio.Scanner) (option *pythonOption, err error) {

	option = &pythonOption{}

	name, err := parse_line(scanner.Bytes(), nameBytes)
	if err != nil {
		return nil, errors.New("parseoption name: " + err.Error())
	}
	option.configName = common.ConfigfileOption(name)
	scanner.Scan()

	s, err := parse_line(scanner.Bytes(), typeBytes)
	if err != nil {
		return nil, errors.New("parseoption type for " + name + ": " + err.Error())
	}
	option.configType = common.ConfigOptionType(s)
	scanner.Scan()

	s, err = parse_line(scanner.Bytes(), categoryBytes)
	if err != nil {
		return nil, errors.New("parseoption category for " + name + ": " + err.Error())
	}
	option.configCategory = common.ConfigCategory(s)
	scanner.Scan()

	s, err = parse_line(scanner.Bytes(), valuesBytes)
	if err != nil {
		return nil, errors.New("parseoption enumvalues for " + name + ": " + err.Error())
	}
	option.enumOptionsString = s
	scanner.Scan()

	s, err = parse_line(scanner.Bytes(), defaultBytes)
	if err != nil {
		return nil, errors.New("parseoption default for " + name + ": " + err.Error())
	}
	option.defaultString = s
	scanner.Scan()

	s, err = parse_line(scanner.Bytes(), descriptionBytes)
	if err != nil {
		return nil, errors.New("parseoption description for " + name + ": " + err.Error())
	}
	option.description = s
	scanner.Scan() // read empty line
	return option, nil
}

func parse_line(scannerString, prefix []byte) (val string, err error) {
	if !bytes.Equal(scannerString[:len(prefix)], prefix) {
		err = errors.New("prefix mismatch")
		return "", err
	}
	// copy the bytes so that if we overwrite ...
	val = string(scannerString[len(prefix):])
	// trim whitespace
	val = strings.TrimSpace(val)
	return val, nil
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func initializeFile(file *os.File) error {
	_, err := file.WriteString("package " + pkgName + "\n")
	if err != nil {
		return err
	}
	file.Write(dynamicGeneration)
	return nil
}

func main() {
	pythonOptionsFilename := "su2options.txt"

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		exit(errors.New("gopath not found"))
	}
	basepath := filepath.Join(gopath, "src", "github.com", "btracey", "su2tools", "config")
	thispath := filepath.Join(basepath, "code_generator")
	newpath := filepath.Join(thispath, "newfiles")
	//alloptionsFilename := "all_options.go"

	// Open relevant files
	pythonOptionsFile, err := os.Open(filepath.Join(thispath, pythonOptionsFilename))
	defer pythonOptionsFile.Close()
	if err != nil {
		exit(err)
	}

	generators := []CodeGenerator{
		&optionFile{Filename: filepath.Join(newpath, "all_options.go")},
		&defaultWriter{Filename: filepath.Join(newpath, "options_default.go")},
		&optionMapWriter{Filename: filepath.Join(newpath, "option_map.go")},
		&fieldMapWriter{Filename: filepath.Join(newpath, "go_to_su2_field_map.go")},
		&su2ToGoFieldMapWriter{Filename: filepath.Join(newpath, "su2_to_go_field_map.go")},
	}

	headerOrderFilename := filepath.Join(newpath, "heading_order.go")
	headerOrder, err := os.Create(headerOrderFilename)

	initializeFile(headerOrder)
	headerOrder.WriteString("import \"github.com/btracey/su2tools/config/common\"\n")
	headerOrder.WriteString("var categoryOrder map[common.ConfigCategory]int = map[common.ConfigCategory]int{\n")

	scanner := bufio.NewScanner(pythonOptionsFile)

	// Initialize writing files
	for i := range generators {
		err := generators[i].init()
		if err != nil {
			exit(errors.New("error initializing: " + err.Error()))
		}
	}

	var headingNumber int
	// Write headings until we reach options
	for scanner.Scan() {
		b := scanner.Bytes()
		line := string(b)
		if len(line) > len(optionStart) {
			if line[:len(optionStart)] == optionStart {
				break
			}
		}
		headerOrder.WriteString("\"" + line + "\":" + strconv.Itoa(headingNumber) + ",\n")
		headingNumber++
	}
	headerOrder.WriteString("}")
	headerOrder.Close()

	// stupid loop and a half problem
	option, err := parseOption(scanner)
	if err != nil {
		exit(errors.New("error parsing option" + err.Error()))
	}
	err = option.Process()
	if err != nil {
		exit(err)
	}
	// Now that we have the option, use it
	for i := range generators {
		err := generators[i].add_option(option)
		if err != nil {
			exit(errors.New("error adding option" + err.Error()))
		}
	}

	// Read in the options
	for scanner.Scan() {
		option, err := parseOption(scanner)
		if err != nil {
			exit(errors.New("error parsing option" + err.Error()))
		}
		err = option.Process()
		if err != nil {
			exit(err)
		}
		// Now that we have the option, use it
		for i := range generators {
			err := generators[i].add_option(option)
			if err != nil {
				exit(errors.New("error adding option" + err.Error()))
			}
		}
	}

	// Finish writing files
	for _, g := range generators {
		g.finalize()
	}

	// Format the go files
	for _, g := range generators {
		err := format_file(g.GetFilename())
		if err != nil {
			exit(errors.New("error formatting: " + err.Error()))
		}
	}
	err = format_file(headerOrderFilename)
	if err != nil {
		exit(errors.New("error formatting header: " + err.Error()))
	}

	// Move the go files up a directory
	for _, g := range generators {
		// get the trailing part of the name
		oldname := g.GetFilename()
		newname := filepath.Join(basepath, filepath.Base(g.GetFilename()))
		err := os.Rename(oldname, newname)
		if err != nil {
			exit(errors.New("error moving: " + err.Error()))
		}
	}
	oldname := headerOrderFilename
	newname := filepath.Join(basepath, filepath.Base(headerOrderFilename))
	err = os.Rename(oldname, newname)
	if err != nil {
		exit(errors.New("error moving: " + err.Error()))
	}
}
