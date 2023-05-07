package uniq

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
)

type Options struct {
	Num_fields int
	Num_chars  int
	CopyPtr    bool
	DoublePtr  bool
	UniqPtr    bool
	RegiPtr    bool
	InputFile  string
	OutputFile string
}

var (
	ErrTooMuchFlagArgs  = errors.New("передано слишком много аргументов")
	ErrUncombinedParams = errors.New("параллельно параметры -c, -d, -u не имеют никакого смысла")
)

func ParseOptions() (*Options, error) {
	num_fields := flag.Int("f", 0, "Не учитывать первые n полей в строке")
	num_chars := flag.Int("s", 0, "Не учитывать первые n символов в строке")
	copyPtr := flag.Bool("c", false, "Подсчитать количество встреч строки во входных данных")
	doublePtr := flag.Bool("d", false, "Вывод тех строк, которые повторились во входных данных")
	uniqPtr := flag.Bool("u", false, "Вывод тех строк, которые не повторились во входных данных")
	regiPtr := flag.Bool("i", false, "Не учитывать регистр букв")

	flag.Parse()

	if *copyPtr && *doublePtr || *copyPtr && *uniqPtr || *doublePtr && *uniqPtr {
		return nil, ErrUncombinedParams
	}

	opt := Options{}

	opt.Num_fields = *num_fields
	opt.Num_chars = *num_chars
	opt.CopyPtr = *copyPtr
	opt.DoublePtr = *doublePtr
	opt.UniqPtr = *uniqPtr
	opt.RegiPtr = *regiPtr

	if len(flag.Args()) >= 1 {
		opt.InputFile = flag.Arg(0)
	}
	if len(flag.Args()) == 2 {
		opt.InputFile = flag.Arg(1)
	}
	if len(flag.Args()) > 2 {
		return nil, ErrTooMuchFlagArgs
	}

	return &opt, nil
}

func Functional(opt *Options, writer io.Writer, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	previous, counter := scanner.Text(), 1
	for scanner.Scan() {
		current := scanner.Text()
		if StrsIsEqual(&previous, &current, opt) {
			counter++
		} else {
			if opt.CopyPtr {
				CopyHandler(writer, &previous, counter)
			} else if opt.DoublePtr {
				DoubleHandler(writer, &previous, counter)
			} else if opt.UniqPtr {
				UniqueHandler(writer, &previous, counter)
			} else if !opt.DoublePtr && !opt.UniqPtr && !opt.CopyPtr {
				fmt.Fprintln(writer, previous)
			}
			counter = 1
			previous = current
		}
	}
	fmt.Fprintln(writer, previous)
}

func StrsIsEqual(previous, current *string, opt *Options) bool {
	if opt.Num_fields > 0 {
		*previous, _ = DeleteFieldsHandler(*previous, opt)
		*current, _ = DeleteFieldsHandler(*current, opt)
	}
	if opt.Num_chars > 0 {
		*previous, _ = DeleteCharsHandler(*previous, opt)
		*current, _ = DeleteCharsHandler(*current, opt)
	}

	if *previous == *current || strings.EqualFold(*previous, *current) && opt.RegiPtr {
		return true
	}
	return false
}

func GetFields(line *string) []string {
	fields := strings.FieldsFunc(*line, func(r rune) bool {
		return r == ' '
	})
	return fields
}

func DeleteFieldsHandler(str string, opt *Options) (string, error) {
	if opt.Num_fields == 0 {
		return str, nil
	}

	fields := GetFields(&str)
	if len(fields) <= opt.Num_fields {
		return "", nil
	}

	fields = fields[opt.Num_fields:]
	result := strings.Join(fields, " ")
	return result, nil
}

func DeleteCharsHandler(str string, opt *Options) (string, error) {
	chars := str
	counter := len(chars)
	if counter <= opt.Num_chars {
		return "", nil
	}
	chars = chars[opt.Num_chars:]
	return chars, nil
}

func UniqueHandler(writer io.Writer, str *string, counter int) {
	if counter == 1 {
		fmt.Fprintln(writer, *str)
	}
}

func DoubleHandler(writer io.Writer, str *string, counter int) {
	if counter > 1 {
		fmt.Fprintln(writer, *str)
	}
}

func CopyHandler(writer io.Writer, str *string, counter int) {
	fmt.Fprintln(writer, counter, *str)
}
