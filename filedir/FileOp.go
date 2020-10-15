package filedir

import (
	"container/list"
	"io/ioutil"
	"os"
)

//递归文件夹获取到所有文件名称
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

//递归文件夹获取到所有文件名称
func TraverseDirBySlice(dirPth string, fileList []string) error {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			TraverseDirBySlice(dirPth+PthSep+fi.Name(), fileList)
		} else {
			fileList = append(fileList, dirPth+PthSep+fi.Name())
		}
	}
	return nil
}
