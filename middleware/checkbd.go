package middleware

import(
  "net/http"
  "twittor/bd"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request){
  if bd.CheckConnection()==0{
    http.Error(w, "Conexión perdida con BD", 500)
    return
  }
  next.ServeHTTP(w, r)
}
}
