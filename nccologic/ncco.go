package nccologic

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type AnswerMessage struct {
	Action    string
	VoiceName string
	Text      string
}

func BuildAnswerURL(as AnswerMessage) {
	//baseDir := "/data/wwwroot/falcon.web.sdp.101.com/webroot/"
	baseDir := "E:\\tmp\\ncco\\"
	dateTime := time.Now().Format("2006-01-02")
	fileDir := baseDir + dateTime
	fmt.Println(fileDir)

	timeNow := time.Now().Unix()
	timeNowStr := strconv.FormatInt(timeNow, 10)
	//filePath := fileDir + timeNowStr + ".json"
	filePath := fileDir + "\\" + timeNowStr + ".json"
	fmt.Println(filePath)

	err := os.MkdirAll(fileDir, 0666)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println("Create Directory OK!")
	}

	jsonData, _ := json.Marshal(as)
	fmt.Println(string(jsonData))

	outfile, err := os.Create(filePath)
	defer outfile.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = outfile.Write(jsonData)
	if err != nil {
		log.Fatal(err)
		return
	}
}
