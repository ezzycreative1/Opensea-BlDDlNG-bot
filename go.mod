{
	"info": {
		"_postman_id": "687f8d9e-d917-4eae-af62-49f4d975c653",
		"name": "solution auth jwt",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "551676"
	},
	"item": [
		{
			"name": "login",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "formdata",
					"formdata": [
						{
							"key": "username",
							"value": "user123",
							"type": "text"
						},
						{
							"key": "password",
							"value": "password123",
							"type": "text"
						}
					]
				},
				"url": {
					"raw": "http://localhost:3000/login",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "3000",
					"path": [
						"login"
					]
				}
			},
			"response": []
		},
		{
			"name": "get auth",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": [
						{
							"key": "token",
							"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxMjMiLCJleHAiOjE3MDk2MTQ0ODl9.uhq4GxCofnAviHwcrQSovogtIFs0UvvoumUJIapo8dY\"",
							"type": "string"
						}
					]
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:3000/auth",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "3000",
					"path": [
						"auth"
					]
				}
			},
			"response": []
		}
	]
}