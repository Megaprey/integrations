package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// === Скрипт: JSON → XML для БКИ ===
func Process_93d9fa01(msg Message, sender MessageSender) error {
	log.Printf("🔄 JSON → XML (BKI): обработка %s", msg.ID)

	payload, ok := msg.Data["payload"].(map[string]interface{})
	if !ok {
		payload = msg.Data
	}

	applicationId := ""
	if env, ok := msg.Data["envelope"].(map[string]interface{}); ok {
		if v, ok := env["applicationId"].(string); ok {
			applicationId = v
		}
	}

	client, _ := payload["client"].(map[string]interface{})
	passport, _ := client["passport"].(map[string]interface{})

	bkiDocRef := ""
	if consents, ok := payload["consents"].([]interface{}); ok {
		for _, c := range consents {
			cm, _ := c.(map[string]interface{})
			if cm["type"] == "BKI_REQUEST" {
				bkiDocRef, _ = cm["documentRef"].(string)
			}
		}
	}

	xmlBody := fmt.Sprintf(`<BkiReportRequest>
  <applicationId>%s</applicationId>
  <client>
    <firstName>%s</firstName>
    <lastName>%s</lastName>
    <birthDate>%s</birthDate>
    <passportSeries>%s</passportSeries>
    <passportNumber>%s</passportNumber>
    <inn>%s</inn>
    <snils>%s</snils>
  </client>
  <requestedFrom>NBKI</requestedFrom>
  <requestedFrom>OKB</requestedFrom>
  <consent>
    <granted>true</granted>
    <documentRef>%s</documentRef>
  </consent>
</BkiReportRequest>`,
		escapeXml(applicationId),
		escapeXml(strVal(client["firstName"])),
		escapeXml(strVal(client["lastName"])),
		escapeXml(strVal(client["birthDate"])),
		escapeXml(strVal(passport["series"])),
		escapeXml(strVal(passport["number"])),
		escapeXml(strVal(client["inn"])),
		escapeXml(strVal(client["snils"])),
		escapeXml(bkiDocRef),
	)

	outMsg := Message{
		ID:            msg.ID,
		Body:          []byte(xmlBody),
		IsJSON:        false,
		Headers:       msg.Headers,
		AggregateName: msg.AggregateName,
		AggregateID:   msg.AggregateID,
		CorrelationID: msg.CorrelationID,
		ReplyTo:       msg.ReplyTo,
	}

	log.Printf("✅ JSON → XML готов (BKI), %d байт", len(xmlBody))

	if sender != nil {
		if err := sender.Send(outMsg); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
			return err
		}
	}
	return nil
}

func strVal(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	if n, ok := v.(float64); ok {
		if n == float64(int64(n)) {
			return fmt.Sprintf("%d", int64(n))
		}
		return fmt.Sprintf("%v", n)
	}
	b, _ := json.Marshal(v)
	return string(b)
}

func escapeXml(s string) string {
	s = replaceAll(s, "&", "&amp;")
	s = replaceAll(s, "<", "&lt;")
	s = replaceAll(s, ">", "&gt;")
	s = replaceAll(s, `"`, "&quot;")
	s = replaceAll(s, "'", "&apos;")
	return s
}

func replaceAll(s, old, new string) string {
	for {
		i := indexOf(s, old)
		if i < 0 {
			return s
		}
		s = s[:i] + new + s[i+len(old):]
	}
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
