CURRENT_DIR=$(shell pwd)


proto-gen:
#script file orqali proto fileni hammasini bittada generatsiya qilish 
	./scripts/gen-proto.sh "${CURRENT_DIR}"
	
exp:
	export DBURL='postgres://postgres:root@localhost:5432/uyishi47_db?sslmode=disable'

mig-up:
	migrate -path migrations -database ${DBURL} -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down



# mig-delete:
#   rm -r db/migrations




gen-proto:
# proto fileni bittalab ruchnoy generatsiya qilish 
	protoc --go_out=services/ \
       --go-grpc_out=services/ \
       services/protos/cart.proto


gen-protoAll:
#proto fileni hammasini bittada generatsiya qilish 
	protoc --go_out=genproto/ \
    --go-grpc_out=genproto/ \
    protos/product/*.proto



gen-proto_curr:
# bu proto fayllarimiz turgan joyga generatsiya qilib beradi bittalab
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		helloworld/helloworld.proto