class Transaction < ApplicationRecord
  belongs_to :account

  after_create :add_to_account_balance
  after_destroy :subtract_from_account_balance

  def add_to_account_balance
    account.balance += value
    account.save
  end

  def subtract_from_account_balance
    account.balance -= value
    account.save
  end
end
