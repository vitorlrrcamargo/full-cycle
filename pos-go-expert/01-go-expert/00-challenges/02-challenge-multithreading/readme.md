# MULTITHREADING CHALLENGE

In this challenge, you will have to use what we have learned with Multithreading and APIs to find the fastest result between two different APIs.

The two requests will be made simultaneously to the following APIs:

https://brasilapi.com.br/api/cep/v1/ + zipCode

http://viacep.com.br/ws/" + zipCode + "/json/

The requirements for this challenge are:

- Accept the API that delivers the fastest response and discard the slowest response.

- The result of the request must be displayed on the command line with the address data, as well as which API sent it.

- Limit the response time to 1 second. Otherwise, the timeout error should be displayed.

# HOW TO RUN THE APPLICATION

In your browser or API tool (e.g. Postman), make a GET request to the / endpoint with the zip code parameter:

http://localhost:8080/?zipCode=13024091

The server will make the requests to both APIs and return the fastest response.
