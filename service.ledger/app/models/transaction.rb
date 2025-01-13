class Transaction < ApplicationRecord
  def amount_in_major_units
    self.amount / 100.00
  end
end
