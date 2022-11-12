# 编译相关信息
# 预定义目标文件
ifndef TARGET
TARGET:=demo
endif
SRC = $(wildcard *.go)

run: ${SRC}
	go run $^

build: ${SRC}
	go build -o ${TARGET} $^

clean:
	${RM} ${TARGET}

show:
	@echo TARGET: ${TARGET}
	@echo SRC: ${SRC}

.PHONY: run clean show build
