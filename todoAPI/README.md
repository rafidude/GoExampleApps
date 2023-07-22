curl localhost:3000/todos

curl -X POST -H "Content-Type: application/json" -d '{"title":"My Todo", "completed":false}' http://localhost:3000/todos

curl -X PUT -H "Content-Type: application/json" -d '{"id": "1ff9df5d-80fa-409d-bed3-e68b944cd962", "title":"Updated Todo", "completed":true}' http://localhost:3000/todos

curl -X DELETE http://localhost:3000/todos/fa7db9bb-07fe-45b9-ab3a-a6ca89bcdeb5
