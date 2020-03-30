#!/usr/bin/env bash
set -euxo pipefail
#-x 解决多条命令连续输出
#-e 错误退出

#$* $@ 相等
# $* 和 $@ 不被双引号" "包围时，它们之间没有任何区别
#"$*"会将所有的参数从整体上看做一份数据，而不是把每个参数都看做一份数据。
#"$@"仍然将每个参数都看作一份数据，彼此之间是独立的。

currentPath=$(dirname $0)
${GOPATH}/bin/migrate -database="mysql://root:root@tcp(127.0.0.1:3306)/firmeve" -path="${currentPath}/../web/backend/database/migrations" $*
