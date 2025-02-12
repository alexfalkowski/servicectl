# frozen_string_literal: true

When('we try to verify a hook') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'hooks', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/hooks.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we try to rotate a secret for the hook') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'hooks', '-i', 'file:.config/client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/hooks.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a succesful hook verification') do
  expect(@status.exitstatus).to eq(0)
end

Then('we should have a succesful rotated the hook secret') do
  expect(@status.exitstatus).to eq(0)
  expect(File).to exist('secrets/hooks-new')
end
