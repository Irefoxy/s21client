# s21client 🍻🫃

Клиент для внутреннего GQL API платформы edu.21-school.ru.

```golang
cli := s21client.New(s21client.DefaultAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))

user, err := cli.R().GetCurrentUser(requests.Request_Variables_GetCurrentUser{})
if err != nil {
  t.Error(err)
}

fmt.Println(user)
```
