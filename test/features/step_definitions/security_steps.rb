# frozen_string_literal: true

When('we try to verify a key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'token', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/security.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we try to rotate a key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'token', '-i', 'file:.config/client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/security.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a succesful key verification') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesfully rotated the key') do
  expect(@status.exitstatus).to eq(0)
end
