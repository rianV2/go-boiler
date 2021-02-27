# Go Boiler API

### POST /player

Add player in DB. 

Body | Type | Description | Example
--- | --- | --- | ---
name (required) | string | Name of player. | "Rastafara Uye"
level | string | Level of player. | "1"
job | string | Job of player. | "Sorcerer"

Example

`/POST localhost:9000/players`

Request
```
{   
	"name": "Rastafara Uye",
	"level": "1",
	"job": "Sorcerer"
}
```

HTTP 200

```
{
	"name": "Rastafara Uye",
	"level": "1",
	"job": "Sorcerer"
}
```

---