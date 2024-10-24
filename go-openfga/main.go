package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/openfga/go-sdk/client"
	. "github.com/openfga/go-sdk/client"
)

func createRelationshipTuple(fgaClient *client.OpenFgaClient, modelId string) {
	options := ClientWriteOptions{
		AuthorizationModelId: &modelId,
	}
	body := ClientWriteRequest{
		Writes: []ClientTupleKey{
			{
				User:     "user:anne",
				Relation: "reader",
				Object:   "document:Z",
			},
		},
	}

	data, err := fgaClient.Write(context.Background()).
		Body(body).
		Options(options).
		Execute()
	if err != nil {
		// .. Handle error
	}
	_ = data // use the response
}

func deleteRelationshipTuple(fgaClient *client.OpenFgaClient, modelId string) {
	options := ClientWriteOptions{
		AuthorizationModelId: &modelId,
	}

	body := ClientWriteRequest{
		Deletes: []ClientTupleKeyWithoutCondition{
			{
				User:     "user:anne",
				Relation: "reader",
				Object:   "document:Z",
			},
		},
	}

	data, err := fgaClient.Write(context.Background()).
		Body(body).
		Options(options).
		Execute()

	if err != nil {
		// .. Handle error
	}

	_ = data // use the response
}

func checkAPI(fgaClient *client.OpenFgaClient, modelId string) {
	options := ClientCheckOptions{
		AuthorizationModelId: &modelId,
	}

	body := ClientCheckRequest{
		User:     "user:anne",
		Relation: "reader",
		Object:   "document:Z",
	}

	data, err := fgaClient.Check(context.Background()).
		Body(body).
		Options(options).
		Execute()

	if err != nil {
		fmt.Println(err)
	}

	obj := *data
	objJSON, _ := json.Marshal(obj)
	println(string(objJSON))

	// println(*(data.Allowed))
}

func listObjectsAPI(fgaClient *client.OpenFgaClient, modelId string) {
	options := ClientListObjectsOptions{
		AuthorizationModelId: &modelId,
	}

	body := ClientListObjectsRequest{
		User:     "user:anne",
		Relation: "reader",
		Type:     "document",
	}

	data, err := fgaClient.ListObjects(context.Background()).
		Body(body).
		Options(options).
		Execute()

	if err != nil {
		fmt.Println("Error encountered")
	}

	obj := *data
	objJSON, _ := json.Marshal(obj)
	println(string(objJSON))
}

func main() {
	// router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "pong"})
	// })
	// router.Run(":8888")

	modelId := "01HW7V5Q20SY6MJ22CNMKN7NCF"

	fgaClient, err := NewSdkClient(&ClientConfiguration{
		ApiUrl:               "http://localhost:8082",
		StoreId:              "01HW7TZ1DQH47JFZWPHMB2KKCJ",
		AuthorizationModelId: "01HW7V5Q20SY6MJ22CNMKN7NCF",
	})

	if err != nil {
		fmt.Println("Error during client configuration")
	}

	// Step - 1 : Create store
	// resp, err := fgaClient.CreateStore(context.Background()).Body(ClientCreateStoreRequest{Name: "FGA Demo"}).Execute()
	// if err != nil {
	// 	 // .. Handle error
	// }

	// Step - 2 : Model configuration
	// var writeAuthorizationModelRequestString = "{\"schema_version\":\"1.1\",\"type_definitions\":[{\"type\":\"user\"},{\"type\":\"document\",\"relations\":{\"reader\":{\"this\":{}},\"writer\":{\"this\":{}},\"owner\":{\"this\":{}}},\"metadata\":{\"relations\":{\"reader\":{\"directly_related_user_types\":[{\"type\":\"user\"}]},\"writer\":{\"directly_related_user_types\":[{\"type\":\"user\"}]},\"owner\":{\"directly_related_user_types\":[{\"type\":\"user\"}]}}}}]}"
	// var body WriteAuthorizationModelRequest
	// if err := json.Unmarshal([]byte(writeAuthorizationModelRequestString), &body); err != nil {
	// 	// .. Handle error
	// 	return
	// }

	// data, err := fgaClient.WriteAuthorizationModel(context.Background()).
	// 	Body(body).
	// 	Execute()
	// if err != nil {
	// 	fmt.Println("Error during model configuration")
	// }
	// // data.AuthorizationModelId = "01HVMMBCMGZNT3SED4Z17ECXCA"

	// Step - 3 : Create Relationship Tuple
	// createRelationshipTuple(fgaClient, modelId)

	// Delete Relationship Tuple (Optional)
	// deleteRelationshipTuple(fgaClient, modelId)

	// Perform a check API
	checkAPI(fgaClient, modelId)

	// Perform a List Objects API
	listObjectsAPI(fgaClient, modelId)

}
