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

### Relationships
The events.organizer_id field in the Events table establishes a relationship with the users.id field in the Users table, indicating the user who created the event.

The Registrations table links users and events. The registrations.user_id field is related to the users.id field, and the registrations.event_id field is related to the events.id field.
*/
