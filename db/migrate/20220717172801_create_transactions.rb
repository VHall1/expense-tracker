class CreateTransactions < ActiveRecord::Migration[6.1]
  def change
    create_table :transactions do |t|
      t.string :title
      t.string :category
      t.decimal :value, precision: 10, scale: 2
      t.datetime :created_at
    end
  end
end
