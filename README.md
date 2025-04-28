# Serverless Go React Native Project

This project is a serverless application built using Go for the backend and React Native for the frontend. 
It utilizes AWS Lambda and DynamoDB to manage tasks in a serverless environment.

## Project Structure

```
serverless-go-react-native
├── backend
│   ├── cmd
│   │   └── lambda
│   │       └── main.go          # Entry point for the Lambda function
│   ├── internal
│   │   ├── db
│   │   │   └── dynamodb.go      # Functions for interacting with DynamoDB
│   │   ├── handlers
│   │   │   └── tasks.go         # Handler for task-related requests
│   │   └── models
│   │       └── task.go          # Task model definition
│   ├── go.mod                    # Go module definition
│   ├── go.sum                    # Module dependency checksums
│   └── template.yaml             # AWS SAM template for serverless application
├── frontend
│   ├── src
│   │   ├── components
│   │   │   └── TaskItem.tsx      # React component for a single task item
│   │   ├── screens
│   │   │   └── TaskListScreen.tsx # Component to display the list of tasks
│   │   ├── services
│   │   │   └── api.ts            # API call functions to the backend
│   │   └── App.tsx               # Main entry point for the React Native app
│   ├── package.json              # npm configuration for frontend
│   └── app.json                  # React Native app configuration
└── README.md                     # Project documentation
```

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- Node.js (version 14 or later)
- AWS CLI configured with your credentials

### Backend Setup

1. Navigate to the `backend` directory:
```
cd backend
```

2. Install Go dependencies:
```
go mod tidy
```

3. Install Serverless framework globally:
```
npm install -g serverless
```

3. Deploy the backend using AWS SAM:
```
sls deploy
```

### Frontend Setup

1. Navigate to the `frontend` directory:
```
cd frontend
```

2. Install npm dependencies:
```
npm install
```

3. Run the React Native application:
```
npm start
```

## Usage

- The backend provides an API for managing tasks. You can create, read, update, and delete tasks through the defined endpoints.
- The frontend application allows users to interact with the task management system, displaying tasks and enabling task operations.
