# Testing Guide

This document describes the unit test suite for the Task Management API, how to run it locally.

## Overview

- Test framework: `testify` (assert and mock)
- Language/runtime: Go 1.21
- Targets covered:
  - Domain: entity validation and helpers (`Task.Validate`, `User.IsAdmin`)
  - Usecases: business rules with mocked repositories/services (tasks and users)
  - Controllers: HTTP handlers using Gin’s test context with mocked usecases
  - Infrastructure: password hashing and JWT token generation/validationalternatively, use mongo-driver `mtest` integration helpers

## Running tests

- Run all tests with verbose output:
  ```bash
  go test ./... -v
  ```
- Only a package:
  ```bash
  go test ./Usecases -v
  go test ./Delivery/controllers -v
  ```
- Coverage report:
  ```bash
  go test ./... -coverprofile=coverage.out
  go tool cover -func=coverage.out
  # HTML report
  go tool cover -html=coverage.out -o coverage.html
  ```

## What’s tested

- Domain
  - Enforces non-empty title/description
  - Admin role helper
- Usecases (with testify mocks)
  - Tasks: list, get by id, create/update validation, delete, not-found
  - Users: register (hash persisted), login (success, wrong password, user not found), promote
- Controllers (with Gin + mocked usecases)
  - Tasks: list, get by id (ok/invalid/not found), create (validation/success), update (invalid id/not found), delete (invalid id/success)
  - Users: register, login (success/invalid), promote
- Infrastructure
  - Password: bcrypt hashing and comparison, wrong password branch
  - JWT: generate/validate roundtrip, malformed token, expired token

## Edge cases covered

- Task: missing title/description, not found on get/delete, update validations
- Auth: invalid credentials (wrong password or user missing)
- JWT: malformed or expired token


## Conventions and tips

- Test file names: `*_test.go`
- Use table-driven tests for variants where helpful
- Prefer constructor helpers for mocks to reduce duplication
- Keep assertions specific and high-signal; prefer `assert.ErrorIs` for domain error types
- When adding new usecases/controllers, create mocks for dependencies and cover
  - success path
  - validation errors
  - not-found or permission errors

## Troubleshooting

- IDE shows broken imports for tests: run `go mod tidy` at `task-manager/`
- JWT errors in tests with v5: use exported sentinels like `jwt.ErrTokenMalformed`, `jwt.ErrTokenExpired`
- Mongo `mtest` errors: ensure correct imports and consider isolating behind build tags or switch to interface-based mocks
