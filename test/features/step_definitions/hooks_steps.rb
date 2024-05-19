# frozen_string_literal: true

When('we try to sign a hook') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'hooks', '-i', 'file:.config/client.yml', '--sign')
  pid = spawn({}, cmd, %i[out err] => ['reports/hooks.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a succesful hook signature') do
  expect(@status.exitstatus).to eq(0)
end
