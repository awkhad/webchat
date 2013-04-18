require 'active_record'
require "yaml"

ENVIRONMENT = "development"
CONFIG = YAML::load(File.open("conf/database.yml"))

namespace :db do

 # rake db:create[env]
  desc "create"
  task :create do 
    #args.with_defaults(:env => "development")
    ActiveRecord::Base.establish_connection(CONFIG[ENVIRONMENT].merge("database" => nil))
    ActiveRecord::Base.connection.create_database(CONFIG[ENVIRONMENT]["database"])
  end

  # rake db:migrate
  # raek db:migrate[version]
  desc "migrate"
  task :migrate, [:version] do |t, args|
    connected_db(ENVIRONMENT)
    ActiveRecord::Migration.verbose = true
    ActiveRecord::Migrator.migrate("db/migrate/", args['version'] ? args['version'].to_i : nil)
  end

  desc "drop"
  task :drop do
    connected_db(ENVIRONMENT)
    ActiveRecord::Base.connection.drop_database CONFIG[ENVIRONMENT]["database"]
  end
end

def connected_db(env)
    ActiveRecord::Base.establish_connection(CONFIG[env])
end

#def load_config
#    YAML::load(File.open("conf/database.yml"))
#end

