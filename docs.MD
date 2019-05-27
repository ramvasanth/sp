# for creating a new relation
make a POST request to this URL http://localhost:3000/relation/create with the following payload and set the Content-Type to application/json in the header 
``
{
  "friends":
    [
      "andy@example.com",
      "john@example.com"
    ]
}
``
# for getting friends list
make a POST request to this URL http://localhost:3000/relation/get with the following payload and set the Content-Type to application/json in the header 
``
{
  "email":"andy@example.com"
}
``

# for getting a common friend between two friends email
make a POST request to this URL http://localhost:3000/relation/common with the following payload and set the Content-Type to application/json in the header 
``
{
  "friends":
    [
      "andy@example.com",
      "john@example.com"
    ]
}
``

#to subscribe to updates from an email address
make a POST request to this URL http://localhost:3000/subscription/create with the following payload and set the Content-Type to application/json in the header 
``
{
  "requestor": "andy@example.com",
  "target":"john@example.com"
}
``

#to block updates from an email address
make a DELETE request to this URL http://localhost:3000/subscription/block with the following payload and set the Content-Type to application/json in the header 
``
{
  "requestor": "andy@example.com",
  "target":"john@example.com"
}
``