class CreateAccounts < ActiveRecord::Migration[6.1]
  def change
    create_table :accounts, id: :uuid do |t|
      t.belongs_to :user, index: true, foreign_key: true, type: :uuid

      t.string :name
      t.decimal :balance, precision: 10, scale: 2, default: 0

      t.timestamps
    end
  end
end
