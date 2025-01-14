class CreateUsers < ActiveRecord::Migration[7.0]
  def change
    create_table :users, id: :uuid do |t|
      t.string :first_name, :last_name
      t.string :email, index: { unique: true }

      t.timestamps
    end
  end
end
