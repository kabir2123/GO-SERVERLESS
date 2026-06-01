package main

import(
	"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
	lambda.Start(handler)
}