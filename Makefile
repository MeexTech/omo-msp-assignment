.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/task.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/team.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/agent.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/family.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/coterie.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/apply.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/meeting.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/category.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/assignment/question.proto
