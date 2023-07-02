curl -X POST -H "Content-Type: application/json" -d '{"id":"1","title":"Article 1","description":"This is article 1","rate":5}' http://localhost:5000/api/v1/articles

curl -X GET http://localhost:5000/api/v1/articles/1

curl -X GET http://localhost:5000/api/v1/articles/

curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Article 1","description":"This is the updated article 1","rate":10}' http://localhost:5000/api/v1/articles/1
