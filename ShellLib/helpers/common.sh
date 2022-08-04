#!/bin/bash

# 执行并记录命令到控制台
function command()
{
	tip_msg "$@"
	$@
	return $?
}
