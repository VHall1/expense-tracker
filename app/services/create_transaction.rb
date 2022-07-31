def CreateTransaction
  # include sidekiq

  def process(params)
    title, category, value, created_at = params[:]

    transaction = Transaction.new(
      title: params[:reference],
      category: params[:spendingCategory],
      value: params[:amount][:minorUnits] * 100.0,
      created_at: params[:eventTimestamp]
    )

    if transaction.save
      # write to google sheets
    end
  end
end