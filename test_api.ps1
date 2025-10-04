$baseUrl = "http://localhost:8080"

Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "       TESTING TODO API" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

Write-Host "1. Health Check" -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/health" -Method Get
    Write-Host "   Status: $($response.status)" -ForegroundColor Green
    Write-Host "   Time: $($response.time)" -ForegroundColor Gray
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

Write-Host "`n2. Creating Task #1: 'Купить молоко'" -ForegroundColor Yellow
try {
    $task1 = @{
        name = "Купить молоко"
        description = "В магазине на углу"
    } | ConvertTo-Json

    $response1 = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method Post -Body $task1 -ContentType "application/json"
    $taskId1 = $response1.id
    Write-Host "   Created Task ID: $taskId1" -ForegroundColor Green
    Write-Host "   Name: $($response1.name)" -ForegroundColor Gray
    Write-Host "   Completed: $($response1.completed)" -ForegroundColor Gray
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n3. Creating Task #2: 'Выучить Go'" -ForegroundColor Yellow
try {
    $task2 = @{
        name = "Выучить Go"
        description = "Изучить основы языка Go"
    } | ConvertTo-Json

    $response2 = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method Post -Body $task2 -ContentType "application/json"
    $taskId2 = $response2.id
    Write-Host "   Created Task ID: $taskId2" -ForegroundColor Green
    Write-Host "   Name: $($response2.name)" -ForegroundColor Gray
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n4. Getting All Tasks (GET /tasks)" -ForegroundColor Yellow
try {
    $allTasks = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method Get
    Write-Host "   Total tasks: $($allTasks.total)" -ForegroundColor Green
    foreach ($task in $allTasks.tasks) {
        Write-Host "   - [ID:$($task.id)] $($task.name) - Completed: $($task.completed)" -ForegroundColor Gray
    }
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n5. Getting Task by ID (GET /tasks/$taskId1)" -ForegroundColor Yellow
try {
    $task = Invoke-RestMethod -Uri "$baseUrl/tasks/$taskId1" -Method Get
    Write-Host "   Task: $($task.name)" -ForegroundColor Green
    Write-Host "   Description: $($task.description)" -ForegroundColor Gray
    Write-Host "   Completed: $($task.completed)" -ForegroundColor Gray
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n6. Updating Task (PUT /tasks/$taskId1)" -ForegroundColor Yellow
try {
    $updateTask = @{
        name = "Купить молоко и хлеб"
        description = "В супермаркете на главной улице"
        completed = $false
    } | ConvertTo-Json

    $updated = Invoke-RestMethod -Uri "$baseUrl/tasks/$taskId1" -Method Put -Body $updateTask -ContentType "application/json"
    Write-Host "   Updated: $($updated.name)" -ForegroundColor Green
    Write-Host "   Description: $($updated.description)" -ForegroundColor Gray
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n7. Marking Task as Complete (PATCH /tasks/$taskId2/complete)" -ForegroundColor Yellow
try {
    $completed = Invoke-RestMethod -Uri "$baseUrl/tasks/$taskId2/complete" -Method Patch
    Write-Host "   Task '$($completed.name)' marked as completed" -ForegroundColor Green
    Write-Host "   Completed: $($completed.completed)" -ForegroundColor Gray
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n8. Getting Completed Tasks (GET /tasks/completed)" -ForegroundColor Yellow
try {
    $completedTasks = Invoke-RestMethod -Uri "$baseUrl/tasks/completed" -Method Get
    Write-Host "   Completed tasks: $($completedTasks.total)" -ForegroundColor Green
    foreach ($task in $completedTasks.tasks) {
        Write-Host "   - [ID:$($task.id)] $($task.name)" -ForegroundColor Gray
    }
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n9. Getting Pending Tasks (GET /tasks/pending)" -ForegroundColor Yellow
try {
    $pendingTasks = Invoke-RestMethod -Uri "$baseUrl/tasks/pending" -Method Get
    Write-Host "   Pending tasks: $($pendingTasks.total)" -ForegroundColor Green
    foreach ($task in $pendingTasks.tasks) {
        Write-Host "   - [ID:$($task.id)] $($task.name)" -ForegroundColor Gray
    }
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n10. Deleting Task (DELETE /tasks/$taskId1)" -ForegroundColor Yellow
try {
    Invoke-RestMethod -Uri "$baseUrl/tasks/$taskId1" -Method Delete
    Write-Host "   Task deleted successfully" -ForegroundColor Green
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n11. Verifying Deletion (GET /tasks)" -ForegroundColor Yellow
try {
    $finalTasks = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method Get
    Write-Host "   Remaining tasks: $($finalTasks.total)" -ForegroundColor Green
    foreach ($task in $finalTasks.tasks) {
        Write-Host "   - [ID:$($task.id)] $($task.name) - Completed: $($task.completed)" -ForegroundColor Gray
    }
} catch {
    Write-Host "   FAILED: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "       TESTING COMPLETED!" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

