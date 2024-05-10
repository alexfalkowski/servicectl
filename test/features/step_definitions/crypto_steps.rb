# frozen_string_literal: true

When('we rotate the aes key') do
  env = {
    'CONFIG_FILE' => '.config/client.yml',
    'AES_CONFIG_FILE' => 'reports/aes_client.yml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'aes', '--rotate')
  pid = spawn(env, cmd, %i[out err] => ['reports/aes.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('it should run sucessfully') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a config file {string}') do |file|
  expect(File.exist?(file)).to be true
end
