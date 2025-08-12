#!/bin/bash

URL="http://localhost:8080/api/v1/ads"

titles=("Продам велосипед" "Сдам квартиру" "Ищу работу" "Продам ноутбук" "Отдам котенка")
texts=("В отличном состоянии" "В центре города, недорого" "Опыт 5 лет в IT" "Игровой, почти новый" "Очень ласковый")
user_ids=(101 102 103 104 105)

declare -a ad_ids

# Создаём объявления
for i in {0..4}
do
  json=$(jq -n \
    --arg title "${titles[$i]}" \
    --arg text "${texts[$i]}" \
    --argjson userid "${user_ids[$i]}" \
    '{title: $title, text: $text, user_id: $userid}')

  echo "Создаю объявление $((i+1))..."
  response=$(curl -s -X POST "$URL" \
    -H "Content-Type: application/json" \
    -d "$json")

  echo "Ответ сервера: $response"  # Для отладки, можно потом убрать
  ad_id=$(echo "$response" | jq -r '.data.id')
  echo "ID созданного объявления: $ad_id"
  ad_ids+=("$ad_id")
  echo "---"
done

# Меняем статус для двух объявлений
for idx in 0 2
do
  ad_id="${ad_ids[$idx]}"
  user_id="${user_ids[$idx]}"

  json=$(jq -n \
    --argjson published true \
    --argjson userid "$user_id" \
    '{published: $published, user_id: $userid}')

  echo "Меняю статус объявления ID=$ad_id..."
  curl -s -X PUT "$URL/$ad_id/status" \
    -H "Content-Type: application/json" \
    -d "$json"
  echo -e "\n---\n"
done
