package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func Process_93d9fa01(msg Message, sender MessageSender) error {
	log.Printf("🔄 JSON → XML (BKI): обработка %s", msg.ID)

	requestId := strVal(msg.Data["requestId_inner431"])
	applicationId := strVal(msg.Data["applicationId_inner432"])

	client, _ := msg.Data["client_inner433"].(map[string]interface{})
	firstName := strVal(client["firstName_inner401"])
	lastName := strVal(client["lastName_inner402"])
	middleName := strVal(client["middleName_inner403"])
	birthDate := strVal(client["birthDate_inner404"])
	passportSeries := strVal(client["passportSeries_inner405"])
	passportNumber := strVal(client["passportNumber_inner406"])
	inn := strVal(client["inn_inner407"])
	snils := strVal(client["snils_inner408"])

	requestedFrom := ""
	if arr, ok := msg.Data["requestedFrom_inner434"].([]interface{}); ok {
		for _, v := range arr {
			requestedFrom += fmt.Sprintf("<requestedFrom>%s</requestedFrom>", escapeXml(strVal(v)))
		}
	}

	requestPurpose := strVal(msg.Data["requestPurpose_inner435"])

	consent, _ := msg.Data["consent_inner436"].(map[string]interface{})
	consentGranted := strVal(consent["granted_inner421"])
	consentDocRef := strVal(consent["documentRef_inner422"])

	// ⚠️ Без namespace — XSD тоже должен быть без targetNamespace
	xmlBody := fmt.Sprintf(`<ReportRequest>
  <requestId>%s</requestId>
  <applicationId>%s</applicationId>
  <client>
    <firstName>%s</firstName>
    <lastName>%s</lastName>
    <middleName>%s</middleName>
    <birthDate>%s</birthDate>
    <passportSeries>%s</passportSeries>
    <passportNumber>%s</passportNumber>
    <inn>%s</inn>
    <snils>%s</snils>
  </client>
  %s
  <requestPurpose>%s</requestPurpose>
  <consent>
    <granted>%s</granted>
    <documentRef>%s</documentRef>
  </consent>
</ReportRequest>`,
		escapeXml(requestId),
		escapeXml(applicationId),
		escapeXml(firstName),
		escapeXml(lastName),
		escapeXml(middleName),
		escapeXml(birthDate),
		escapeXml(passportSeries),
		escapeXml(passportNumber),
		escapeXml(inn),
		escapeXml(snils),
		requestedFrom,
		escapeXml(requestPurpose),
		escapeXml(consentGranted),
		escapeXml(consentDocRef),
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
	if b, ok := v.(bool); ok {
		return fmt.Sprintf("%v", b)
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
