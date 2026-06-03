package handlers
//Instead of writing this response 
//code again and again in every handler, you make one reusable function.
import (
     "encoding/json" //convert GO to JSON
	 "github.com/aws/aws-lambda-go/events"
)
                
func apiResponse()(status int , body interface{})(*events.APIGatewayProxyResponse , error){
      resp := events.APIGatewayProxyResponse{Headers : map[string]string["Content-Type" : "application/json"]}
	  resp.StatusCode = status
	  stringBody, _ := json.Marshal(body)
	  resp.Body = string(stringBody)
	  return &resp, nil
}

//status -> HTTP status code //body -> data that we send back
//here interface means the body can be string , map , struct , array

//FULL FLOW

/*Frontend / Client
   ↓
API Gateway
   ↓
Lambda handler
   ↓
Your Go handler function
   ↓
ApiResponse(...)
   ↓
API Gateway sends response back
   ↓
Frontend receives it
*/