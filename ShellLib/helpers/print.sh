#!/bin/bash

# 打印分隔线
function printLine()
{
	echo -e "========================================";
}

# 打印分隔消息
function printLineTip()
{
	echo -e "==========  $1  ==========";
}


# 打印列表（数组）
# eg : printList "${array[*]}"
function printList()
{
	if [[ $2 ]]; then
		echo -e "$2 : "
	fi

	i=0
	arr=$1
    for item in ${arr[*]}; do
        echo $i ":" $item
        let i++
    done
}