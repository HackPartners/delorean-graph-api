# Delorian Graph API

RESTful API to query the sations graph. Trailing slashes are tolerated. Query params are not case sensative, path segments are however.

## Station

Example station object returned by this API:

```
{
    "stanox":00000,
    "name":"EXAMPL"
}
```

## Stations

Returns an array of JSON objects representing the stations requested

**:from** is stanox of target station

**:direction** is either **up** or **down**

**GET** all stations:

```
{api_url}/stations
```

**GET** stations after target station:

```
{api_url}/stations/:from/:direction
```

```
{api_url}/stations?from=:from&?direction=:direction
```

## Paths

Returns a 2d array of JSON objects represtenting the stations in the paths requested

**:from**,**:to** is stanox of target station(s)

**:direction** is either **up** or **down**

**GET** paths after target station:

```
{api_url}/paths/:from/:direction
```

```
{api_url}/paths?from=:from&?direction=:direction
```

**GET** paths between target stations:

```
{api_url}/paths/:from/to/:to/:direction
```

```
{api_url}/paths?from=:from&to=:to&?direction=:direction
```
