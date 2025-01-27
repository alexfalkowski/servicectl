# frozen_string_literal: true

When('we rotate token secret of kind {string}') do |kind|
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'token', '-i', "file:.config/#{kind}.yml", '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/token.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a succesful rotation of the secret of kind {string}') do |kind|
  expect(@status.exitstatus).to eq(0)
  expect(File).to exist("secrets/#{kind}-new")
end
