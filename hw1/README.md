# Copy Trading

- [Описание](#описание)
- [1. Пользовательские сценарии](#1-пользовательские-сценарии)
- [2. Общая архитектура и схема взаимодействия сервисов](#2-общая-архитектура-и-схема-взаимодействия-сервисов)
- [3. Назначение сервисов и их зоны ответственности](#3-назначение-сервисов-и-их-зоны-ответственности)
- [4. Контракты взаимодействия сервисов друг с другом](#4-контракты-взаимодействия-сервисов-друг-с-другом)

## Описание

Я хочу создать веб-сервис который будет делать Copy Trading на различных биржах (Binance, Bybit, bitget).

Основные пользователи: трейдеры - привязывают свой торговый аккаунт на бирже в нашем сервисе, и мы можем получать информацию о его сделках для копирования с этой биржи, подписчики - обычные пользователи, которые подписываются на трейдера за определенную сумму, также привязывают свой торговый аккаунт и мы с него совершаем сделки, копируя сделки трейдера.

Трейдер получает большую часть средств с подписки подписчика.

При выборе трейдера пользователь может посмотреть статистику его торговли в виде графиков. Также после оформления подписки подписчику доступна история его сделок и аналитика по его кошельку.

## 1. Пользовательские сценарии

### Сценарий 1: Регистрация трейдера и привязка торгового аккаунта

![alt-image](https://www.plantuml.com/plantuml/svg/bLJ1QbD15Ds_hnYwxmy8aORYmeKW6EdwaGOjQ2nf19SreLMm9DgDEsFn1wPfNfoysCaldFj7dZlJmVTy9IYbj3DtphddddkIlPDXFXY-VVqgoxl3DmFJ6UI7lO79JqmdEsRbg7jqdFU7PWUNA72zYn9hEFvxAwCTVNQm_zUxAnHoYZcSl3SoneO7o-DQCRoS6cpWiPOB_8HdmSe81SzJ_88BbC-VFFjlJekIFvP52kCfSRKRN5_W25_2XlgZTzsNUVz5ByjQBOFl9BpHBZBI5d8UoFN0mWgB8CEQLYlhcDs7dFs1mJVsA-LKFkaiV30o2HwqdSlOyD3WXSiEzfM4qvBbAsmOotD2ukoz69baVFq7T0LF4nUn4mfQ8nzK0OyUAyCQIPK2hZ54P4f3p0XO43WFnQK-YfpRj1FS8VP7QihYjD9KBkX4HQ6D7i_8uZ2NAGlSc27Iy-ecXC7htyOswuRFL2D_prGp70r7oL3LQ8j_eRrRGFl-1TGPPTo44sY21p447jTXwykx611y6rqZeeo-XV2dJq3OLWZ9nBIVFkOOl9o4ovNEcxJX4lCqQ_qX-PoQN6W2jMdIb6AjTkbGSR71xuZmSXOZcXRTL90uaf2jCGdXNyFnYk7JiU7gTErgtZvNhMHkgMX3peyAhZjFYTMmRL5GJu3BzdhzGtwn_GO0)

### Сценарий 2: Сбор информации о сделках трейдера для копирования

![alt-image](https://www.plantuml.com/plantuml/svg/VLBTIYfH6BtVJ_7XlIzmO0vwJbrq4GJr0eDA2QQY1bsc8X51KWHTXc-mCqUtseoliFORjRvH1bD9GVtsv_fPQq-XtGbRdPlhccdokrggDiDwHt8O8V9t6CFvhlWkFf1mjFfpWWXB95Za96pBnV_ptv1zh1J7SKbOpCEXS5pmRRcCVO_J5rpALJpTvNg68tmAc-wFRqjNOVsoilSdYnczqPb5JCv7FpXWj1MMAqV_Z0a2mNlgqN4veS8zfouMdETPrXy44WI6GIN_Lpt-4Rp1fi6ITJihE2NGVdeyZJ1NhEunvZvDpAGEOtMdUG-MO4WkUOfGYoAZ3IyD7E9L4iUanRCJ8TFIFp68CxL6eobuvJmK5S2Kc0IVI7P9h16-rhcoRdV28LQjlKQ5rX9LLeeCk-Nz1Gbc15B5Fy1IhvTn8SgujFHj5vjBql2K7oYSzGpNS4eoLL-lIwzvdwXIBvi2F_WqVmC0)

### Сценарий 3: Регистрация подписчика, выбор трейдера и оформление подписки

![alt-image](https://www.plantuml.com/plantuml/svg/dLJDQjj04BxlKoovvmKABWav5moYvEouEHZQDDZEtPLAsf9GWo-zLIrzWOqRfHkhbLzXvetwpGuocvL889UGrSxtCz_C-c0s7qtd5-_Ucj5u_dvgiukJsNWwEJcTsj7CPkOSjvFnv7nqDhTxjAAARlPfGpLv_FkFypqfEpvAwdvInGjQa-SFbWlQuk1mh0M3ZqjBMsgevclwGmqk7ETKuRoal-G3vUjNsRCuR_4N5R3eTbnPXwi4J3NlqCGLtLj0TlG1DCoC6Gmi_K3HR_dAkIZoPKUBFv4NFJiOcCpklqGSBonz1_s65o17U-3ERKiHx5voON7eS9CtnqSWaG30yerSwBH1qy2vNWn_DQZUWP1KWrnlLOagf0MBS81ZWrPnvzGv-Tv6bArjf9G2WDSXe3iKLyhxc7-0nSclLXYNoqYOhv56vD9frYkaTODU5pOGLdGlQa9jOHKN7-KISwrR3-dG_SCXb1eT42um-cWOYW5wawfX8tK8oTQfLQfno4DIkbw7wQ0a3n4KK5lPft4OfRmAQQn3Z5_KOmyF5r5pcfO7HCkVoeoNKfOa1aPjDUlTjtXC_O4R00K_x8MZTgjlZOikT1YG0Yo3vKemwSkcIgdIvnDpNqaihObIumrQcDG33q5IktrijHFk9vz3VpBHGoZZAscca0K4nMU0c-vgEtDmUlO6Fyl_0G00)

### Сценацрий 4: Копирование сделок трейдера на аккаунт подписчика

![alt-image](https://www.plantuml.com/plantuml/svg/bPFVIZ916CRFvofUFFS6FjWF_VuST104UWEBIWcceWOTgX4HGM44GIVb7GnRezkgwoqySqSz3EkskIdKGJEpx_l-tcUUTopr1tvlS7RQKbskpLgpwxS7Ki0KnWuH8BGZGOGOQsw7N0DehBY_4pl26mmM_1WLnE_BtsfbYp9XyXnQcCWghEoDlSpNLNj-lV7lpnu5zW9h7ZIF2vYaFoN5954OYH63IuIELpci_9YbII23mPKTPrZb76lYZ9X1k_Z_yzg9tpvkAELvWgUyKJ32XSVStSNp52sJuk-D0xy4Zz1ENcn7E-lic0a6x-0dxhHAogL858d4t3Dnpbf6xJKqyyOf3kOJ5o7SQA1hM_sPbPhh7A8p8O9aZ5dxxHK34OZFYoCt8soQ87cF7GJLwdIwWWSiE1StrncJOhmY_gv2A-5VSeToxWxbemD2oBfbFcNjHh1slu5VMM6cR_eAMFZYdlmIXcLcym1YSX8_d83i04pUCe9JP4RJCwYSCmXS6br96_2urPNUPRkgHhkkIlp7N_e7)

### Сценарий 5: Просмотр подписчиком истории сделок и аналитики

![alt-image](https://www.plantuml.com/plantuml/svg/XL5DJi9G6Do_KtnXpmLCQaIy0IQi2x8WyOUKieSYKMDY8Xjt4czGa4fJf5nXtesSxsdL8eaBNcZdvyrCZ_gX7uI3YtFZjyEhG1g3LhyTT5kTGFo-D4oFQBVTxVcNeLGmHO9v5HdMIFbtOwEAqfhr7TuB4Zl40gaT2M9i-AF2HaYfJPscjklzID8AiM2f2dhVsudO4HyJYdDaJdOQ-6UTasDZF4ymuoKvAUyy8piKlO-Nv7oT8Ys9YUr98fvd6b8zP8C3mHF1BRMnlIMyGCmMKSdKFfHCJRDECQly4kTxAmksoDr1zTPPtYXYnwQc0OfUD9gHixPZuXnOPoulWlmdGiQeapskYCsNdrIBV5EILbYuUdT5jMzZ97-DNR9Ycvb-KMBBGbggg64fpdRY8j65r4CJFEkAc6i2lFq_BZ71gyjx_JdY1m00)

## 2. Общая архитектура и схема взаимодействия сервисов

![alt-image](https://www.plantuml.com/plantuml/svg/dLHDRjD05DxFAKPPmI8tGAZ1i611LCcksB2PEgCsgkjOOmTH8QIs23P5gk02Y1ia9MbTfoHNU7CZldcHoU005csLQBxtyyttlkTu9o-KASR7YSYFXcccZ3gMoJ3LnKccvSWKXoDXT5oez23HiaDVQ4xJBbMqe1B77_Qy8rKkN-hz81dgj10YK_6HEd2rtsXkJ-cAIdjcBphohP3uyycy0cdVqmfW0hXmn0XUIbhHaXRs4ztG4ec9FQSvy2NTKSdtxhq8_vyLiCB_14de1whPUc2xdyOPlb51bSSMjh6TqPIMZiyTPsuoC5MlqWUHKGDjPA2oOcoq3BLvFOprGsQDWd-pGiAInPokBLkhRDuGZB8JwQvfi0TH2xlZcxZbmxYIgjguxQwlA5hlwnekNz0j96o8k6ktrUqhjCyW-Qxscswviz-YikJ5pkcxSzT-12wbYq_PbsjlIS-UX-AT40DLg7sLEwBFoDvW9pJZOQ7TVlZBWuZFMizUhitJ9wqi_XAuZpU0HX4de_60srXWBUvDV8ZvT2x4hn-OVDJjFdRFVJCKReTwVmXjHR050Je6zI1X0_SQU1Dm_aM11uB80vm2g0Vbh8zxlqEn6sw6gSzzyFIvpd_DW9xJdaaxEXtWHVOJ)

## 3. Назначение сервисов и их зоны ответственности

### Сервис управления пользователями (UMS)

- **Назначение**: Аутентификация и авторизация пользователей (трейдеров и подписчиков), управление профилями.
- **Зона ответственности**:
  - Регистрация и вход пользователей.
  - Хранение персональных данных.
  - Управление правами доступа.

### Сервис управления аккаунтами (AMS)

- **Назначение**: Управление привязками торговых аккаунтов пользователей к сервису.
- **Зона ответственности**:
  - Сохранение и управление API ключами биржевых аккаунтов.
  - Обеспечение безопасного взаимодействия с биржами от имени пользователя.

### Сервис управления подписками (SMS)

- **Назначение**: Обработка подписок подписчиков на трейдеров.
- **Зона ответственности**:
  - Создание и управление подписками.
  - Проверка статуса подписок.
  - Расчет выплат трейдерам.

### Сервис сбора сделок (TCS)

- **Назначение**: Сбор данных о сделках трейдеров для последующего копирования.
- **Зона ответственности**:
  - Периодическое получение данных о новых сделках трейдеров.
  - Сохранение сделок в базе данных.

### Сервис копирования сделок (CTS)

- **Назначение**: Копирование сделок трейдеров на аккаунты подписчиков.
- **Зона ответственности**:
  - Обработка новых сделок и размещение соответствующих ордеров на аккаунтах подписчиков.
  - Управление параметрами копирования (например, коэффициент копирования).

### Сервис аналитики (AS)

- **Назначение**: Предоставление статистики и аналитики по торговле трейдеров и подписчиков.
- **Зона ответственности**:
  - Расчет показателей эффективности торговли.
  - Формирование отчетов и графиков.
  - Обработка и агрегация данных о сделках.

### Платежный сервис (PS)

- **Назначение**: Обработка платежей за подписки и выплаты трейдерам.
- **Зона ответственности**:
  - Интеграция с платежными шлюзами.
  - Обработка транзакций.
  - Управление финансовыми операциями.

### Сервис уведомлений (NS)

- **Назначение**: Отправка уведомлений пользователям о важных событиях.
- **Зона ответственности**:
  - Отправка email и push-уведомлений.
  - Настройка предпочтений уведомлений пользователей.

### Сервис интеграции с биржами (EIS)

- **Назначение**: Обеспечение унифицированного взаимодействия с различными биржами.
- **Зона ответственности**:
  - Абстрагирование API бирж.
  - Обработка запросов на размещение ордеров и получение данных.

## 4. Контракты взаимодействия сервисов друг с другом

### 1. Сервис управления пользователями (UMS) <-> Веб-клиент

- **API для регистрации и входа пользователей**:
  - `POST /register` — регистрация нового пользователя.
  - `POST /login` — аутентификация пользователя.

### 2. Сервис управления аккаунтами (AMS) <-> Веб-клиент**

- **API для управления биржевыми аккаунтами**:
  - `POST /accounts` — привязка нового торгового аккаунта.
  - `GET /accounts` — получение списка привязанных аккаунтов.

### 3. Сервис управления подписками (SMS) <-> Веб-клиент**

- **API для управления подписками**:
  - `POST /subscriptions` — создание новой подписки.
  - `GET /subscriptions` — получение с писка подписок пользователя.

### 4. Сервис управления подписками (SMS) <-> Платежный сервис (PS)

- API для обработки платежей:
  - `POST /payments` — инициирование платежа за подписку.
  - `POST /payments/notify` — обработка уведомлений о статусе платежа.

### 5. Сервис сбора сделок (TCS) <-> Сервис управления аккаунтами (AMS)

- API для получения списка активных трейдеров:
  - `GET /traders` — получение списка трейдеров с привязанными аккаунтами.

### 6. Сервис копирования сделок (CTS) <-> Сервис управления подписками (SMS)

- API для получения активных подписчиков трейдера:
  - `GET /subscriptions/{traderId}` — получение списка подписчиков трейдера.

### 7. Сервис копирования сделок (CTS) <-> Сервис управления аккаунтами (AMS)

- API для получения данных аккаунтов подписчиков:
  - `GET /accounts/{subscriberId}` — получение API ключей подписчика.

### 8. Сервис аналитики (AS) <-> База данных сделок (TradeDB)

- Интерфейс для получения данных о сделках:
  - Прямое чтение данных из базы или через API:
    - `GET /trades/{userId}` — получение списка сделок пользователя.

### 9. Сервис интеграции с биржами (EIS) <-> Биржи

- Унифицированные методы взаимодействия с биржами:
  - `POST /orders` — размещение ордеров.
  - `GET /trades` — получение истории сделок.

### 10. Сервис уведомлений (NS) <-> Другие сервисы

- API для отправки уведомлений:
  - `POST /notifications` — создание нового уведомления для пользователя.
