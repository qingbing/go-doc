#!/bin/bash
# author qingbing<780042175@qq.com>
# 本脚步目的 ： 从当前目录自动生成 .md 文件的索引文件，并把文件放在当前目录
#     
# 本脚步使用方法
# ./make_index.sh

# 获取脚本所在目录
rPath=$(cd $(dirname $0); pwd)
# 加载公共函数
. $rPath/ShellLib/index.sh
loadFiles $rPath/ShellLib

# 索引文件名
indexFileName=README.md

indexFile=${rPath}/${indexFileName}

# 判断文件目录中是否包含 md 文件
function hasMdFile()
{
	local sPath=$1
	echo "Browse Path: ${sPath}"
	for _file in $(ls $sPath); do
		local filePath="${sPath}/${_file}"

		# 目录
		if [[ -d $filePath ]]; then
			hasMdFile $filePath
			if [[ $? -ne 0 ]]; then
				return 1
			fi
			continue
		fi

		local fileExt=`extensionname $_file`
		# 文件扩展判断为 md
		if [[ "${fileExt}" == "md" || "${fileExt}" == "MD" ]]; then
			return 1
		fi
	done
	return 0
}

# 目录循环
function generateMdIndex()
{
	local sPath=$1
	local prefix=$2
	local space=$3
	local ingoreFiles=$4

	# 定义变量
	local files=()
	local folders=()
	local fileIdx=0
	local folderIdx=0
	for file in $(ls $sPath); do
		# 忽略文件
		in_array $file "${ingoreFiles[*]}"
		isIngoreFile=$?
		if [[ $isIngoreFile -eq 1 ]]; then
			continue;
		fi
		# 文件路径
		filePath="${sPath}/${file}"

		# 目录
		if [[ -d $filePath ]]; then
			folders[folderIdx]=$filePath;
			let "folderIdx=$folderIdx + 1"
			continue;
		fi

		#文件扩展
		local fileExt=`extensionname $file`
		# 文件扩展判断为 md
		if [[ "${fileExt}" == "md" || "${fileExt}" == "MD" ]]; then
			files[fileIdx]=$filePath;
			let "fileIdx=$fileIdx + 1"
		fi
	done

	# 文件加载（执行）
	for file in ${files[*]}; do
		# 获取文件名
		relativeFilename=`basename $file`
		filename=`filename $relativeFilename`

		echo "${space}- [$filename]($prefix/$relativeFilename)" >> $indexFile
    done
    # 递归加载文件夹
	for file in ${folders[*]}; do
		basepath=`basename $file`
		hasMdFile $file;
		if [[ $? -eq 0 ]]; then
			continue;
		fi
		echo "${space}- $basepath" >> $indexFile
		generateMdIndex $file $prefix/$basepath "${space}    "
    done
}

echo "# go-doc
go 语言学习笔记
" > $indexFile;


ingoreFiles=($indexFileName tmp)

generateMdIndex $rPath '.' "" "${ingoreFiles[*]}"

