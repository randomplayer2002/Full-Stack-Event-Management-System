/*
This table holds details about the events created in the system.
*/

CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    title TEXT,
    description TEXT,
    date TIMESTAMP,
    location TEXT,
    capacity INT,
    organizer_id UUID
);

/*
id: A unique identifier for each event.
title: The title or name of the event.
description: A brief description of the event.
date: The date and time of the event.
location: The venue or location where the event will take place.
capacity: The maximum number of attendees allowed for the event.
organizer_id: The ID of the user who created the event, establishing a relationship with the Users table.
*/
