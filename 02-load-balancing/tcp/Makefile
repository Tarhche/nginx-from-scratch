default: run

up start:
	docker compose up -d --build

down stop:
	docker compose down --remove-orphans --volumes

logs log:
	docker compose logs load-balancer

lb-terminal:
	docker exec -it load-balancer sh

lb-mysql-terminal:
	docker exec -it load-balancer mysql -P 3305 -u root -proot

db1-terminal:
	docker exec -it mysql-node-01 mysql -u root -proot

db2-terminal:
	docker exec -it mysql-node-02 mysql -u root -proot

.PHONY: up start down stop logs log lb-terminal db1-terminal db2-terminal
