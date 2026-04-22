#!/bin/bash

echo "🚀 Star 项目 - 完整测试"
echo "========================"
echo ""

BASE_URL="http://127.0.0.1:8000"
RECEIVE_ID="ou_873d43281881570cdc0757ba3b977775"
USERNAME="testuser"
PASSWORD="123456"

echo "📋 配置："
echo "   App ID: cli_a940a8f1f439dbca"
echo "   接收者: 张富伟 (ou_873d43281881570cdc0757ba3b977775)"
echo ""

# 检查服务是否运行
echo "1️⃣ 检查服务..."
if ! curl -s "$BASE_URL/swagger" > /dev/null; then
  echo "❌ 服务未运行！请先启动服务："
  echo "   cd goframe_demo/practices/star"
  echo "   gf run main.go"
  exit 1
fi
echo "✅ 服务正常"
echo ""

# 注册/登录
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
  echo "响应: $LOGIN_RESPONSE"
  exit 1
fi

echo "✅ Token 获取成功"
echo ""

# 创建测试单词
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
      },
      {
        "word": "star",
        "definition": "A fixed luminous point in the night sky",
        "exampleSentence": "Look at the stars!",
        "chineseTranslation": "星星",
        "pronunciation": "/stɑː(r)/",
        "proficiencyLevel": 4
      }
    ]
  }' > /dev/null

echo "✅ 测试单词已创建"
echo ""

# 发送学习提醒
echo "4️⃣ 发送学习提醒到飞书..."
echo "   接收者: $RECEIVE_ID"
echo ""

SEND_RESPONSE=$(curl -s -X POST "$BASE_URL/v1/lark/reminder/send" \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{
    "receiveId": "'"$RECEIVE_ID"'"
  }')

echo "响应:"
echo "$SEND_RESPONSE"
echo ""

echo "========================"
echo "🎉 测试完成！"
echo ""
echo "📚 请检查你的飞书是否收到学习提醒！"
echo "💡 如果有问题，请查看服务日志！"
