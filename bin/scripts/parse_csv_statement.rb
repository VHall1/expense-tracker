require 'csv'

# Date,Counter Party,Reference,Type,Amount (GBP),Balance (GBP),Spending Category,Notes
CSV.foreach('tmp/statement.csv') do |row|
  p row[4]
  Transaction.create!(
    title: row[2],
    category: row[6],
    value: row[4],
    created_at: row[0]
  )
end
