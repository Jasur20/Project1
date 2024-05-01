import telebot
from telebot import types
import random
import requests


bot=telebot.TeleBot('6470818555:AAFqBtXXGzUpYLH5YDsZYEvJrFTV-JBnYiY')

first = ["Сегодня — идеальный день для новых начинаний.","Оптимальный день для того, чтобы решиться на смелый поступок!","Будьте осторожны, сегодня звёзды могут повлиять на ваше финансовое состояние.","Лучшее время для того, чтобы начать новые отношения или разобраться со старыми.","Плодотворный день для того, чтобы разобраться с накопившимися делами."]
second = ["Но помните, что даже в этом случае нужно не забывать про","Если поедете за город, заранее подумайте про","Те, кто сегодня нацелен выполнить множество дел, должны помнить про","Если у вас упадок сил, обратите внимание на","Помните, что мысли материальны, а значит вам в течение дня нужно постоянно думать про"]
second_add = ["отношения с друзьями и близкими.","работу и деловые вопросы, которые могут так некстати помешать планам.","себя и своё здоровье, иначе к вечеру возможен полный раздрай.","бытовые вопросы — особенно те, которые вы не доделали вчера.","отдых, чтобы не превратить себя в загнанную лошадь в конце месяца."]
third = ["Злые языки могут говорить вам обратное, но сегодня их слушать не нужно.","Знайте, что успех благоволит только настойчивым, поэтому посвятите этот день воспитанию духа.","Даже если вы не сможете уменьшить влияние ретроградного Меркурия, то хотя бы доведите дела до конца.","Не нужно бояться одиноких встреч — сегодня то самое время, когда они значат многое.","Если встретите незнакомца на пути — проявите участие, и тогда эта встреча посулит вам приятные хлопоты."]

@bot.message_handler(commands=['help','start'])
def get_help(message):
    if message.text=="/start":
        keyboard = types.InlineKeyboardMarkup()

        key_info = types.InlineKeyboardButton(text='Информация о карте', callback_data='info')
        keyboard.add(key_info)
        key_balance = types.InlineKeyboardButton(text='Узнать баланс', callback_data='balance')
        keyboard.add(key_balance)
        key_last_tr = types.InlineKeyboardButton(text='Последняя транзакция', callback_data='last_tr')
        keyboard.add(key_last_tr)
        bot.send_message(message.from_user.id, text='Что Вы хотите узнать?', reply_markup=keyboard)
    elif message.text=="/help":
        bot.send_message(message.from_user.id,'hello!')

@bot.message_handler(content_types=['text'])
def get_text_message(message):
    if message.text=="Привет" or message.text=="Hi":
        bot.send_message(message.from_user.id,'Привет, чем я могу помочь?')
    else:
        bot.send_message(message.from_user.id,"Я тебя не понимаю, введите - /help")

@bot.callback_query_handler(func=lambda call: True)
def callback_worker(call):
  # Если нажали на одну из 12 кнопок — выводим гороскоп
  if call.data == 'info':
      BASE_URL = ('https://api.github.com')
      query_params = {
          "limit": 3
        }
      response = requests.get(f"{BASE_URL}", params=query_params)
      print(response.text)
      bot.send_message(call.message.chat.id, response.text)

  if call.data=='balance':
      BASE_URL = 'https://fakestoreapi.com'
      query_params = {
          "limit": 3
      }
      response = requests.get(f"{BASE_URL}/balance", params=query_params)
      bot.send_message(call.message.chat.id,response.text)

  if call.data=='last_tr':
      BASE_URL = 'https://fakestoreapi.com'
      query_params = {
          "limit": 3
      }
      response = requests.get(f"{BASE_URL}/lastTranc", params=query_params)


bot.infinity_polling()

