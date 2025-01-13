Rails.application.routes.draw do
  get "login/:email", to: "auth#login", as: "auth_login"
  post "register", to: "auth#register", as: "auth_register"
end
