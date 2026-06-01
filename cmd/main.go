package main

import(
	"go-serverless/pkg/handlers"
	"os"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func main(){
	region := os.Getenv("AWS_REGION")

	//creating an AWS session
	awsSession , err := session.NewSession(&aws.Config{
		Region: aws.String(region) //sets region
	},)
    //if session fails to create the program stops
	if err != nil{
		return
	}
	//create DynamoDB client using this AWS session
	//client can be used to perform operations like putItem,GetItem,UpdateItem etc

	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler) //handler is the entry function for every Lambda request/event


}

//why not put handler inside main ?
//main runs only once while creating the client after that all req go to handler 
//otherwise we will end up having multiple clients coz every time a request comes it first creates a client then goes to the handler




const tableName = "LambdaInGoUser"
            //input                             //output
func handler(req events.APIGatewayProxyRequest)(*events.APIGatewayProxyResponse , error){
	switch req.HTTPMethod{
	case "GET":
		return handlers.GetUser(req , tableName , dynaClient)
	case "POST":
		return handlers.CreateUser(req , tableName , dynaClient)
	case "PUT":
		return handlers.UpdateUser(req , tableName , dynaClient)
	case "DELETE":
		return handlers.DeleteUser(req , tableName , dynaClient)
	
    default:
		return handlers.UnhandledMethod()
	}


}