class CreateUsers < ActiveRecord::Migration[6.1]
  def change
    enable_extension 'pgcrypto'

    create_table :users, id: :uuid do |t|
      t.string :name
      t.string :email

      t.timestamps
    end
  end
end
