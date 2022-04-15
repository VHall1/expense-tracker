class CreateTransactions < ActiveRecord::Migration[6.1]
  def change
    create_table :transactions, id: :uuid do |t|
      t.belongs_to :account, index: true, foreign_key: true, type: :uuid

      t.decimal :value, precision: 10, scale: 2
      t.string :description, null: true

      t.timestamps
    end
  end
end
