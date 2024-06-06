# week21




## Objective: 

***Your task in this assessment is to create a Golang gRPC service that provides specific functionalities for managing user details and includes a search capability.***

### Primary objectives:

- [ ] Simulate a database by maintaining a list of user details within a variable.
- [ ] Create gRPC endpoints for fetching user details based on a user ID and retrieving a list of user details based on a list of user IDs.
- [ ] Implement a search functionality to find user details based on specific criteria.


## Sample User Model:

```json
{
	"id": 1,
	"fname": "Steve",
	"city": "LA",
	"phone": 1234567890,
	"height": 5.8,
	"Married": true
}
```



## Maintain Code Quality and Design: 

- Ensure that the code is well-structured and follows best practices.
- Apply suitable design patterns to enhance the maintainability and extensibility of your service.


## Develop and Test APIs:

- Implement the specified gRPC service methods to accomplish the tasks.
- Write comprehensive unit tests to verify the correctness of your APIs.


## Input Validation and Response Handling:

- Validate incoming requests to ensure they adhere to the defined requirements.
- Handle requests appropriately and respond with meaningful messages, especially in the case of errors.


## Implement Search Functionality:

- Create a search endpoint that allows users to search for specific user details based on criteria (e.g., city, phone number, marital status, etc.).


## Cover Edge Cases:

- Identify potential edge cases and consider these in your implementation to provide robust and reliable functionality.


### Extra Brownie Points:

- Use design patterns creatively to improve your service's architecture and efficiency.
- Containerize the entire application using Docker for easy deployment and scaling.

