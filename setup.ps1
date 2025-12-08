# --- CONFIGURATION ---
$RepoUrl = "https://github.com/w3rr0/MyLC.git"
$DirName = "server"

Write-Host "=== Starting MyLC environment setup ===" -ForegroundColor Cyan

# 1. Check whether Docker is running
docker info > $null 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Docker is not running!" -ForegroundColor Red
    Write-Host "Launch Docker Desktop and try again."
    exit
}

# 2. Clone or update the repository
if (Test-Path -Path $DirName) {
    Write-Host "Directory $DirName already exist. Downloading changes (git pull)..." -ForegroundColor Cyan
    Set-Location -Path $DirName
    git pull
} else {
    Write-Host "Cloning repository..." -ForegroundColor Green
    git clone $RepoUrl $DirName
    Set-Location -Path $DirName
}

# 3. ULaunch Docker Compose
Write-Host "🐳 Building and launching project inside Docker..." -ForegroundColor Green
docker compose down -v
docker compose up --build