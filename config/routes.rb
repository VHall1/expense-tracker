Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html

  resources :transactions

  post '/webhooks/starling/feed-item' => 'webhooks#starling_feed_item'

  root 'transactions#index'
end
