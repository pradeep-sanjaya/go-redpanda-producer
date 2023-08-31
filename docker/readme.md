### Start redpanda
docker compose up -d

### Get cluster info
docker exec -it redpanda-0 rpk cluster info

### Create topic
docker exec -it redpanda-0 rpk topic create topic1

### Produce message in topic1
docker exec -it redpanda-0 rpk topic produce topic1


### Consume messages from topic1
docker exec -it redpanda-0 rpk topic consume topic1 --num 1




