#!/bin/bash

# git 命令目录
if [[ -n GIT_COMMAND ]]; then
	GIT_COMMAND=git
fi

# 分支切换
function git_checkout()
{
	command "$GIT_COMMAND checkout $1";
	if [[ $? -ne 0 ]]; then
		error_msg "切换分支到 $1 失败，可能文件冲突，请手动操作"
		exit;
	fi
}

# 拉取远程分支到本地
function git_pull()
{
	command "$GIT_COMMAND pull $1 $2"
	if [[ $? -ne 0 ]]; then
		error_msg "拉取远程分支 $1-$2 失败，可能文件冲突，对比分支后手动操作"
		exit;
	fi
}

# 提交本地分支到远程
function git_push()
{
	command "$GIT_COMMAND push $1 $2"
	if [[ $? -ne 0 ]]; then
		error_msg "推送分支 $1-$2 到远程失败，可能文件冲突，请手动操作"
		exit;
	fi
}

# git commit
function git_commit()
{
	command "$GIT_COMMAND commit -m $1"
}