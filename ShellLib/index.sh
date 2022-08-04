#!/bin/bash

#################### 递归记载目录 ####################
function loadFiles()
{
	# 要加载目录
	sPath=$1
	if [[ -z $sPath ]]; then
		# 当未传递参数时定位到当前脚本目录的 helpers 下
		echo 'loading 目录时，必须指定目录, eg: "loadFiles $sPath/lib"'
		exit 1
	fi
	# 定义变量
	files=()
	folders=()
	fileIdx=0
	folderIdx=0
	for file in $(ls $sPath); do
		_filePath="${sPath}/${file}"
		# 目录
		if [[ -d $_filePath ]]; then
			folders[folderIdx]=$_filePath;
			let "folderIdx=$folderIdx + 1"
			continue;
		fi
		# 可执行文件需要加载
		if [[ -x $_filePath ]]; then
			files[fileIdx]=$_filePath;
			let "fileIdx=$fileIdx + 1"
			continue;
		fi
	done
	# 文件加载（执行）
	for file in ${files[*]}; do
    	. $file
    done
    # 递归加载文件夹
	for folder in ${folders[*]}; do
		loadFiles $folder
    done
}

# 当前脚本目录
# rPath=$(cd $(dirname $0); pwd)
# 库文件目录
# hPath="${sPath}/helpers"
# # 加载 helpers 目录里面的所有可执行文件
# loadFiles $hPath
