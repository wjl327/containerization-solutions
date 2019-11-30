package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	opers := []string{"点击首页", "拉取推荐页", "修改资料", "查看动态", "发表动态",
		"发表评论", "发弹幕", "送礼物", "刷小视频"}
	for {
		sleepTime := 300 + rand.Int31n(1000)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		userId := rand.Int63n(100000)
		oper := opers[rand.Int31n(int32(len(opers)))]
		now := time.Now()
		logTime := now.Format("2006-01-02 15:04:05")
		status := "succ"
		if rand.Int31n(100) < 10 {
			status = "fail"
		}
		res := logTime + "|" + strconv.Itoa(int(userId)) + "|" + oper + "|" + status + "\n"
		appendFile("../log/helloElk.log", res)
	}

}

func appendFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
