# frozen_string_literal: true

When('we try to verify a hook') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'hooks', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/hooks.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we try to rotate a secret for the hook') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'hooks', '-i', 'file:.config/client.yml', '-o', 'file:reports/hooks_client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/hooks.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a succesful hook verification') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesfully rotated the hook secret') do
  expect(@status.exitstatus).to eq(0)

  src = Nonnative.configurations('.config/client.yml')
  dest = Nonnative.configurations('reports/hooks_client.yml')

  expect(src.hooks.secret).to_not eq(dest.hooks.secret)
end
