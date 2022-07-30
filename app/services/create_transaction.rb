def CreateTransaction
  # include sidekiq

  def process(params)
    title, category, value, created_at = params[:]

    transaction = Transaction.new(
      title: title,
      category: category,
      value: value,
      created_at: created_at
    )

    if transaction.save
      # write to google sheets
    end
  end
end