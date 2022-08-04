#!/bin/bash

# 检查必要的目录是否存在，不存在则程序终止
function file_check_dir()
{
	if [[ ! -d $1 ]]; then
		echo $1
		echo "====== 程序终止 ======"
		exit
	fi
}

# 确定目录存在，不存在就创建
function file_create_dir()
{
	if [[ ! -d $1 ]]; then
		mkdir -p $1
	fi
}

# 获取路径中的文件名，不带后缀
function filename()
{
	file=$1
	echo "${file%.*}"
}

# 获取文件名的后缀
function extensionname()
{
	file=$1
	echo "${file##*.}"
}
