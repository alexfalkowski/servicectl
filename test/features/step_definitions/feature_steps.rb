# frozen_string_literal: true

When('we try to connect to feature') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'feature', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/feature.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a succesful feature connection') do
  expect(@status.exitstatus).to eq(0)
end
