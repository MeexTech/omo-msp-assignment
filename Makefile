.PHONY: proto
proto:
    protoc --proto_path=. --micro_out=. --go_out=. proto/organization/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/organization/task.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/organization/team.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/organization/volunteer.proto
