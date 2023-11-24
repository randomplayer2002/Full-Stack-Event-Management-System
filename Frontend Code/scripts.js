// scripts.js

document.addEventListener('DOMContentLoaded', function () {
    // Fetch and display events on page load
    fetchEvents();

    // Add event listeners for interactions
    document.getElementById('event-list').addEventListener('click', handleEventClick);
});

function fetchEvents() {
    // Make an API call to get the list of events and dynamically populate the event list section
    // Example:
    // const events = fetch('/api/events').then(response => response.json());
    // events.forEach(event => appendEventCard(event));
}

function appendEventCard(event) {
    // Dynamically create and append event cards to the event list section
    const eventList = document.getElementById('event-list');
    const eventCard = document.createElement('div');
    eventCard.className = 'event-card';
    eventCard.innerHTML = `
        <h3>${event.title}</h3>
        <p>${event.date}</p>
    `;
    eventList.appendChild(eventCard);
}

function handleEventClick(event) {
    // Handle the event card click event to display event details or registration form
    // Example:
    // const eventId = event.target.dataset.id;
    // const eventDetails = fetch(`/api/events/${eventId}`).then(response => response.json());
    // displayEventDetails(eventDetails);
}

function displayEventDetails(event) {
    // Display event details in the event details section
    const eventDetailsSection = document.getElementById('event-details');
    eventDetailsSection.innerHTML = `
        <h2>${event.title}</h2>
        <p>Date: ${event.date}</p>
        <p>Location: ${event.location}</p>
        <p>Description: ${event.description}</p>
        <button onclick="showRegistrationForm('${event.id}')">Register</button>
    `;

    // Show the event details section and hide the others
    document.getElementById('event-list').classList.add('hidden');
    document.getElementById('event-details').classList.remove('hidden');
    document.getElementById('registration-form').classList.add('hidden');
}

function showRegistrationForm(eventId) {
    // Show the registration form for the selected event
    // You can handle the registration logic here
    document.getElementById('event-list').classList.add('hidden');
    document.getElementById('event-details').classList.add('hidden');
    document.getElementById('registration-form').classList.remove('hidden');
}
