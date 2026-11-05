package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type InnerBkiSource_e5f6a7b5 struct {
    Bureau         string `mapstructure:"bureau"         json:"bureau_inner441"`
    Score          int    `mapstructure:"score"          json:"score_inner442"`
    ActiveLoans    int    `mapstructure:"activeLoans"    json:"activeLoans_inner443"`
    TotalDebt      int    `mapstructure:"totalDebt"      json:"totalDebt_inner444"`
    OverdueMaxDays int    `mapstructure:"overdueMaxDays" json:"overdueMaxDays_inner445"`
}
type InnerBkiSummary_e5f6a7b5 struct {
    WorstStatus           string `mapstructure:"worstStatus"           json:"worstStatus_inner451"`
    HasBankruptcy         bool   `mapstructure:"hasBankruptcy"         json:"hasBankruptcy_inner452"`
    HasCurrentDelinquency bool   `mapstructure:"hasCurrentDelinquency" json:"hasCurrentDelinquency_inner453"`
}
type InnerMessage_e5f6a7b5 struct {
    RequestID string                  `mapstructure:"requestId" json:"requestId_inner461"`
    Status    string                  `mapstructure:"status"    json:"status_inner462"`
    Sources   []InnerBkiSource_e5f6a7b5 `mapstructure:"sources"   json:"sources_inner463"`
    Summary   InnerBkiSummary_e5f6a7b5  `mapstructure:"summary"   json:"summary_inner464"`
}

func Process_f6a71a03(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (out→inner, ID e5f6a7b5)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var inner InnerMessage_e5f6a7b5
    if err := mapstructure.Decode(raw, &inner); err != nil { log.Printf("❌ decode: %v", err); sendOrWarn(msg, output); return msg }
    newBody, err := json.Marshal(inner)
    if err != nil { log.Printf("❌ marshal: %v", err); sendOrWarn(msg, output); return msg }
    msg.Body = newBody; msg.IsJSON = true; msg.Data = inner
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