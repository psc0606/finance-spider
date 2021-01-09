package model

// 股票基金
// no,PUBLISHDATE,SYMBOL,SNAME,NAV,PCHG,M12RETRUN,SLNAVG,ZJZC
type Fund struct {
	Model
	// 序号
	// No int `gorm:"size:255;column:no" json:"NO" structs:"NO"`
	// 净值日期
	PublishDate string `gorm:"type:date;column:publish_date" json:"PUBLISHDATE" structs:"PUBLISHDATE"`
	// 基金代码
	Symbol string `gorm:"type:varchar(16);column:symbol" json:"SYMBOL" structs:"SYMBOL"`
	// 基金名称
	Sname string `gorm:"type:varchar(32);column:sname" json:"SNAME" structs:"SNAME"`
	// 单位净值
	Nav float64 `gorm:"column:nav" json:"NAV" structs:"NAV"`
	// 增长率(%)
	Pchg float64 `gorm:"column:pchg" json:"PCHG" structs:"PCHG"`
	// 一年回报(%)
	M12retrun float64 `gorm:"column:m12_return" json:"M12RETRUN" structs:"M12RETRUN"`
	// 上市至今回报(%)
	Slnavg float64 `gorm:"column:slnavg" json:"SLNAVG" structs:"SLNAVG"`
	// 规模(单位RMB)
	Zjzc int `gorm:"type:bigint;column:zjzc" json:"ZJZC" structs:"ZJZC"`
}
