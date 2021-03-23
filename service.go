package gutils

import (
	"encoding/json"
	"log"
	"regexp"

	"github.com/streadway/amqp"
	"go.elastic.co/apm"
)

// >> Service communication example:
//  serviceResMap, ok := gutils.ServiceCommunicator(userResMap, "service", "service.queue", data, rabbitmq.ClientRPC)
// 	if !ok {
// 		return gutils.ReturnServiceError()
// 	}

// ServiceCommunicator : Service communicator function for rabbitmq
func ServiceCommunicator(dataMap map[string]interface{}, serviceName string, topicName string, data amqp.Delivery, clientRPC func([]byte, string, string, string) []byte) (map[string]interface{}, bool) {
	if dataMap["user_id"] == nil {
		dataMap["user_id"] = "service"
	}

	reqByte, marshalErr := json.Marshal(dataMap)
	if marshalErr != nil {
		log.Println("ServiceCommunicator Marshal error: ", marshalErr)
		return nil, false
	}

	resByte := clientRPC(reqByte, serviceName, topicName, data.CorrelationId)

	resMap := map[string]interface{}{}

	unmarshalErr := json.Unmarshal(resByte, &resMap)
	if unmarshalErr != nil {
		log.Println("ServiceCommunicator Unmarshal error: ", marshalErr)
		return nil, false
	}

	return resMap, true
}

// ServiceCommunicatorV2 : Service communicator with tracing function for rabbitmq
func ServiceCommunicatorV2(dataMap map[string]interface{}, serviceName string, topicName string, data amqp.Delivery, clientRPC func([]byte, string, string, amqp.Delivery) []byte) (map[string]interface{}, bool) {
	if dataMap["user_id"] == nil {
		dataMap["user_id"] = "service"
	}

	reqByte, marshalErr := json.Marshal(dataMap)
	if marshalErr != nil {
		log.Println("ServiceCommunicator Marshal error: ", marshalErr)
		return nil, false
	}

	resByte := clientRPC(reqByte, serviceName, topicName, data)

	resMap := map[string]interface{}{}

	unmarshalErr := json.Unmarshal(resByte, &resMap)
	if unmarshalErr != nil {
		log.Println("ServiceCommunicator Unmarshal error: ", marshalErr)
		return nil, false
	}

	return resMap, true
}

// ServiceCommunicatorAPM : Service communicator function for rabbitmq
func ServiceCommunicatorAPM(txApm *apm.Transaction, dataMap map[string]interface{}, serviceName string, topicName string, data amqp.Delivery, clientRPC func(*apm.Transaction, []byte, string, string, string) []byte) (map[string]interface{}, bool) {

	if dataMap["user_id"] == nil {
		dataMap["user_id"] = "service"
	}

	reqByte, marshalErr := json.Marshal(dataMap)
	if marshalErr != nil {
		log.Println("ServiceCommunicator Marshal error: ", marshalErr)
		return nil, false
	}

	resByte := clientRPC(txApm, reqByte, serviceName, topicName, data.CorrelationId)

	resMap := map[string]interface{}{}

	unmarshalErr := json.Unmarshal(resByte, &resMap)
	if unmarshalErr != nil {
		log.Println("ServiceCommunicator Unmarshal error: ", marshalErr)
		return nil, false
	}

	return resMap, true
}

// ReturnServiceError : Return error response when somethong wrong about service communication
func ReturnServiceError() []byte {
	errRes := errorResponse{}
	errRes.Success = false
	errRes.Message = "Service Communication error."

	resByte, _ := json.Marshal(errRes)

	return resByte
}

func GetFGServiceName(exchangeName string) string {
	if match, _ := regexp.MatchString("(fblfg_order)", exchangeName); match {
		return "fblfg-order-sev"
	} else if match, _ := regexp.MatchString("(order|address|zipcode|sub-district|district|province)", exchangeName); match {
		return "order-service"
	} else if match, _ := regexp.MatchString("(content)", exchangeName); match {
		return "content-service"
	} else if match, _ := regexp.MatchString("(courier_partner)", exchangeName); match {
		return "courier-partner"
	} else if match, _ := regexp.MatchString("(courier)", exchangeName); match {
		return "courier-service"
	} else if match, _ := regexp.MatchString("(activity_logging)", exchangeName); match {
		return "act-log-sev"
	} else if match, _ := regexp.MatchString("(customer)", exchangeName); match {
		return "cus-service"
	} else if match, _ := regexp.MatchString("(dashboard)", exchangeName); match {
		return "dashbrd-service"
	} else if match, _ := regexp.MatchString("(employee_authentication)", exchangeName); match {
		return "emp-auth-sev"
	} else if match, _ := regexp.MatchString("(employee)", exchangeName); match {
		return "emp-service"
	} else if match, _ := regexp.MatchString("(invoice)", exchangeName); match {
		return "invoice-service"
	} else if match, _ := regexp.MatchString("(mail)", exchangeName); match {
		return "mail-service"
	} else if match, _ := regexp.MatchString("(payment)", exchangeName); match {
		return "payment-service"
	} else if match, _ := regexp.MatchString("(product|warehouse)", exchangeName); match {
		return "product-service"
	} else if match, _ := regexp.MatchString("(sale)", exchangeName); match {
		return "sale-service"
	} else if match, _ := regexp.MatchString("(shop)", exchangeName); match {
		return "shop-service"
	} else if match, _ := regexp.MatchString("(social)", exchangeName); match {
		return "social-service"
	} else if match, _ := regexp.MatchString("(support)", exchangeName); match {
		return "support-service"
	} else if match, _ := regexp.MatchString("(user)", exchangeName); match {
		return "user-service"
	}

	return "none"
}
