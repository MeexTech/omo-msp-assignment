.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/task.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/team.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/agent.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/family.proto
