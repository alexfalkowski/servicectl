# frozen_string_literal: true

When('we rotate the aes key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'aes', '-i', 'file:.config/client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/aes.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we verify the aes key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'aes', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/aes.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we rotate the hmac key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'hmac', '-i', 'file:.config/client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/hmac.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we verify the hmac key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'hmac', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/hmac.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we rotate the rsa key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'rsa', '-i', 'file:.config/client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/rsa.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we verify the rsa key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'rsa', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/rsa.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we rotate the ed25519 key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'ed25519', '-i', 'file:.config/client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/ed25519.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we verify the ed25519 key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'ed25519', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/ed25519.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a succesful rotation of aes keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful verification of aes keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful rotation of hmac keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful verification of hmac keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful rotation of rsa keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful verification of rsa keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful rotation of ed25519 keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful verification of ed25519 keys') do
  expect(@status.exitstatus).to eq(0)
end
