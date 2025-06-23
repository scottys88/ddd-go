# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Domain-Driven Design (DDD) implementation in Go for what appears to be an event management system (likely sports/running events). The codebase follows DDD principles with clear separation between entities and value objects.

## Code Architecture

### Directory Structure
- `entity/` - Contains original domain entities (kept for reference)
- `valueobject/` - Contains value objects without identity (Transaction, EntrantStatus, DistanceValue, Duration)
- `aggregates/` - Contains aggregate roots and their child entities (Event aggregate)
- `domain/` - Currently empty, likely for domain services

### Key Domain Concepts

#### Event Aggregate (aggregates/event.go)
- **Event** (Aggregate Root): Contains event details, dates, and collections of distances and entrants
- **Distance**: Child entity representing race distances with entries
- **Entry**: Child entity representing individual race results with times and positions  
- **Entrant**: Child entity linking persons to events with registration details

#### Value Objects (valueobject/)
- **EntrantStatus**: Enum for registration/race status (registered, confirmed, started, finished, dnf, dq)
- **DistanceValue**: Distance measurement with value and unit
- **Duration**: Race time representation with hours, minutes, seconds
- **Transaction**: Value object for transfers between entities

#### Standalone Entities (entity/)
- **Person**: Individual participants with contact information (separate aggregate)
- **Medal**: Awards/medals for events
- **Organiser**: Links persons to events they organize

### Design Patterns
- Uses `github.com/google/uuid` for entity identifiers
- Entities have clear identity (ID fields)
- Value objects follow DDD principles (Transaction has no exposed identity)
- Consistent use of JSON tags for serialization
- Audit fields pattern in Entry entity (CreatedAt, UpdatedAt, ArchivedAt)

## Development Commands

### Basic Go Commands
```bash
# Build the project
go build ./...

# Run tests (when available)
go test ./...

# Run tests with verbose output
go test -v ./...

# Check for dependency issues
go mod tidy

# Verify dependencies
go mod verify
```

### Module Information
- Module: `github.com/scottys88/ddd-go`
- Go version: 1.20
- Dependencies: `github.com/google/uuid v1.6.0`

## Notes
- No tests currently exist in the codebase
- The `entrantstatus.go` file exists but is empty
- Aggregates and domain directories are present but empty
- This appears to be an early-stage DDD implementation focused on defining core entities