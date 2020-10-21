package filedir

import (
	"bufio"
	"container/list"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*CheckFileIsExist 判断文件是否存在
 * Return  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/*WriteFileContent 写入文件内容，目录|文件不存在则创建目录|文件
 * Return  存在返回 true 不存在返回false
 */
func WriteFileContent(filename string, content string,append bool) (bool, error) {
	dir:=filepath.Dir(filename)
	if !CheckFileIsExist(dir) {
		os.MkdirAll(dir,os.ModePerm)
	}
	//如果不是追加模式，则删除旧文件再写入
	if !append {
		os.Remove(filename)
	}
	var flag int
	if append {
		flag= os.O_RDWR|os.O_CREATE|os.O_APPEND
	}else {
		flag= os.O_RDWR|os.O_CREATE
	}
	outputFile, err := os.OpenFile(filename, flag, 0666)
	if err != nil {
		return false, err
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	defer outputWriter.Flush()
	//写入内容
	outputWriter.WriteString(content)
	return true, nil
}


/*TraverseDir 递归文件夹获取到所有文件名称
 *
 *  dirPth 目录
 */
func TraverseDir(dirPth string, fileList *list.List) error {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			TraverseDir(dirPth+PthSep+fi.Name(), fileList)
		} else {
			fileList.PushBack(dirPth + PthSep + fi.Name())
		}
	}
	return nil
}

/*TraverseDir 递归文件夹获取到所有文件名称
 *
 *  dirPth 目录
 */
func TraverseDirBySlice(dirPth string) ([]string, error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)

	var curFile []string
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			down, _ := TraverseDirBySlice(dirPth + PthSep + fi.Name())
			curFile = append(curFile, down...)
		} else {
			curFile = append(curFile, dirPth+PthSep+fi.Name())
		}
	}

	return curFile, nil
}
