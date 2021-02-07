trap killgroup SIGINT

killgroup(){
  echo killing...
  kill 0
}

docker exec -i db_performance_test psql -U postgres -c "drop database if exists pg_performance_test" && \
docker exec -i db_performance_test psql -U postgres -c "create database pg_performance_test"