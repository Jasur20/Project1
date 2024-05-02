import telebot
from telebot import types
import random
import requests

bot=telebot.TeleBot('6470818555:AAFqBtXXGzUpYLH5YDsZYEvJrFTV-JBnYiY')

card_number=''
password=''

@bot.message_handler(content_types=['text'])
def start(message):
    global card_number
    bot.message_handler(message.from_user.id,'Введите номер своей карты')
    card_number=message.text
    if len(card_number)!=16:
        bot.send_message(message.from_user.id,'Неверный ввод номера карты.Попробуйте снова!!')
    else:
        bot.register_next_step_handler(message,get_password)

def get_password(message):
    global password
    password=message.text
    if len(password)<8:
        bot.send_message(message.from_user.id,'Неверный пароль!!Повторите попытку.')
    else:
        bot.send_message(message.from_user.id,'Вы успешно вошли в свой личный кабинет')



@bot.message_handler(commands=['help','start'])
def get_help(message):
    if message.text=="/help":
        keyboard = types.InlineKeyboardMarkup()

        key_info = types.InlineKeyboardButton(text='Информация о карте', callback_data='info')
        keyboard.add(key_info)
        key_balance = types.InlineKeyboardButton(text='Узнать баланс', callback_data='balance')
        keyboard.add(key_balance)
        key_last_tr = types.InlineKeyboardButton(text='Последняя транзакция', callback_data='last_tr')
        keyboard.add(key_last_tr)
        key_createac=types.InlineKeyboardButton(text='Создать личный кабинет',callback_data='create',url='https://brt.tj/ru')
        keyboard.add(key_createac)
        key_autorization=types.InlineKeyboardButton(text='Войти в личный кабинет',callback_data='in')
        keyboard.add(key_autorization)
        bot.send_message(message.from_user.id, text='Что Вы хотите узнать?', reply_markup=keyboard)
    elif message.text=="/start":
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
      if response.status_code==200:
          bot.send_message(call.message.chat.id, response.text)
      elif response.status_code==400:
          bot.send_message(call.message.chat.id,'Данные введены некоректно.Исправьте.')
      elif response.status_code==500:
          bot.send_message(call.message.chat.id,'Сервер временно недоступен.')
      else:
          bot.send_message(call.message.chat.id,'Проблемы с сервисом.')

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
  if call.data=='in':
      @bot.message_handler(content_types=['text'])
      def st(message):
          bot.send_message(message.from_user.id,'hi')
      st(message="hello")



bot.infinity_polling()
