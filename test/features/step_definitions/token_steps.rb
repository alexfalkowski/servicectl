# frozen_string_literal: true

When('we rotate the token key') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'token', '-i', 'file:.config/client.yml', '--rotate')
  pid = spawn({}, cmd, %i[out err] => ['reports/token.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a succesful rotation of token key') do
  expect(@status.exitstatus).to eq(0)
end
