/*
This table keeps track of users who have registered for events.
*/

CREATE TABLE IF NOT EXISTS registrations (
    user_id UUID,
    event_id UUID,
    PRIMARY KEY (user_id, event_id)
);

/*
user_id: The ID of the user who registered for an event, establishing a relationship with the Users table.
event_id: The ID of the event for which the user registered, establishing a relationship with the Events table.
*/
