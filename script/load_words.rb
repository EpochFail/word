require "pg"

connection = PGconn.connect(:hostaddr=>"127.0.0.1", :port=>"5432", :dbname=>"jsname_dev", :user=>"jsname", :password=>"test")

File.readlines("result.txt").each do |line|
  word = line.strip
  connection.exec("insert into words (word, rating) values ('#{word}', 0)")
  puts word
end

