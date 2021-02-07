trap killgroup SIGINT

killgroup(){
  echo killing...
  kill 0
}

docker run --rm -p 5434:5432 -e POSTGRES_PASSWORD=postgres --name db_performance_test postgres:12-alpine