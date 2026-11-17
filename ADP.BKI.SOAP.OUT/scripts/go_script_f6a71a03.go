package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func Process_f6a71a03(msg Message, sender MessageSender) error {
	log.Printf("🔄 XML → JSON (BKI): обработка %s", msg.ID)

	result, err := parseSoapBody(msg.Body, "ReportResponse")
	if err != nil {
		log.Printf("❌ Ошибка парсинга XML: %v", err)
		return err
	}

	// ⚠️ Разворачиваем вложенный <ReportResponse> (node-soap оборачивает)
	if inner, ok := result["ReportResponse"].(map[string]interface{}); ok {
		log.Printf("🔍 Разворачиваю вложенный ReportResponse")
		result = inner
	}

	finalJSON := map[string]interface{}{
		"requestId_inner461": strValF6(result["requestId"]),
		"status_inner462":    strValF6(result["status"]),
	}

	// sources
	if sources, ok := result["sources"].(map[string]interface{}); ok {
		if source, ok := sources["source"].(map[string]interface{}); ok {
			finalJSON["sources_inner463"] = []interface{}{buildSource(source)}
		} else if sourceArr, ok := sources["source"].([]interface{}); ok {
			finalSources := []interface{}{}
			for _, s := range sourceArr {
				sm, _ := s.(map[string]interface{})
				finalSources = append(finalSources, buildSource(sm))
			}
			finalJSON["sources_inner463"] = finalSources
		}
	} else if sourcesArr, ok := result["sources"].([]interface{}); ok {
		finalSources := []interface{}{}
		for _, s := range sourcesArr {
			sm, _ := s.(map[string]interface{})
			finalSources = append(finalSources, buildSource(sm))
		}
		finalJSON["sources_inner463"] = finalSources
	}

	if summary, ok := result["summary"].(map[string]interface{}); ok {
		finalJSON["summary_inner464"] = map[string]interface{}{
			"worstStatus_inner451":           strValF6(summary["worstStatus"]),
			"hasBankruptcy_inner452":         summary["hasBankruptcy"],
			"hasCurrentDelinquency_inner453": summary["hasCurrentDelinquency"],
		}
	}

	jsonBytes, err := json.Marshal(finalJSON)
	if err != nil {
		log.Printf("❌ Ошибка сериализации JSON: %v", err)
		return err
	}

	outMsg := Message{
		ID:            msg.ID,
		Body:          jsonBytes,
		Data:          finalJSON,
		IsJSON:        true,
		Headers:       msg.Headers,
		AggregateName: msg.AggregateName,
		AggregateID:   msg.AggregateID,
		CorrelationID: msg.CorrelationID,
		ReplyTo:       msg.ReplyTo,
	}

	log.Printf("✅ XML → JSON готов (BKI): %s", string(jsonBytes))

	if sender != nil {
		if err := sender.Send(outMsg); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
			return err
		}
	}
	return nil
}

func buildSource(sm map[string]interface{}) map[string]interface{} {
	if sm == nil {
		sm = map[string]interface{}{}
	}
	return map[string]interface{}{
		"bureau_inner441":         strValF6(sm["bureau"]),
		"score_inner442":          toNumber(sm["score"]),
		"activeLoans_inner443":    toNumber(sm["activeLoans"]),
		"totalDebt_inner444":      toNumber(sm["totalDebt"]),
		"overdueMaxDays_inner445": toNumber(sm["overdueMaxDays"]),
	}
}

func parseSoapBody(data []byte, targetElement string) (map[string]interface{}, error) {
	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	inBody := false

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == "Body" {
				inBody = true
			} else if inBody && t.Name.Local == targetElement {
				val, err := decodeElement(decoder, &t)
				if err != nil {
					return nil, err
				}
				if m, ok := val.(map[string]interface{}); ok {
					return m, nil
				}
				return map[string]interface{}{targetElement: val}, nil
			}
		}
	}
	return nil, fmt.Errorf("элемент <%s> не найден", targetElement)
}

func decodeElement(decoder *xml.Decoder, start *xml.StartElement) (interface{}, error) {
	var text strings.Builder
	children := make(map[string]interface{})
	hasChildren := false

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			hasChildren = true
			childVal, err := decodeElement(decoder, &t)
			if err != nil {
				return nil, err
			}
			addToMap(children, t.Name.Local, childVal)
		case xml.CharData:
			text.Write(t)
		case xml.EndElement:
			if t.Name.Local == start.Name.Local {
				if hasChildren {
					if shouldFlatten(start.Name.Local, children) {
						for _, v := range children {
							if arr, ok := v.([]interface{}); ok {
								return arr, nil
							}
						}
					}
					return children, nil
				}
				s := strings.TrimSpace(text.String())
				if s == "" {
					return nil, nil
				}
				if n, err := strconv.ParseFloat(s, 64); err == nil {
					return n, nil
				}
				if b, err := strconv.ParseBool(s); err == nil {
					return b, nil
				}
				return s, nil
			}
		}
	}
}

func shouldFlatten(parentName string, children map[string]interface{}) bool {
	flattenParents := map[string]bool{
		"sources": true, "documents": true, "consents": true,
		"items": true, "checks": true, "flags": true,
		"reasonCodes": true, "proceedings": true,
	}
	if !flattenParents[parentName] {
		return false
	}
	if len(children) != 1 {
		return false
	}
	for _, v := range children {
		if _, ok := v.([]interface{}); ok {
			return true
		}
	}
	return false
}

func addToMap(m map[string]interface{}, key string, value interface{}) {
	if existing, ok := m[key]; ok {
		if arr, ok := existing.([]interface{}); ok {
			m[key] = append(arr, value)
		} else {
			m[key] = []interface{}{existing, value}
		}
	} else {
		m[key] = value
	}
}

func toNumber(v interface{}) interface{} {
	if s, ok := v.(string); ok {
		if n, err := strconv.ParseFloat(s, 64); err == nil {
			return n
		}
	}
	return v
}

func strValF6(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}
