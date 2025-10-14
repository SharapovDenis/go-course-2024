#!/bin/bash

# URL для API
USER_URL="http://localhost:8080/api/v1/users"
AD_URL="http://localhost:8080/api/v1/ads"

# Данные для пользователей
names=("Иван" "Мария" "Алексей" "Ольга" "Дмитрий")
emails=("ivan@example.com" "maria@example.com" "alex@example.com" "olga@example.com" "dmitry@example.com")

# Данные для объявлений
titles=("Продам велосипед" "Сдам квартиру" "Ищу работу" "Продам ноутбук" "Отдам котенка")
texts=("В отличном состоянии" "В центре города, недорого" "Опыт 5 лет в IT" "Игровой, почти новый" "Очень ласковый")

declare -a user_ids
declare -a ad_ids

echo "=== Создаю пользователей ==="
for i in {0..4}
do
  json=$(jq -n \
    --arg name "${names[$i]}" \
    --arg email "${emails[$i]}" \
    '{name: $name, email: $email}')

  echo "Создаю пользователя ${names[$i]}..."
  response=$(curl -s -X POST "$USER_URL" \
    -H "Content-Type: application/json" \
    -d "$json")

  echo "Ответ сервера: $response"  # Можно закомментировать позже
  user_id=$(echo "$response" | jq -r '.data.id')

  if [[ "$user_id" == "null" || -z "$user_id" ]]; then
    echo "Ошибка при создании пользователя ${names[$i]}, пропускаю..."
    continue
  fi

  echo "ID созданного пользователя: $user_id"
  user_ids+=("$user_id")
  echo "---"
done

echo "=== Создаю объявления ==="
for i in {0..4}
do
  json=$(jq -n \
    --arg title "${titles[$i]}" \
    --arg text "${texts[$i]}" \
    --argjson userid "${user_ids[$i]}" \
    '{title: $title, text: $text, user_id: $userid}')

  echo "Создаю объявление $((i+1))... $json"
  response=$(curl -s -X POST "$AD_URL" \
    -H "Content-Type: application/json" \
    -d "$json")

  echo "Ответ сервера: $response"  # Для отладки
  ad_id=$(echo "$response" | jq -r '.data.id')
  echo "ID созданного объявления: $ad_id"
  ad_ids+=("$ad_id")
  echo "---"
done

echo "=== Меняю статус двух объявлений ==="
for idx in 0 2
do
  ad_id="${ad_ids[$idx]}"
  user_id="${user_ids[$idx]}"

  json=$(jq -n \
    --argjson published true \
    --argjson userid "$user_id" \
    '{published: $published, user_id: $userid}')

  echo "Меняю статус объявления ID=$ad_id..."
  curl -s -X PUT "$AD_URL/$ad_id/status" \
    -H "Content-Type: application/json" \
    -d "$json"
  echo -e "\n---\n"
done
