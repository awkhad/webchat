class CreateUsersTable < ActiveRecord::Migration 
  def change
    create_table :user do |t|
      t.string :name
      t.string :email
      t.string :salt
      t.string :encryptpasswd
      t.datetime :created
      t.datetime :updated
    end
  end
end
