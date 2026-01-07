# API DOC

# Users

#### /create_user

- Method: POST
- Description: Create a new user.
- Request Body:
  - email: string (required)
  - password: string (required)
- Response:
  - 201 Created: User created successfully.
  - 400 Bad Request: Invalid input data.
  - 500 Internal Server Error: Server error.
