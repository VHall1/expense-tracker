module TransactionsHelper
  def list_transactions_for_month(date)
    Transaction.where(
      "bank_created_at >= ?",
      date.at_beginning_of_month
    )
  end

  # [CATEGORY, TOTAL]
  # https://stackoverflow.com/questions/59421316/rails-group-by-column-and-select-column
  def list_transactions_per_category_for_month(date)
    transactions = list_transactions_for_month(date)
    transactions
      .pluck(:category, :value)
      .group_by(&:shift)
      .transform_values(&:flatten)
      .transform_values(&:sum)
  end

  def get_income_for_month(date)
    transactions = list_transactions_for_month(date)
    transactions.where(category: 'INCOME')
      .map(&:value)
      .sum
  end

  def get_expenses_for_month(date)
    transactions = list_transactions_for_month(date)
    transactions.where.not(category: 'INCOME')
      .map(&:value)
      .sum
  end

  def get_diff_for_month(date)
    get_income_for_month(date) + get_expenses_for_month(date)
  end
end
