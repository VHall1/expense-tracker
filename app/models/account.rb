class Account < ApplicationRecord
  belongs_to :user

  has_many :transactions, dependent: :destroy

  CHOICES = [
    ['Checking', 'Accounts::Checking'],
    ['Savings', 'Accounts::Savings'],
  ]

  def self.setup_account(user, initial_balance)
    self.user = user
    self.balance = initial_balance
    self
  end
end
