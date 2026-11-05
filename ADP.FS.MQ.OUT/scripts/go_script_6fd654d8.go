package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type FsBkiInputOut_b8c9d0e8 struct {
    Score          int `mapstructure:"score_inner601"          json:"score"`
    ActiveLoans    int `mapstructure:"activeLoans_inner602"    json:"activeLoans"`
    TotalDebt      int `mapstructure:"totalDebt_inner603"      json:"totalDebt"`
    OverdueMaxDays int `mapstructure:"overdueMaxDays_inner604" json:"overdueMaxDays"`
}
type FsAntifraudInputOut_b8c9d0e8 struct {
    RiskScore int    `mapstructure:"riskScore_inner611" json:"riskScore"`
    RiskLevel string `mapstructure:"riskLevel_inner612" json:"riskLevel"`
}
type FsInternalInputOut_b8c9d0e8 struct {
    ClientSince       string `mapstructure:"clientSince_inner621"       json:"clientSince"`
    AvgMonthlyBalance int    `mapstructure:"avgMonthlyBalance_inner622" json:"avgMonthlyBalance"`
    SalaryClient      bool   `mapstructure:"salaryClient_inner623"      json:"salaryClient"`
}
type FsProductInputOut_b8c9d0e8 struct {
    Amount     int    `mapstructure:"amount_inner631"     json:"amount"`
    TermMonths int    `mapstructure:"termMonths_inner632" json:"termMonths"`
    Type       string `mapstructure:"type_inner633"       json:"type"`
}
type FsInputsOut_b8c9d0e8 struct {
    Kyc       string                      `mapstructure:"kyc_inner641"       json:"kyc"`
    Bki       FsBkiInputOut_b8c9d0e8       `mapstructure:"bki_inner642"       json:"bki"`
    Antifraud FsAntifraudInputOut_b8c9d0e8 `mapstructure:"antifraud_inner643" json:"antifraud"`
    Internal  FsInternalInputOut_b8c9d0e8  `mapstructure:"internal_inner644"  json:"internal"`
    Product   FsProductInputOut_b8c9d0e8   `mapstructure:"product_inner645"   json:"product"`
}
type OutMessage_b8c9d0e8 struct {
    RequestID     string                `mapstructure:"requestId_inner651"     json:"requestId"`
    ApplicationID string                `mapstructure:"applicationId_inner652" json:"applicationId"`
    ClientID      string                `mapstructure:"clientId_inner653"      json:"clientId"`
    FeatureGroups []string              `mapstructure:"featureGroups_inner654" json:"featureGroups"`
    Inputs        FsInputsOut_b8c9d0e8  `mapstructure:"inputs_inner655"        json:"inputs"`
}

func Process_6fd654d8(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (inner→out, ID b8c9d0e8)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v; delete(msg.Data, key) } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var out OutMessage_b8c9d0e8
    if err := mapstructure.Decode(raw, &out); err != nil { log.Printf("❌ decode: %v", err); sendOrWarn(msg, output); return msg }
    newBody, err := json.Marshal(out)
    if err != nil { log.Printf("❌ marshal: %v", err); sendOrWarn(msg, output); return msg }
    msg.Body = newBody; msg.IsJSON = true; msg.Data = out
    log.Printf("   • Body: %s", string(msg.Body))
    sendOrWarn(msg, output); return msg
}

func sendOrWarn(msg Message, output MessageSender) {
	if output != nil {
		if err := output.Send(msg); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
		}
	} else {
		log.Printf("⚠️ output sender не установлен")
	}
}