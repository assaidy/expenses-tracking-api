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
