package gutils

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

// >> Service communication example:
//  serviceResMap, ok := utils.ServiceCommunicator(userResMap, "service", "service.queue", data)
// 	if !ok {
// 		return utils.ReturnServiceError()
// 	}

func ServiceCommunicator(dataMap map[string]interface{}, serviceName string, topicName string, data amqp.Delivery, clientRPC func([]byte, string, string, string) []byte) (map[string]interface{}, bool) {
	dataMap["user_id"] = "service"

	reqByte, marshalErr := json.Marshal(dataMap)
	if marshalErr != nil {
		fmt.Println("Marshal error: ", marshalErr)

		return nil, false
	}

	resByte := clientRPC(reqByte, serviceName, topicName, data.CorrelationId)

	resMap := map[string]interface{}{}

	unmarshalErr := json.Unmarshal(resByte, &resMap)
	if unmarshalErr != nil {
		fmt.Println("Unmarshal error: ", marshalErr)

		return nil, false
	}

	return resMap, true
}

func ReturnServiceError() []byte {
	errRes := ErrorResponse{}
	errRes.Success = false
	errRes.Message = "Service Communication error."

	resByte, _ := json.Marshal(errRes)

	return resByte
}
