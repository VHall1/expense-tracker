class CreateTransactions < ActiveRecord::Migration[7.0]
  def change
    create_table :transactions do |t|
      t.integer :amount
      t.string :description, limit: 50
      t.datetime :settled_at
      t.text :notes

      t.timestamps
    end
  end
end
