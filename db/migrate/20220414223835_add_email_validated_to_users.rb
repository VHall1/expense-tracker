class AddEmailValidatedToUsers < ActiveRecord::Migration[6.1]
  def change
    add_column :users, :email_validated, :boolean, default: false
  end
end
