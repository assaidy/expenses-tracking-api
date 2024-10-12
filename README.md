# Expenses tracking api with JWT auth.

### User

#### 1. Register 
- Method: `POST /users/register`
- Access level: public
- Expected responses:
    - `201` created
    - `400` bad request
    - `409` conflict

#### 2. Login 
- Method: `POST /users/login` 
- Access level: public
- Expected responses:
    - `200` ok
    - `400` bad request
    - `401` unauthorized
- returns jwt token

#### 3. Get profile
- Method: `GET /users`
- Access level: protected
- Expected responses:
    - `200` ok
    - `401` unauthorized
    - `404` not found

#### 4. Update profile
- Method: `PUT /users` 
- Access level: protected
- Expected responses:
    - `200` ok
    - `400` bad request
    - `401` unauthorized
    - `409` conflict

#### 5. Delete profile
- Method: `DELETE /users` 
- Access level: protected
- Expected responses:
    - `200` ok
    - `401` unauthorized

### Expenses

#### 1. Create Expense
- Method: `POST /expenses` 
- Access level: protected
- Expected responses:
    - `201` created
    - `401` unauthorized
    - `400` bad request
    - `404` category not found

#### 2. Update Expense
- Method: `PUT /expenses/:id<int>` 
- Access level: protected
- Expected responses:
    - `200` ok
    - `401` unauthorized
    - `400` bad request
    - `404` expense/category not found

#### 3. Delete Expense
- Method: `DELETE /expenses/:id<int>` 
- Access level: protected
- Expected responses:
    - `200` ok
    - `401` unauthorized
    - `404` expense not found

#### 4. Get All Expenses
- Method: `GET /expenses` 
- Access level: protected
- Filters:
    - page: default 1
    - limit: default 10
    - date_ragne: week, month, 3months, custome [start_date & end_date]
- Expected responses:
    - `200` ok
    - `401` unauthorized
    - `400` bad request
