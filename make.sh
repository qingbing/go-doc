#!/bin/bash
# author qingbing<780042175@qq.com>
# 本脚步使用方法
# ./push.sh

# 获取脚本所在目录
rPath=$(cd $(dirname $0); pwd)
# 加载公共函数
. $rPath/ShellLib/index.sh
loadFiles $rPath/ShellLib


# 1.进入根目录
printLineTip 1.进入根目录
cd $rPath

# 2.重新生成目录
printLineTip 2.重新生成索引文件
./makeIndex.sh

# 3.添加 README.md
printLineTip 3.添加README.md
git add README.md

# 4.git commit
printLineTip 4.commit
commitMsg=$1
# 检查commit提交信息
if [[ ! -n $commitMsg ]]; then
	commitMsg=`nowtime hour`
fi
git_commit $commitMsg

# 提交 git
printLineTip 5.push
git_push origin master

printLineTip 提交成功

