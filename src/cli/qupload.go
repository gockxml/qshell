package cli

import (
	"fmt"
	"qiniu/log"
	"qshell"
	"strconv"
)

func QiniuUpload(cmd string, params ...string) int {
	if len(params) == 1 || len(params) == 2 {
		var uploadConfigFile string
		var threadCount int64
		var err error
		if len(params) == 2 {
			threadCount, err = strconv.ParseInt(params[0], 10, 64)
			if err != nil {
				log.Error("Invalid <ThreadCount> value,", params[0])
				return 1
			}
			uploadConfigFile = params[1]
		} else {
			uploadConfigFile = params[0]
		}
		if threadCount < qshell.MIN_UPLOAD_THREAD_COUNT ||
			threadCount > qshell.MAX_UPLOAD_THREAD_COUNT {
			fmt.Println("You can set <ThreadCount> value between 1 and 100 to improve speed")
			threadCount = qshell.MIN_UPLOAD_THREAD_COUNT
		}
		return qshell.QiniuUpload(int(threadCount), uploadConfigFile)
	} else {
		CmdHelp(cmd)
	}
	return 0
}
