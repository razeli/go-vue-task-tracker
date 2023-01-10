#!/bin/bash

cd vue-frontend
npm run build

cd ..
rm -rf app/dist
cp -R vue-frontend/dist app/ui/

docker-compose down

docker-compose up -d --build

echo "you can access the application now"
echo "http://<your ip>:8080"
