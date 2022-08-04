#!/bin/bash

# 判断一个变量是否在数组中
# in_array 1233 "${array[*]}"
function in_array()
{
	val=$1
	arr=$2
	if [[ "$val" == null ]]; then
		return 0
	fi
	for i in ${arr[@]}; do
		if [[ "$i" == "$val" ]]; then
			return 1;
		fi
	done
	return 0
}

