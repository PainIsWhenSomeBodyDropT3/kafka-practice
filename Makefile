zoo_up:
	# docker build -t zookeeper:3.9.1 -f D-zookeeper.yaml --no-cache .
	docker-compose -f DC-kafka-zookeeper-cluster.yaml up -d

zoo_down:
	docker-compose -f DC-kafka-zookeeper-cluster.yaml down 
	 
zoo_restart: zoo_down zoo_up 



kraft_up:
	# docker build -t consumer:1.0 consumer/ #--no-cache producer/
	# docker build -t producer:1.0 producer/ #--no-cache producer/
	docker-compose -f DC-kafka_kraft.yaml up -d

kraft_down:
	docker-compose -f DC-kafka_kraft.yaml down

kraft_restart: kraft_down kraft_up
