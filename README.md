# Shopping Basket Manager Service 

## Overview

This repository contains the code for the Shopping Basket Service project, developed as the Internet Engineering course midterm project during the first semester of 1402-1403 Azar.


## Introduction

The project focuses on implementing a shopping basket maintenance service. It provides functionality to manage users' shopping baskets via APIs.

Features
--------

The implemented service offers the following API endpoints:

- `GET /basket/`: Retrieves a list of baskets
- `POST /basket/`: Creates a new basket
- `PATCH /basket/<id>`: Updates the specified basket
- `GET /basket/<id>`: Retrieves a specific basket
- `DELETE /basket/<id>`: Deletes the specified basket

Project Structure
-----------------

The project structure is organized as follows:

- `/handler`: Contains HTTP request handlers for managing baskets
- `/model`: Includes data models and repository interfaces for database operations
- `/request`: Stores request/response structures
- `/config`: Configuration files

Implementation Details
----------------------

The shopping basket structure includes the following information:

- `id`
- `created_at`: Time when the basket was created
- `updated_at`: Time when the basket was last updated
- `data`: Variable length data (maximum size: 2048 bytes)
- `state`: Status of the basket (COMPLETED or PENDING)

- The project supports basket creation and updates, with restrictions on modifying COMPLETED baskets.
- Implementation involves specifying data and state in basket creation and updates; other attributes are managed by the backend.
- Adding User: Involves adding users to the user table, associating users with baskets, and implementing token-based authentication.

### User Authentication and Permissions
Each user can perform CRUD operations only on their respective baskets after authentication.
Token-based(JWT) authentication is implemented to authenticate users.
Users receive a token upon successful login, which they must use in API requests to access their baskets.
CRUD operations on baskets are restricted to the authenticated user.

### Backend Implementation

- Database: PostgreSQL
- Implemented in Go

How to Run
----------

Follow these steps to set up and run the project on your local machine:

#### Prerequisites
- Ensure you have Go installed on your system.
- Have Git available on your machine.

#### Steps

1. **Install Dependencies**

    ```bash
    go mod tidy
    ```

2. **Set Up the Environment**

    ```bash
    make up
    ```

3. **Apply Database Migrations**

    ```bash
    make migrate-up
    ```

4. **Run the Server**

    ```bash
    make run-server
    ```
