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

Then('we should have a succesful rotation of aes keys') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/aes_client.yml')).to be true

  src = Nonnative.configurations('.config/client.yml')
  dest = Nonnative.configurations('reports/aes_client.yml')

  expect(src.crypto.aes.key).to_not eq(dest.crypto.aes.key)
end

Then('we should have a succesful rotation of hmac keys') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/hmac_client.yml')).to be true

  src = Nonnative.configurations('.config/client.yml')
  dest = Nonnative.configurations('reports/hmac_client.yml')

  expect(src.crypto.hmac.key).to_not eq(dest.crypto.hmac.key)
end
