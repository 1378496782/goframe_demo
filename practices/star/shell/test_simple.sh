#!/bin/bash

echo "🚀 Star 项目 - 简单测试"
echo "===================="
echo ""

BASE_URL="http://127.0.0.1:8000"
USERNAME="testuser"
PASSWORD="123456"

echo "📋 现在不需要传 receiveId 了！"
echo "   会自动发送给默认接收者：张富伟"
echo ""

# 检查服务
echo "1️⃣ 检查服务..."
if ! curl -s "$BASE_URL/swagger" > /dev/null; then
  echo "❌ 服务未运行！请先启动服务："
  echo "   cd goframe_demo/practices/star"
  echo "   gf run main.go"
  exit 1
fi
echo "✅ 服务正常"
echo ""

# 登录
echo "2️⃣ 登录用户..."
curl -s -X POST "$BASE_URL/v1/users/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "'"$USERNAME"'",
    "password": "'"$PASSWORD"'",
    "email": "test@example.com"
  }' > /dev/null

LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/v1/users/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "'"$USERNAME"'",
    "password": "'"$PASSWORD"'"
  }')

TOKEN=$(echo "$LOGIN_RESPONSE" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
  echo "❌ 获取 token 失败！"
  exit 1
fi

echo "✅ Token 获取成功"
echo ""

# 确保有单词
echo "3️⃣ 创建测试单词..."
curl -s -X POST "$BASE_URL/v1/word/create/batch" \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{
    "words": [
      {
        "word": "hello",
        "definition": "Used as a greeting",
        "exampleSentence": "Hello, nice to meet you!",
        "chineseTranslation": "你好",
        "pronunciation": "/həˈləʊ/",
        "proficiencyLevel": 3
      },
      {
        "word": "world",
        "definition": "The earth, together with all countries and peoples",
        "exampleSentence": "The world is beautiful",
        "chineseTranslation": "世界",
        "pronunciation": "/wɜːld/",
        "proficiencyLevel": 2
      }
    ]
  }' > /dev/null

echo "✅ 测试单词已创建"
echo ""

# 发送提醒（不传 receiveId！）
echo "4️⃣ 发送学习提醒（不传 receiveId，使用默认值）..."
echo ""

SEND_RESPONSE=$(curl -s -X POST "$BASE_URL/v1/lark/reminder/send" \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{}')

echo "响应:"
echo "$SEND_RESPONSE"
echo ""

echo "===================="
echo "🎉 测试完成！"
echo ""
echo "📚 请检查你的飞书是否收到学习提醒！"
echo "💡 再也不用手动传 receiveId 了！"
