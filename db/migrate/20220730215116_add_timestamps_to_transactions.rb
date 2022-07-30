class AddTimestampsToTransactions < ActiveRecord::Migration[6.1]
  def change
    rename_column :transactions, :created_at, :bank_created_at
    change_column :transactions, :bank_created_at, :date

    # nullable for now
    add_timestamps :transactions
  end
end
