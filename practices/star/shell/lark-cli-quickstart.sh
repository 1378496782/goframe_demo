#!/bin/bash

# ========================================
# 飞书 CLI 快速上手指南脚本
# ========================================

echo "🚀 欢迎使用飞书 CLI 快速上手指南！"
echo ""

# 检查是否安装了 feishu-cli
if ! command -v feishu-cli &> /dev/null; then
    echo "❌ feishu-cli 未安装，请先安装"
    exit 1
fi

# 检查是否安装了 lark-cli
if ! command -v lark-cli &> /dev/null; then
    echo "❌ lark-cli 未安装，请先安装"
    exit 1
fi

# 询问用户邮箱
read -p "📧 请输入你的飞书邮箱（用于测试消息发送）: " USER_EMAIL

if [ -z "$USER_EMAIL" ]; then
    echo "❌ 邮箱不能为空"
    exit 1
fi

echo ""
echo "========================================"
echo "第一步：发送测试消息"
echo "========================================"

# 创建测试卡片
cat > /tmp/test_card.json << 'EOF'
{
  "header": {
    "template": "blue",
    "title": {"tag": "plain_text", "content": "🎉 飞书 CLI 测试消息"}
  },
  "elements": [
    {"tag": "markdown", "content": "**恭喜！你成功运行了飞书 CLI 快速上手指南！**\n\n这是一条自动发送的测试消息，用于验证飞书 CLI 是否正常工作。"},
    {"tag": "hr"},
    {"tag": "div", "fields": [
      {"is_short": true, "text": {"tag": "lark_md", "content": "**状态**\n✅ 成功"}},
      {"is_short": true, "text": {"tag": "lark_md", "content": "**日期**\n2026-04-21"}}
    ]},
    {"tag": "note", "elements": [{"tag": "plain_text", "content": "由 lark-cli-quickstart.sh 自动发送"}]}
  ]
}
EOF

echo "📤 正在发送测试消息到 $USER_EMAIL ..."

feishu-cli msg send \
  --receive-id-type email \
  --receive-id "$USER_EMAIL" \
  --msg-type interactive \
  --content-file /tmp/test_card.json

if [ $? -eq 0 ]; then
    echo "✅ 消息发送成功！请查收飞书"
else
    echo "❌ 消息发送失败"
fi

echo ""
echo "========================================"
echo "第二步：查看今日日程"
echo "========================================"
echo "📅 正在获取今日日程..."

lark-cli calendar agenda

echo ""
echo "========================================"
echo "第三步：查看个人信息"
echo "========================================"
echo "👤 正在获取你的个人信息..."

lark-cli contact me

echo ""
echo "========================================"
echo "🎉 快速上手完成！"
echo "========================================"
echo ""
echo "📚 下一步建议："
echo "1. 尝试发送更多类型的消息（图片、文件、富文本）"
echo "2. 查看第二阶段：结合你的项目"
echo "3. 探索更多飞书 CLI 功能"
echo ""

# 清理临时文件
rm -f /tmp/test_card.json /tmp/welcome_card.json
