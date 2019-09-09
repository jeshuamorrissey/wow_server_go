package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/jeshuamorrissey/wow_server_go/dbc_tool/dbc"
)

const (
	// BinDBCDir is the directory containing the binary DBC files.
	BinDBCDir = "D:\\Games\\World of Warcraft (Vanilla) - Tools\\DBC"

	// JSONDBCDir is the directory containing the JSON DBC files.
	JSONDBCDir = "D:\\Users\\Jeshua\\go\\src\\github.com\\jeshuamorrissey\\wow_server_go\\dbc_tool\\dbc\\data"
)

func main() {
	// data, err := ioutil.ReadFile(filepath.Join(BinDBCDir, "ChrClasses.dbc"))
	// if err != nil {
	// 	panic(err)
	// }

	// table, err := utils.LoadBinaryDBC(data, new(dbc.Class))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v", table.Records[0])

	classes, err := dbc.LoadClasses(JSONDBCDir)
	if err != nil {
		panic(err)
	}

	// Save the classes to a DBC file.
	classesBinary := classes.ToBinary()
	err = ioutil.WriteFile(filepath.Join(BinDBCDir, "ChrClasses.dbc"), classesBinary, 0777)
	if err != nil {
		panic(err)
	}
}
