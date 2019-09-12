package yoso

import "os"

type Dep struct {
	Cfg    Config
	Res    Result
	File   *os.File
	Writer CSV

	//couter for partition file
	PartCounter int

	//counter for row limiter per file
	LimitCounter int

	//get all files fullpath
	ResultFiles []string
}

type Config struct {
	// Path is the location of file that will be generated.
	Path string

	// Make sure separator that being used isn't found
	// inside data that will be write into file.
	Separator rune

	// FileName is design to be only contain file name in it,
	// without file path and file extension.
	FileName string

	// Header is used once when creating file
	// If len of Header is 0 then it means header will not
	// be written in file
	Header []string

	// If UsePart is true then after Path+FileName
	// will be add "_part"+PartCounter.
	// example: temp/listStudent_part1.csv
	UsePart bool

	// LimitPerPart to determine how many row per csv file.
	// Make sure to assign LimitPerPart more than 1 if UsePart is true.
	LimitPerPart int
}

type Result struct {
	//retrieve all files data that have been written
	Files []string
}
