#!/bin/sh

curl -X GET http://localhost:3333/artistes
curl -X POST http://localhost:3333/artistes
curl -X PATCH http://localhost:3333/artistes/10
curl -X DELETE http://localhost:3333/artistes/10
curl -X GET http://localhost:3333/artistes/10/albums
curl -X POST http://localhost:3333/artistes/10/albums
curl -X PATCH http://localhost:3333/artistes/10/albums/10
curl -X DELETE http://localhost:3333/artistes/10/albums/10
curl -X GET http://localhost:3333/artistes/10/acts
curl -X POST http://localhost:3333/artistes/10/acts
curl -X GET http://localhost:3333/artistes/10/members

