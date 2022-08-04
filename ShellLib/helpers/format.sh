#!/bin/bash

# 当前时间格式化，使用方法
# time=`nowtime minute -`
# echo $time
function nowtime()
{
	# 参数获取和校验
	type=$1
	dateExp=$2
	# 类型
	if [[ ! -n $type ]]; then
		type=date
	fi
	# 分割字符
	if [[ ! -n $dateExp ]]; then
		dateExp=""
		secExp=""
		exp=""
	else
		secExp=":"
		exp=" "
	fi

	# 格式化字符串
	if [[ hour = $type ]]; then
		fs="+%Y${dateExp}%m${dateExp}%d${exp}%H";
	elif [[ date = $type ]]; then
		fs="+%Y${dateExp}%m${dateExp}%d";
	elif [[ second = $type ]]; then
		fs="+%Y${dateExp}%m${dateExp}%d${exp}%H${secExp}%M${secExp}%S";
	elif [[ minute = $type ]]; then
		fs="+%Y${dateExp}%m${dateExp}%d${exp}%H${secExp}%M";
	elif [[ month = $type ]]; then
		fs="+%Y${dateExp}%m";
	elif [[ year = $type ]]; then
		fs="+%Y";
	else
		fs="+%Y${dateExp}%m${dateExp}%d";
	fi
	date "$fs"
}