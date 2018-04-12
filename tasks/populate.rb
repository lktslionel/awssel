require "aws-sdk-ssm"

LOCALSTACK_SSM_ENDPOINT = "http://localhost:4583"

RAW_ENTRIES = <<-ENTRIES
/os/prod/proxy/HOME                                 | /home/username          
/os/dev/support/proxy/PROXY_APP_DIR                 | /os/app                 
/os/api/API_ID                                      | 0ac41e50-3364-11e8-9d7e 
/os/proxy/PROXY_USER                                | os-operator             
/os/proxy/PROXY_PASS                                | 784f43631c05`
/os/prod/support/IT/core/sd-web/OS_SDWEB_URL        | sd-web.com
/os/prod/support/IT/core/sd-web/OS_SDWEB_HTTP_URL   | http://www.sd-web.com
/os/prod/support/IT/core/sd-web/OS_SDWEB_HTTPS_URL  | https://www.sd-web.com
/os/qa/support/IT/core/esb/OS_ESB_MULE_HOST         | esb
/os/qa/support/IT/core/esb/OS_ESB_MULE_PORT         | 8080
/os/qa/support/IT/core/esb/OS_ESB_MULE_HTTP_URL     | http://esb
ENTRIES

def seed(action)
  api_operation = action + "_parameter"
  parameters = []

  RAW_ENTRIES.split("\n").each do |entry|
    path, val = entry.split("|")
    parameters << {path: path.chomp.strip, value: val.chomp.strip}
  end

  begin
    client = Aws::SSM::Client.new(region: "eu-west-1",
                                  endpoint: LOCALSTACK_SSM_ENDPOINT)

    parameters.each do |param|
      case action
      when "put"
        opts = {
          name: param[:path],
          value: param[:value],
          type: "string",
        }
      when "delete"
        opts = {
          name: param[:path],
        }
      end
      puts "#{action} entry : #{opts} "
      client.send api_operation.to_sym, opts
    end
  rescue Aws::SSM::Errors::ServiceError => e
    puts "ERROR: #{e}"
  end
end
