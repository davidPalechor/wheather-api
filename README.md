# Weather status API
## Dependencies
- Docker 18.06.1
- Docker compose 1.22.0

## Running app
- `docker-compose build`
- `docker-compose up`
- Go to `localhost:8080/v1/{$ENDPOINT}`
### Available endpoints:
- [GET] `weather?city=$City&country=$Country` where $City is city's name capitalized e.g. Bogota, London, etc. And $Country is the country code lowercased e.g. co (Colombia), uk (United Kingdom), ca (Canada), etc. (See a full code list [here](https://www.nationsonline.org/oneworld/country_code_list.html))
- [PUT]`scheduler/weather` Params - {city: $City, country: $Country} following rules above.

## Endpoint Swagger Docs
- `localhost:8080/swagger`

## Applying migrations
- `docker exec -it wheatherapi_web_1 bee migrate -conn="root:root@tcp(db:3306)/weather"`

## Testing 
- `docker exec -it wheatherapi_web_1 bash test.sh`

## MySQL shell
- `docker exec -it wheatherapi_db_1 mysql -u root -proot"`
