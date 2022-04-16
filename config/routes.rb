Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html

  resources :users

  get '/register' => 'users#new'
  match '/login' => 'users#login', via: [:get, :post]

  resources :accounts do
    resources :transactions
  end

  root 'accounts#index'
end
