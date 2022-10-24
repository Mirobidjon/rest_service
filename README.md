# rest_service

# About

Rest Service stores phone numbers (ID/PHONE) in its postgreSQL database and provides a REST API to access them.

Basic Rest Service (You can use the existing one or create your own). Five endpoints for the following operations:
    1. Create a new phone number    (/phone POST)
    2. Get a phone number by ID     (/phone/{id} GET)
    3. Update a phone number by ID  (/phone/{id} PUT)
    4. Delete a phone number by ID  (/phone/{id} DELETE)
    5. Get all phone numbers        (/phone GET)


Use the following command to run the application:

    docker run -p 8080:80 us.gcr.io/learn-cloud-0809/rest_service:latest

# ENVIRONMENT VARIABLES

The following environment variables are used by the application, you can use .env.example as a template:

    POSTGRES_HOST: The host of the database
    POSTGRES_PORT: The port of the database
    POSTGRES_DATABASE: The name of the database
    POSTGRES_USER: The user of the database
    POSTGRES_PASSWORD: The password of the database

