
if [ "$(docker ps -q -f name=HMS-mysql-db)" ]; then
  docker rm -f HMS-mysql-db
fi

if [ "$(docker ps -q -f name=HMS-mongo-db)" ]; then
  docker rm -f HMS-mongo-db
fi
