# frozen_string_literal: true

When('we rotate the aes key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'aes', '-i', 'file:.config/client.yml', '-o', 'file:reports/aes_client.yml',
                                '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/aes.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we verify the aes key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'aes', '-i', 'file:.config/client.yml', '-o', 'file:reports/hmac_client.yml',
                                '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/aes.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we rotate the hmac key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'hmac', '-i', 'file:.config/client.yml', '-o', 'file:reports/hmac_client.yml',
                                '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/hmac.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we verify the hmac key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'hmac', '-i', 'file:.config/client.yml', '-o', 'file:reports/hmac_client.yml',
                                '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/hmac.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we rotate the rsa key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'rsa', '-i', 'file:.config/client.yml', '-o', 'file:reports/rsa_client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/rsa.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we verify the rsa key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'rsa', '-i', 'file:.config/client.yml', '-o', 'file:reports/rsa_client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/rsa.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we rotate the ed25519 key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'ed25519', '-i', 'file:.config/client.yml', '-o', 'file:reports/ed25519_client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/ed25519.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we verify the ed25519 key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'ed25519', '-i', 'file:.config/client.yml', '-o', 'file:reports/ed25519_client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/ed25519.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a succesful rotation of aes keys') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/aes_client.yml')).to be true

  src = Nonnative.configurations('.config/client.yml')
  dest = Nonnative.configurations('reports/aes_client.yml')

  expect(src.crypto.aes.key).to_not eq(dest.crypto.aes.key)
end

Then('we should have a succesful verification of aes keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful rotation of hmac keys') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/hmac_client.yml')).to be true

  src = Nonnative.configurations('.config/client.yml')
  dest = Nonnative.configurations('reports/hmac_client.yml')

  expect(src.crypto.hmac.key).to_not eq(dest.crypto.hmac.key)
end

Then('we should have a succesful verification of hmac keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful rotation of rsa keys') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/rsa_client.yml')).to be true

  src = Nonnative.configurations('.config/client.yml')
  dest = Nonnative.configurations('reports/rsa_client.yml')

  expect(src.crypto.rsa.public).to_not eq(dest.crypto.rsa.public)
  expect(src.crypto.rsa.private).to_not eq(dest.crypto.rsa.private)
end

Then('we should have a succesful verification of rsa keys') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful rotation of ed25519 keys') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/ed25519_client.yml')).to be true

  src = Nonnative.configurations('.config/client.yml')
  dest = Nonnative.configurations('reports/ed25519_client.yml')

  expect(src.crypto.ed25519.public).to_not eq(dest.crypto.ed25519.public)
  expect(src.crypto.ed25519.private).to_not eq(dest.crypto.ed25519.private)
end

Then('we should have a succesful verification of ed25519 keys') do
  expect(@status.exitstatus).to eq(0)
end
