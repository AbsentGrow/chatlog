#!/bin/bash
# opensearch/init/install-ik.sh

set -e # 遇到错误立即退出

echo ">>> Starting IK plugin installation..."

# 1. 离线安装 IK 插件
PLUGIN_ZIP="/usr/share/opensearch/plugins-local/opensearch-analysis-ik-3.2.0.0.zip"
if [ ! -f "$PLUGIN_ZIP" ]; then
  echo "ERROR: Plugin ZIP not found at $PLUGIN_ZIP" >&2
  exit 1
fi

/usr/share/opensearch/bin/opensearch-plugin install --batch "file://$PLUGIN_ZIP"

# 2. 替换 IK 配置文件 (热更新配置)
IK_CONFIG_SRC="/usr/share/opensearch/init/IKAnalyzer.cfg.xml"
IK_CONFIG_DST="/usr/share/opensearch/plugins/ik/config/IKAnalyzer.cfg.xml"

if [ -f "$IK_CONFIG_SRC" ]; then
  echo ">>> Copying custom IK config..."
  cp "$IK_CONFIG_SRC" "$IK_CONFIG_DST"
  echo ">>> IK config copied to $IK_CONFIG_DST"
else
  echo ">>> No custom IK config found, using default."
fi

echo ">>> IK plugin installation and config completed."