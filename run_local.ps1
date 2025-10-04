$env:DB_HOST = "localhost"
$env:DB_PORT = "5433"
$env:DB_USER = "postgres"
$env:DB_PASSWORD = "Kekos228@"
$env:DB_NAME = "tasks_db"
$env:DB_SSLMODE = "disable"
$env:PORT = "8080"

Write-Host "======================================" -ForegroundColor Cyan
Write-Host "Starting local development server..." -ForegroundColor Green
Write-Host "DB: postgresql://$env:DB_USER@$env:DB_HOST:$env:DB_PORT/$env:DB_NAME" -ForegroundColor Yellow
Write-Host "API: http://localhost:$env:PORT" -ForegroundColor Yellow
Write-Host "======================================" -ForegroundColor Cyan
Write-Host ""

go run cmd/server/main.go

