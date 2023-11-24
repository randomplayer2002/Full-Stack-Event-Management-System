/*
This table stores information about registered users in the system.
*/

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username TEXT,
    password TEXT,
    email TEXT
);

/*

id: A unique identifier for each user.
username: The username chosen by the user during registration.
password: The hashed password of the user for security.
email: The email address of the user.

*/
