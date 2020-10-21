package filedir

import (
	"container/list"
	"io/ioutil"
	"os"
)

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
