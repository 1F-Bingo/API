## 1F Bingo

> *Mom! I'm playing bingo, leave me be.*

---

This platform serves as the API gateway for end clients to consume. Documentation lives [here](https://1f.skelmis.co.nz/api/v1/docs).

### MVP

A site that allows users to create a bingo board from a collection of possible entries. Users can compete against friends in leagues with a leaderboard for most pieces as well as who gets bingo first.

Admins can update bingo item entries and have updates push out to end clients easily.

##### Considerations

- Bingo board randomization should be done server side to avoid end users selecting ideal boards via proxying.
- A maximum number of randomization's should be offered to stop people just re-rolling endlessly until they get one the best one.

---

### Taking it further

Users can create custom bingo items for their own leagues

---

### Running the API

#### Both

Environment variables required to run

- `TOKEN_HOUR_LIFESPAN=int`, JWT token lifetime in hours
- `API_SECRET=string`, API secret to sign JWT's with

#### Developing

Run the server locally using `go build`, documentation can be generated using `swag init`.

#### Production

Build and run the `Dockerfile`, when running the generated Docker image you should also set the following enviroment variables:

- `GIN_MODE=release`
- `PROXY_IP=<Your proxy ip>`