package gutils

import (
	"encoding/json"
	"log"

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
