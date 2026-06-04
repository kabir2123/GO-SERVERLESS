package user
import(

	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

)

var(

)

type User struct(
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`

)

var(
	ErrorFailedToFetchRecord = "Failed to fetch Record"
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorInvalidUserData = "invalid user data"
	ErrorCouldNotMarshalItem = "could not marshal item"
	ErrorCouldNotDeleteItem = "could not delete item"
	ErrorCouldNnotDynamoPutItem = "could not dynamo put item"
	ErrorUserDoesNotExist = "user doesnot exist"
)

func FetchUser(email , tableName string , dynaClient dynamodbiface.DynamoDBAPI)(*User , error){
	//creating the request object that will be sent to DynamoDB(dynamodb.GetItemInput)
	input := &dynamodb.GetItemInput{
		//setting up primary key which is email here
		Key: map[string]*dynamodb.AttributeValue{
			//main query
			"email":{
				S: aws.String(email)
			}
		},
		TableName: aws.String(tableName)

	}
	result , err := dynaClient.GetItem(input) //sends request to DynamoDB
	if err!= nil{
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	//Create empty User object

	item := new(User)
	/*{
    Email: "",
    FirstName: "",
    LastName: "",}*/
	err = dynamodbattribute.UnmarshalMap(result.Item ,  item) //convert to json for go 
	if err != nil{
		return nil , errors.New(ErrorFailedToFetchRecord)
	}
	return item, nil
}


func FetchUsers(tableName string , dynaClient dynamodbiface.DynamoDBAPI)(*[]User , error){
	input := &dynamodb.ScanInput{ //scan reads every item in the table
		TableName: aws.String(tableName)
	}
	result , err := dynaClient.Scan(input)
	if err != nil {
		return nil , errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]User)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items ,  item)
	if err != nil {
		return nil , errors.New(ErrorFailedToFetchRecord)
	}
	return item , nil
}

func CreateUser(req events.APIGatewayProxyRequest , tableName string , dynaClient dynamodbiface.DynamoDBAPI)(*User, error){
    var u User
	json.Unmarshal([]byte(req.body) , &u); err!= nil{
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(u.Email){
		return nil , errros.New(ErrorInvalidEmail)
	}

	currentUser, _ := FetchUser(u.Email , tableName , dynaClient)
	if currentUser != nil && len(currentUser.Email) != 0{
		return nil , erros.New(ErrorUserAlreadyExists)
	}

	av, err := dynamodbattribute.marshalMap(u)

	if err != nil{
		return nil , errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(tableName)


	}
	_, err = dynaClient.PutItem(input)
	if err != nil{
		return nil , errors.New(ErrorCouldNnotDynamoPutItem)
	}

	return &u , nil

}
