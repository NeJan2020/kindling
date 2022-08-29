docker rmi 10.10.102.213:8443/cloudnevro-test/kindling:tcp-queue-state
docker tag kindling-collector:latest 10.10.102.213:8443/cloudnevro-test/kindling:tcp-queue-state
docker images
docker push 10.10.102.213:8443/cloudnevro-test/kindling:tcp-queue-state
