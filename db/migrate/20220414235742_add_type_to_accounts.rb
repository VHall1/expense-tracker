class AddTypeToAccounts < ActiveRecord::Migration[6.1]
  def change
    add_column :accounts, :type, :string, index: true
  end
end
