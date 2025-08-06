# Gator CLI

**Gator** — это CLI-приложение на Go для работы с RSS-лентами, пользователями и подписками. Отлично подходит для изучения Go, SQL и архитектуры CLI-инструментов.

---

## 📦 Требования

Перед началом убедитесь, что у вас установлены:

- [Go](https://golang.org/doc/install) 1.21+
- [PostgreSQL](https://www.postgresql.org/download/)

---

## 🚀 Установка

```bash
go install github.com/sergleq/gator@latest
```
---

## Команды

# Зарегистрировать пользователя
gator register alice

# Войти как пользователь
gator login alice

# Добавить ленту
gator addfeed "Boot.dev Blog" https://www.wagslane.dev/index.xml

# Подписаться на ленту
gator follow https://www.wagslane.dev/index.xml

# Посмотреть подписки
gator following

# Посмотреть всех пользователей
gator users

# Посмотреть все ленты
gator feeds

# Разорвать подписку
gator unfollow https://www.wagslane.dev/index.xml

# Сбросить базу
gator reset

# Агрегация постов из RSS
gator agg 1m


