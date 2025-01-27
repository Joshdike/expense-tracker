
The Expense Tracker API is a RESTful service designed to help users manage their expenses.
It allows users to create, read, update, and delete expense records.
Each expense is stored in a database and includes details such as ID, date, amount, category, description, payment method, and timestamps.

Features:

Create Expense: Add a new expense record to the database.

Read Expense: Retrieve a single or list of expense records.

Update Expense: Modify an existing expense record.

Delete Expense: Remove an expense record from the database.

Each expense includes:

Transactionid: Unique identifier for the expense.

date: Date of the expense.

amount: Amount of the expense.

category: Category of the expense (e.g., Food, Travel, Utilities).

description: Additional details about the expense.

payment_method: Payment method used (e.g., Cash, Credit Card, PayPal).

created_at: Timestamp when the expense was created.

updated_at: Timestamp when the expense was last updated.

API Endpoints-
Base URL
Eg. https://localhost:3000

Endpoints:

Create an Expense-

Method: POST        
Endpoint: /expense

Get All Expenses-

Method: GET          
Endpoint: /expense


Get a Single Expense-

Method: GET        
Endpoint: /expense/{id}


Update an Expense-

Method: PUT          
Endpoint: /expense/{id}


Delete an Expense-

Method: DELETE             
Endpoint: /expenses/{id}

Technologies Used-

Database: [PostgreSQL]             
Other Tools: [Docker]
