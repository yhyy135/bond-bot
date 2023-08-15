package main

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"
)

func BondParser(data []byte) any {
	var bond any
	json.NewDecoder(&JsonpWrapper{
		Underlying: bytes.NewBuffer(data),
		Prefix:     "_",
	}).Decode(&bond)

	return bond
}

func BondFilter(data any) string {
	type Bonds struct {
		Result struct {
			Data []struct {
				Name   string `json:"SECURITY_NAME_ABBR"`
				Code   string `json:"SECURITY_CODE"`
				Date   string `json:"VALUE_DATE"`
				Rating string `json:"RATING"`
			} `json:"data"`
		} `json:"result"`
	}
	var (
		message string = ""
		bonds   Bonds
	)
	json.Unmarshal(func(data any) []byte {
		b, _ := json.Marshal(data)
		return b
	}(data), &bonds)

	var today, tomorrow []string
	var timezone, _ = time.LoadLocation("Asia/Shanghai")
	for _, v := range bonds.Result.Data {
		// 匹配今天
		if v.Date == time.Now().In(timezone).Format("2006-01-02")+" 00:00:00" {
			today = append(today, "🆕 "+v.Name+"("+v.Code+") 🔝"+v.Rating+"\n")
		}

		// 匹配明天
		if v.Date == time.Now().In(timezone).Add(time.Hour*24).Format("2006-01-02")+" 00:00:00" {
			tomorrow = append(tomorrow, "🆕 "+v.Name+"("+v.Code+") 🔝"+v.Rating+"\n")
		}
	}

	if len(today) > 0 {
		message += "\n🎉今天别错过" + parseNum2Emoji(len(today)) + "\n"
		for _, val := range today {
			message += val
		}
	}

	if len(tomorrow) > 0 {
		message += "\n😁明天有戏" + parseNum2Emoji(len(tomorrow)) + "\n"
		for _, val := range tomorrow {
			message += val
		}
	}

	if len(message) == 0 {
		message = "这两天啥都木有，好好打工静候捡钱🤑"
	} else {
		message = "🔔新债提醒🔔\n" + message
	}
	return message
}

func parseNum2Emoji(num int) string {
	numStr := strconv.Itoa(num)
	emojiNum := ""
	for _, char := range numStr {
		s := string(char)
		emojiDigit := string(s + "\ufe0f\u20e3")
		emojiNum += emojiDigit
	}

	return emojiNum
}
