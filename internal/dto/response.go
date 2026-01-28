package dto

// Response is the standard API response format
type Response struct {
	Code    string      `json:"code" example:"SUCCESS"`
	Message string      `json:"message" example:"Operation completed"`
	Data    interface{} `json:"data"`
}

// ErrorResponse is the standard error response format
type ErrorResponse struct {
	Code    string `json:"code" example:"VALIDATION_ERROR"`
	Message string `json:"message" example:"Invalid input"`
	Data    any    `json:"data"`
}

// HealthResponse is the health check response
type HealthResponse struct {
	Status    string `json:"status" example:"ok"`
	Version   string `json:"version" example:"1.0.0"`
	Timestamp string `json:"timestamp" example:"2024-01-18T10:00:00Z"`
}

// DocsResponse is the docs endpoint response
type DocsResponse struct {
	Content string `json:"content"`
}

// PingResponse is a simple ping response
type PingResponse struct {
	Message   string `json:"message" example:"pong"`
	RequestID string `json:"request_id" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// EchoRequest is the echo endpoint request
type EchoRequest struct {
	Message string `json:"message" binding:"required" example:"Hello World"`
}

// EchoResponse is the echo endpoint response
type EchoResponse struct {
	Original  string `json:"original" example:"Hello World"`
	Reversed  string `json:"reversed" example:"dlroW olleH"`
	Length    int    `json:"length" example:"11"`
	RequestID string `json:"request_id" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// LoginRequest is the login endpoint request
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// LoginResponse is the login endpoint response
type LoginResponse struct {
	Token     string `json:"token" example:"eyJhbGciOiJIUzI1NiIs..."`
	ExpiresAt string `json:"expires_at" example:"2024-01-19T10:00:00Z"`
}

// UserInfoResponse is the protected user info response
type UserInfoResponse struct {
	UserID    string `json:"user_id" example:"user-123"`
	RequestID string `json:"request_id" example:"550e8400-e29b-41d4-a716-446655440000"`
}
