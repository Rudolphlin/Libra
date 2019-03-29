/* ######################################################################
# File Name: md5Photo.go
# Author: Laiyinglin
# Main: Rudolph_Lin@hotmail.com
# Created Time: 2019-03-11 10:27:30
####################################################################### */
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	ReSetPhotoName()
}
func ReSetPhotoName() {
	list := os.Args
	fmt.Println(len(list))
	for _, photoFolder := range list {
		fmt.Println(photoFolder)
		files, err := ioutil.ReadDir(photoFolder)
		if err == nil {
			for _, file := range files {
				if file.IsDir() {
					continue
				} else {
					fileName := file.Name()
					fmt.Println(photoFolder + "/" + fileName)
					newFileName := GetMD5Hash(photoFolder + "/" + fileName)

					doIndex := strings.LastIndex(fileName, ".")
					if doIndex != -1 && doIndex != 0 {
						newFileName += fileName[doIndex:]
						fmt.Println(newFileName)
					}

					reNameErr := os.Rename(photoFolder+"/"+fileName, photoFolder+"/"+newFileName)
					if reNameErr != nil {
						fmt.Println("rename fail , err is", reNameErr)
						break
					}
				}
			}
		} else {
			fmt.Println("err = ", err)
		}
	}
}

func GetMD5Hash(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("file open fail , err is  ", err)
		return ""
	}
	defer f.Close()

	md5Hash := md5.New()
	if _, err3 := io.Copy(md5Hash, f); err3 != nil {
		fmt.Println("md5 copy fail err ", err3)
		return ""
	}

	str := md5Hash.Sum(nil)
	return hex.EncodeToString(str)
}
