package session

import(
  "net/http"
  "time"

  "../setting"
  "github.com/labstack/echo"
)

// ブラウザのcookieにsession.IDを書き込む
func WriteCookie(c echo.Context, sessionID ID) error{
  cookie := new(http.Cookie)
  cookie.Name = setting.Session.CookieName
  cookie.Value = string(sessionID)
  cookie.Expires = time.Now().Add(setting.Session.CookieExpire)
  c.SetCookie(cookie)
  return nil
}

// ブラウザのcookieからsession.IDを読み込む
func ReadCookie(c echo.Context)(ID, error){
  var sessionID ID
  cookie, err := c.Cookie(setting.Session.CookieName)
  if err != nil{
    return sessionID, err
  }
  sessionID = ID(cookie.Value)
  return sessionID, nil
}
