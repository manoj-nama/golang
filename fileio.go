package main

import (
	"fmt"
	"io/ioutil"
)
type Page struct {
	FileName string
	Data []byte
}
func (p *Page) saveFile() error{
	filename := p.FileName
	return ioutil.WriteFile(filename,p.Data,0600)
} 

func loadPage(name string)  (*Page, error) {
	data,err := ioutil.ReadFile(name)
	if err != nil {
		return nil,err
	}
	return &Page{FileName: name, Data: data},nil
}

func main() {
	p1 := &Page{FileName: "TestPage.txt",Data: []byte("sample File writer ...\r\n")}
	p1.saveFile();
	p2,err :=loadPage("TestPage.txt")
	if err!=nil{
		fmt.Println(string("Error occured"))
	}else{
		fmt.Println(string(p2.Data))	
	}
}
