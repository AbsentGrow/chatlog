#!/bin/sh
# Summary: Docker entrypoint for container runtime.
# Details: 以非 root 身份运行主进程（通过 PUID/PGID 切换），
# 并在容器启动时调整必要目录权限；保持 `set -eu` 失败即退出，确保行为可预期。
set -eu

[ -n "${UMASK:-}" ] && umask "$UMASK"

if [ "$(id -u)" = '0' ]; then
  PUID=${PUID:-1000}
  PGID=${PGID:-1000}

  DATA_DIRS="/app /usr/local/bin"
  for DIR in ${DATA_DIRS}; do
    if [ -d "$DIR" ]; then
      chown -R "${PUID}:${PGID}" "$DIR" || true
    fi
  done

  exec gosu "${PUID}:${PGID}" "$@"
else
  exec "$@"
fi
