



namespace :code do 

  namespace :deps do
    desc "Install deps pkgs"  
    task :add do
      # Remove prg name from global ARGV
      ARGV.shift
      opts = ARGV.join(" ")
      sh "dep ensure -add #{opts}"
    end
  end  

  namespace :test do 
    desc "Start testing env"
    task "prepare" do 
      sh "docker-compose up -d" 
    end
    
    desc "Run tests"
    task "run" do 
      
    end
    
    desc "Clean testing env"
    task "clean" do 
      sh "docker-compose down" 
    end
  end 

end 
