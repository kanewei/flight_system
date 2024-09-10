# flight systme

## Staff
The internal staff can operate the functions related to airplanes, airports, and flights.

## User
The user can operate functions for registration, flight search, ticket, and order.

## Airport
For internal use only, able to add and query airports.

## Airplane
For internal use only, able to add and query airplanes.

## Flight
Internal staff can create flights, and users can search for flights. Flight searches can be conducted with paginated queries.

After the flight is established, seats are also created, including overbooked seats. These overbooked seats, along with regular seats, are marked with an 'is_over_sold' tag. Users can book seats, but it will cost more. If they do not book, seats are assigned randomly, and there is a chance they may receive an overbooked seat.

## Ticket
Once the user selects a flight ticket, there is a 30-minute reservation period to make the payment. If the payment is completed within this time, the order will be finalized; otherwise, the ticket will be released.

## Order
When the flight ticket is purchased, an order will be created.