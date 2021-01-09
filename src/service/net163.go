package service

import (
	"encoding/json"
	"finance-spider/src/model"
	"finance-spider/src/util"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

const API = "http://quotes.money.163.com/fn/service/netvalue.php?host=/fn/service/netvalue.php&query=STYPE:FDO;TYPE3:GPX&fields=no,PUBLISHDATE,SYMBOL,SNAME,NAV,PCHG,M12RETRUN,SLNAVG,ZJZC&sort=no&order=asc&count=20&type=query&req=6048"

func FundDownload() error {
	//count := (a["count"]).(int)
	//total := (a["total"]).(int)
	currentPage := 0
	for {
		body, _ := util.Download(API + "&page=" + strconv.Itoa(currentPage))
		//println(body)

		var a map[string]interface{}
		err := json.Unmarshal([]byte(body), &a)
		if err != nil {
			fmt.Printf("unmarshal err=%v\n", err)
			return err
		}
		funds := (a["list"]).([]interface{})
		for _, m := range funds {
			var fund = &model.Fund{}
			_ = mapstructure.Decode(m.(map[string]interface{}), fund)
			if err := db.Create(fund).Error; nil != err {
				logger.Error("error create record", fund, err)
				// fmt.Printf("error create record: %v, err: %v\n", fund, err)
				continue
			}
		}
		page := int((a["pagecount"]).(float64))
		if currentPage >= page-1 {
			break
		}
		currentPage++
	}
	return nil
}
