# Gator CLI

**Gator** — это CLI-приложение на Go для работы с RSS-лентами, пользователями и подписками. Отлично подходит для изучения Go, SQL и архитектуры CLI-инструментов. (учебный проект boot.dev)

---

## Команды

### Зарегистрировать пользователя
```bash
gator register alice
```
### Войти как пользователь
```bash
gator login alice
```
### Добавить ленту
```bash
gator addfeed "Boot.dev Blog" <url>
```
### Подписаться на ленту
```bash
gator follow <url>
```
### Посмотреть подписки
```bash 
gator following
```
### Посмотреть всех пользователей
```bash
gator users
```
### Посмотреть все ленты
```bash
gator feeds
```
### Разорвать подписку
```bash
gator unfollow <url>
```
### Сбросить базу
```bash
gator reset
```
### Агрегация постов из RSS
```bash
gator agg 1m
```


