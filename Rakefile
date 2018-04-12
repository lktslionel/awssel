require_relative "tasks/populate"

namespace :code do
  namespace :deps do
    desc "Install deps pkgs"
    task :add do
      # Remove prg name from global ARGV
      ARGV.shift
      opts = ARGV.join(" ")
      sh "dep ensure -add #{opts}"
    end

    desc "Update and sync deps pkgs"
    task :sync do
      sh "dep ensure"
    end
  end

  namespace :test do
    desc "Start testing env"
    task "prepare" do
      sh "docker-compose up -d"
    end

    desc "Add test entries"
    task "seed" do
      seed("put")
    end

    desc "Run tests"
    task "run" do
      sh "/usr/local/bin/go test -timeout 30s github.com/lktslionel/awssel/env -run -v"
    end

    desc "Clean testing env"
    task "clean" do
      seed("delete")
      sh "docker-compose down"
    end
  end
end
