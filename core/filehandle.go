package core

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func MakeDir(dir string) {
	defer os.MkdirAll(dir, os.ModePerm)
}

func CreateFile(filename string) {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Printf(""+
			"[CreateFile]"+
			"create file error! filename: %s | error: %s",
			filename,
			err,
		)
	}
}

func MoveFilePath(filefullname string) {
	// 结果同步到proxy后结果文件名去掉"."
	filename := path.Base(filefullname)
	resultFileDir := beego.AppConfig.String("ResultFileDir")
	newfilefullname := fmt.Sprintf("%s/%s", resultFileDir, filename)
	err := os.Rename(filefullname, newfilefullname)
	if err != nil {
		log.Printf(
			"[MoveFile]move file faild! file:%s, newfile:%s, error:%s",
			filefullname,
			newfilefullname,
			err,
		)
	}
}

func ReMoveFile(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Printf(
			"[ReMoveFile]remove file faild! file:%s. error:%s",
			file,
			err,
		)
	}
}

func WriteResult(resultMap map[string]string, sync bool) {
	resultFileDir := beego.AppConfig.String("ResultFileDir")
	waitingSyncDir := beego.AppConfig.String("WaitingSyncResultDir")
	var filename string
	if sync {
		filename = fmt.Sprintf("%s/%s.json", resultFileDir, resultMap["tid"])
		log.Printf(
			"[WriteResult] "+
				"task(tid:%s | sync_type:sync) begin write result. filename:%s",
			resultMap["tid"],
			filename,
		)
	} else {
		filename = fmt.Sprintf("%s/%s.json", waitingSyncDir, resultMap["tid"])
		log.Printf(
			"[WriteResult] "+
				"task(tid:%s | sync_type:async) begin write result. filename:%s | status:%s",
			resultMap["tid"],
			filename,
			resultMap["status"],
		)
	}

	content, _ := json.Marshal(resultMap)
	//f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		log.Printf(
			"[WriteResult]write result file error! filename: %s | error: %s",
			filename,
			err.Error(),
		)
	} else {
		_, err := f.Write([]byte(content))
		if err != nil {
			log.Printf(
				"[WriteResult]write result file error! filename: %s | error: %s",
				filename,
				err.Error(),
			)
		} else {
			log.Printf(
				"[WriteResult]write result file success! filename: %s | task status: %s",
				filename,
				resultMap["status"],
			)
		}
	}
}

func WriteFile(filename string, content []byte) (bool, error) {
	err := ioutil.WriteFile(filename, content, 0777)
	if err != nil {
		log.Printf(
			"[WriteFile]write result file error! filename: %s | error: %s",
			filename,
			err.Error(),
		)
		return false, err
	}
	return true, nil
}

func ReadFile(filename string) string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf(
			"[writeFile]Read result file error! filename: %s | error: %s",
			filename,
			err.Error(),
		)
	}
	return string(f)
}

func GetFileNameByTid(tid string) string {
	waitingSyncDir := beego.AppConfig.String("WaitingSyncResultDir")
	filename := fmt.Sprintf(
		"%s/%s.json",
		waitingSyncDir,
		tid,
	)
	return filename
}
