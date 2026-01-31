<# 
.SYNOPSIS
  Database migration helper for golang-migrate

.DESCRIPTION
  Thin, production-safe wrapper around golang-migrate.
  - Loads environment variables from .env
  - Runs migrations from project root
  - Protects against accidental destructive commands

.REQUIRES
  - golang-migrate installed and on PATH
  - DATABASE_URL set in .env
#>

# -----------------------------
# Path resolution
# -----------------------------
$ScriptDir   = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = Resolve-Path "$ScriptDir\.."
$EnvFile     = Join-Path $ProjectRoot ".env"
$Migrations  = "migrations"

# -----------------------------
# Helpers
# -----------------------------
function Fail($msg) {
    Write-Error $msg
    exit 1
}

function Confirm($msg) {
    Write-Host "$msg [y/N]"
    return (Read-Host) -eq 'y'
}

function Show-Help {
    Write-Host @"
Usage:
  .\migrate.ps1 up
  .\migrate.ps1 up <N>
  .\migrate.ps1 down [N]
  .\migrate.ps1 create <name>
  .\migrate.ps1 force <version>
  .\migrate.ps1 version
  .\migrate.ps1 drop
  .\migrate.ps1 help

Commands:
  up [N]        Apply all or N up migrations
  down [N]      Roll back 1 or N migrations
  create NAME   Create a new migration
  force V       Set migration version without running SQL
  version       Show current migration version
  drop          Drop all tables (DEV ONLY)
  help          Show this help

Examples:
  .\migrate.ps1 up
  .\migrate.ps1 down 1
  .\migrate.ps1 create add_users_table
  .\migrate.ps1 version
"@
}


# Load .env
if (-not (Test-Path $EnvFile)) {
    Fail ".env file not found at $EnvFile"
}

Get-Content $EnvFile | ForEach-Object {
    if ($_ -match '^\s*([^#=]+)\s*=\s*(.*)\s*$') {
        $key   = $matches[1].Trim()
        $value = $matches[2].Trim()

        if ($value.StartsWith('"') -and $value.EndsWith('"')) {
            $value = $value.Substring(1, $value.Length - 2)
        }

        Set-Item "env:$key" $value
    }
}

if (-not $env:DATABASE_URL) {
    Fail "DATABASE_URL is missing in .env"
}

# Argument parsing
$Command = $args[0]
$Arg     = $args[1]

if (-not $Command -or $Command -eq "help" -or $Command -eq "--help" -or $Command -eq "-h") {
    Show-Help
    exit 0
}

# Execute migrate from project root
Push-Location $ProjectRoot

try {
    switch ($Command) {

        "up" {
            if ($Arg) {
                migrate -path $Migrations -database $env:DATABASE_URL up $Arg
            } else {
                migrate -path $Migrations -database $env:DATABASE_URL up
            }
        }

        "down" {
            $count = if ($Arg) { $Arg } else { "1" }
            if (Confirm "Rolling back $count migration(s). Continue?") {
                migrate -path $Migrations -database $env:DATABASE_URL down $count
            }
        }

        "create" {
            if (-not $Arg) {
                Fail "Migration name required"
            }
            migrate create -ext sql -dir $Migrations -seq $Arg
        }

        "force" {
            if (-not $Arg) {
                Fail "Version number required"
            }
            migrate -path $Migrations -database $env:DATABASE_URL force $Arg
        }

        "version" {
            migrate -path $Migrations -database $env:DATABASE_URL version
        }

        "drop" {
            Write-Host "DANGEROUS: This will DROP ALL TABLES."
            if (Confirm "This should NEVER be run in production. Continue?") {
                migrate -path $Migrations -database $env:DATABASE_URL drop -f
            }
        }

        default {
            Write-Error "Unknown command: $Command"
            Show-Help
            exit 1
        }
    }
}
finally {
    Pop-Location
}