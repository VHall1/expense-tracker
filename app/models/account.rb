class Account < ApplicationRecord
  belongs_to :user

  has_many :transactions, dependent: :destroy

  CHOICES = [
    ['Checking', 'Accounts::Checking'],
    ['Savings', 'Accounts::Savings'],
  ]
end
