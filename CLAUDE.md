# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based tabletop RPG character creation and simulation system called "gottrpg". The project models D&D-style character creation with attributes, classes, races, spells, and dice mechanics.

## Commands

### Build and Run
```bash
go build -o gottrpg main.go  # Build the application
./gottrpg                     # Run the built binary
go run main.go               # Run directly
```

### Testing
```bash
go test ./tests/...          # Run all tests
go test ./tests/diceroll_test.go  # Run specific test file
go test -v                   # Run tests with verbose output
```

### Module Management
```bash
go mod tidy                  # Clean up dependencies
go mod download              # Download dependencies
```

## Architecture

### Core Components

**Server Framework (`server/`)**
- Custom HTTP server wrapper around Go's standard library
- Router system for URL pattern matching and route registration
- Fluent API for server configuration (timeouts, headers, etc.)
- Package name: `goback` (appears to be a custom web framework)

**Character System (`character/`)**
- Complete D&D 5e character modeling with:
  - Attributes, proficiencies, saving throws
  - Multi-classing support via slice of Class structs
  - Race/subrace system with ability score bonuses
  - Background system with skills and languages
  - Spell system and feats

**Simulation Engine (`simulation/`)**
- Polyhedral dice definitions (d4, d6, d8, d10, d12, d20)
- Statistical functions for dice distributions
- Level progression and mathematical utilities
- Comprehensive test coverage for dice mechanics

**Routing (`routes/`)**
- RESTful API endpoints for character operations
- Currently minimal implementation with placeholder character route

**Logging (`log/`)**
- Custom logging system using logrus
- Structured logging with entry management

### Key Patterns

- Uses custom types for enums (Morality, SocietalAttitude)
- Fluent API pattern for server configuration
- Comprehensive struct modeling for game mechanics
- Statistical simulation for dice behavior

### Dependencies

- `github.com/sirupsen/logrus` - Structured logging
- `github.com/mazen160/go-random` - Random number generation
- Go 1.22.0 required

### Current State

The project appears to be in early development with:
- Basic server infrastructure in place
- Character model structure defined
- Dice simulation system implemented
- Main function contains unrelated algorithm practice code that should be cleaned up
- Server initialization is commented out in main.go